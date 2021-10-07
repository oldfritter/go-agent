// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/oldfritter/go-agent/v3/oldfritter"
)

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world")
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "New Relic Go Agent Version: "+oldfritter.Version)
}

func noticeError(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "noticing an error")

	txn := oldfritter.FromContext(r.Context())
	txn.NoticeError(errors.New("my error message"))
}

func noticeErrorWithAttributes(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "noticing an error")

	txn := oldfritter.FromContext(r.Context())
	txn.NoticeError(oldfritter.Error{
		Message: "uh oh. something went very wrong",
		Class:   "errors are aggregated by class",
		Attributes: map[string]interface{}{
			"important_number": 97232,
			"relevant_string":  "zap",
		},
	})
}

func customEvent(w http.ResponseWriter, r *http.Request) {
	txn := oldfritter.FromContext(r.Context())

	io.WriteString(w, "recording a custom event")

	txn.Application().RecordCustomEvent("my_event_type", map[string]interface{}{
		"myString": "hello",
		"myFloat":  0.603,
		"myInt":    123,
		"myBool":   true,
	})
}

func setName(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "changing the transaction's name")

	txn := oldfritter.FromContext(r.Context())
	txn.SetName("other-name")
}

func addAttribute(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "adding attributes")

	txn := oldfritter.FromContext(r.Context())
	txn.AddAttribute("myString", "hello")
	txn.AddAttribute("myInt", 123)
}

func addSpanAttribute(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "adding span attributes")

	txn := oldfritter.FromContext(r.Context())
	sgmt := txn.StartSegment("segment1")
	defer sgmt.End()
	sgmt.AddAttribute("mySpanString", "hello")
	sgmt.AddAttribute("mySpanInt", 123)
}

func ignore(w http.ResponseWriter, r *http.Request) {
	if coinFlip := (0 == rand.Intn(2)); coinFlip {
		txn := oldfritter.FromContext(r.Context())
		txn.Ignore()
		io.WriteString(w, "ignoring the transaction")
	} else {
		io.WriteString(w, "not ignoring the transaction")
	}
}

func segments(w http.ResponseWriter, r *http.Request) {
	txn := oldfritter.FromContext(r.Context())

	func() {
		defer txn.StartSegment("f1").End()

		func() {
			defer txn.StartSegment("f2").End()

			io.WriteString(w, "segments!")
			time.Sleep(10 * time.Millisecond)
		}()
		time.Sleep(15 * time.Millisecond)
	}()
	time.Sleep(20 * time.Millisecond)
}

func mysql(w http.ResponseWriter, r *http.Request) {
	txn := oldfritter.FromContext(r.Context())
	s := oldfritter.DatastoreSegment{
		StartTime: txn.StartSegmentNow(),
		// Product, Collection, and Operation are the most important
		// fields to populate because they are used in the breakdown
		// metrics.
		Product:    oldfritter.DatastoreMySQL,
		Collection: "users",
		Operation:  "INSERT",

		ParameterizedQuery: "INSERT INTO users (name, age) VALUES ($1, $2)",
		QueryParameters: map[string]interface{}{
			"name": "Dracula",
			"age":  439,
		},
		Host:         "mysql-server-1",
		PortPathOrID: "3306",
		DatabaseName: "my_database",
	}
	defer s.End()

	time.Sleep(20 * time.Millisecond)
	io.WriteString(w, `performing fake query "INSERT * from users"`)
}

func message(w http.ResponseWriter, r *http.Request) {
	txn := oldfritter.FromContext(r.Context())
	s := oldfritter.MessageProducerSegment{
		StartTime:       txn.StartSegmentNow(),
		Library:         "RabbitMQ",
		DestinationType: oldfritter.MessageQueue,
		DestinationName: "myQueue",
	}
	defer s.End()

	time.Sleep(20 * time.Millisecond)
	io.WriteString(w, `producing a message queue message`)
}

func external(w http.ResponseWriter, r *http.Request) {
	txn := oldfritter.FromContext(r.Context())
	req, _ := http.NewRequest("GET", "http://example.com", nil)

	// Using StartExternalSegment is recommended because it does distributed
	// tracing header setup, but if you don't have an *http.Request and
	// instead only have a url string then you can start the external
	// segment like this:
	//
	// es := oldfritter.ExternalSegment{
	// 	StartTime: txn.StartSegmentNow(),
	// 	URL:       urlString,
	// }
	//
	es := oldfritter.StartExternalSegment(txn, req)
	resp, err := http.DefaultClient.Do(req)
	es.End()

	if nil != err {
		io.WriteString(w, err.Error())
		return
	}
	defer resp.Body.Close()
	io.Copy(w, resp.Body)
}

