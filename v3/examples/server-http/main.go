// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

// An application that illustrates Distributed Tracing or Cross Application
// Tracing when using http.Server or similar frameworks.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	oldfritter "github.com/oldfritter/go-agent/v3/oldfritter"
)

type handler struct {
	App *oldfritter.Application
}

func (h *handler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	// The call to StartTransaction must include the response writer and the
	// request.
	txn := h.App.StartTransaction("server-txn")
	defer txn.End()

	writer = txn.SetWebResponse(writer)
	txn.SetWebRequestHTTP(req)

	if req.URL.String() == "/segments" {
		defer txn.StartSegment("f1").End()

		func() {
			defer txn.StartSegment("f2").End()

			io.WriteString(writer, "segments!")
			time.Sleep(10 * time.Millisecond)
		}()
		time.Sleep(10 * time.Millisecond)
	} else {
		// Transaction.WriteHeader has to be used instead of invoking
		// WriteHeader on the response writer.
		writer.WriteHeader(http.StatusNotFound)
	}
}

func makeApplication() (*oldfritter.Application, error) {
	app, err := oldfritter.NewApplication(
		oldfritter.ConfigAppName("HTTP Server App"),
		oldfritter.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
		oldfritter.ConfigDebugLogger(os.Stdout),
		oldfritter.ConfigDistributedTracerEnabled(true),
	)
	if nil != err {
		return nil, err
	}

	// Wait for the application to connect.
	if err = app.WaitForConnection(5 * time.Second); nil != err {
		return nil, err
	}

	return app, nil
}

func main() {

	app, err := makeApplication()
	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	server := http.Server{
		Addr:    ":8000",
		Handler: &handler{App: app},
	}

	server.ListenAndServe()
}
