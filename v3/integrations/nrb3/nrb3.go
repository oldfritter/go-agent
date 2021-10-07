// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package nrb3

import (
	"net/http"
	"time"

	"github.com/oldfritter/go-agent/v3/internal"
	oldfritter "github.com/oldfritter/go-agent/v3/oldfritter"
)

func init() { internal.TrackUsage("integration", "b3") }

// NewRoundTripper creates an `http.RoundTripper` to instrument external
// requests.  The RoundTripper returned creates an external segment and adds B3
// tracing headers to each request if and only if a `oldfritter.Transaction`
// (https://godoc.org/github.com/oldfritter/go-agent#Transaction) is found in the
// `http.Request`'s context.  It then delegates to the original RoundTripper
// provided (or http.DefaultTransport if none is provided).
func NewRoundTripper(original http.RoundTripper) http.RoundTripper {
	if nil == original {
		original = http.DefaultTransport
	}
	return &b3Transport{
		idGen:    internal.NewTraceIDGenerator(int64(time.Now().UnixNano())),
		original: original,
	}
}

// cloneRequest mimics implementation of
// https://godoc.org/github.com/google/go-github/github#BasicAuthTransport.RoundTrip
func cloneRequest(r *http.Request) *http.Request {
	// shallow copy of the struct
	r2 := new(http.Request)
	*r2 = *r
	// deep copy of the Header
	r2.Header = make(http.Header, len(r.Header))
	for k, s := range r.Header {
		r2.Header[k] = append([]string(nil), s...)
	}
	return r2
}

type b3Transport struct {
	idGen    *internal.TraceIDGenerator
	original http.RoundTripper
}

func txnSampled(txn *oldfritter.Transaction) string {
	if txn.IsSampled() {
		return "1"
	}
	return "0"
}

func addHeader(request *http.Request, key, val string) {
	if val != "" {
		request.Header.Add(key, val)
	}
}

func (t *b3Transport) RoundTrip(request *http.Request) (*http.Response, error) {
	if txn := oldfritter.FromContext(request.Context()); nil != txn {
		// The specification of http.RoundTripper requires that the request is never modified.
		request = cloneRequest(request)
		segment := &oldfritter.ExternalSegment{
			StartTime: txn.StartSegmentNow(),
			Request:   request,
		}
		defer segment.End()

		md := txn.GetTraceMetadata()
		addHeader(request, "X-B3-TraceId", md.TraceID)
		addHeader(request, "X-B3-SpanId", t.idGen.GenerateSpanID())
		addHeader(request, "X-B3-ParentSpanId", md.SpanID)
		addHeader(request, "X-B3-Sampled", txnSampled(txn))
	}

	return t.original.RoundTrip(request)
}