func roundtripper(w http.ResponseWriter, r *http.Request) {
	// NewRoundTripper allows you to instrument external calls without
	// calling StartExternalSegment by modifying the http.Client's Transport
	// field.  If the Transaction parameter is nil, the RoundTripper
	// returned will look for a Transaction in the request's context (using
	// FromContext). This is recommended because it allows you to reuse the
	// same client for multiple transactions.
	client := &http.Client{}
	client.Transport = oldfritter.NewRoundTripper(client.Transport)

	request, _ := http.NewRequest("GET", "http://example.com", nil)
	// Since the transaction is already added to the inbound request's
	// context by WrapHandleFunc, we just need to copy the context from the
	// inbound request to the external request.
	request = request.WithContext(r.Context())
	// Alternatively, if you don't want to copy entire context, and instead
	// wanted just to add the transaction to the external request's context,
	// you could do that like this:
	//
	//	txn := oldfritter.FromContext(r.Context())
	//	request = oldfritter.RequestWithTransactionContext(request, txn)

	resp, err := client.Do(request)
	if nil != err {
		io.WriteString(w, err.Error())
		return
	}
	defer resp.Body.Close()
	io.Copy(w, resp.Body)
}

func async(w http.ResponseWriter, r *http.Request) {
	txn := oldfritter.FromContext(r.Context())
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(txn *oldfritter.Transaction) {
		defer wg.Done()
		defer txn.StartSegment("async").End()
		time.Sleep(100 * time.Millisecond)
	}(txn.NewGoroutine())

	segment := txn.StartSegment("wg.Wait")
	wg.Wait()
	segment.End()
	w.Write([]byte("done!"))
}

func customMetric(w http.ResponseWriter, r *http.Request) {
	txn := oldfritter.FromContext(r.Context())
	for _, vals := range r.Header {
		for _, v := range vals {
			// This custom metric will have the name
			// "Custom/HeaderLength" in the New Relic UI.
			txn.Application().RecordCustomMetric("HeaderLength", float64(len(v)))
		}
	}
	io.WriteString(w, "custom metric recorded")
}

func browser(w http.ResponseWriter, r *http.Request) {
	txn := oldfritter.FromContext(r.Context())
	hdr := txn.BrowserTimingHeader()
	// BrowserTimingHeader() will always return a header whose methods can
	// be safely called.
	if js := hdr.WithTags(); js != nil {
		w.Write(js)
	}
	io.WriteString(w, "browser header page")
}

func main() {
	app, err := oldfritter.NewApplication(
		oldfritter.ConfigAppName("Example App"),
		oldfritter.ConfigFromEnvironment(),
		oldfritter.ConfigDebugLogger(os.Stdout),
	)
	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	http.HandleFunc(oldfritter.WrapHandleFunc(app, "/", index))
	http.HandleFunc(oldfritter.WrapHandleFunc(app, "/version", versionHandler))
	http.HandleFunc(oldfritter.WrapHandleFunc(app, "/notice_error", noticeError))
	http.HandleFunc(oldfritter.WrapHandleFunc(app, "/notice_error_with_attributes", noticeErrorWithAttributes))
	http.HandleFunc(oldfritter.WrapHandleFunc(app, "/custom_event", customEvent))
	http.HandleFunc(oldfritter.WrapHandleFunc(app, "/set_name", setName))
	http.HandleFunc(oldfritter.WrapHandleFunc(app, "/add_attribute", addAttribute))
	http.HandleFunc(oldfritter.WrapHandleFunc(app, "/add_span_attribute", addSpanAttribute))
	http.HandleFunc(oldfritter.WrapHandleFunc(app, "/ignore", ignore))
	http.HandleFunc(oldfritter.WrapHandleFunc(app, "/segments", segments))
	http.HandleFunc(oldfritter.WrapHandleFunc(app, "/mysql", mysql))
	http.HandleFunc(oldfritter.WrapHandleFunc(app, "/external", external))
	http.HandleFunc(oldfritter.WrapHandleFunc(app, "/roundtripper", roundtripper))
	http.HandleFunc(oldfritter.WrapHandleFunc(app, "/custommetric", customMetric))
	http.HandleFunc(oldfritter.WrapHandleFunc(app, "/browser", browser))
	http.HandleFunc(oldfritter.WrapHandleFunc(app, "/async", async))
	http.HandleFunc(oldfritter.WrapHandleFunc(app, "/message", message))

	http.HandleFunc("/background", func(w http.ResponseWriter, req *http.Request) {
		// Transactions started without an http.Request are classified as
		// background transactions.
		txn := app.StartTransaction("background")
		defer txn.End()

		io.WriteString(w, "background transaction")
		time.Sleep(150 * time.Millisecond)
	})

	http.ListenAndServe(":8000", nil)
}
