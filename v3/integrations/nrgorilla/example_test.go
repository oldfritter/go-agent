// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package nrgorilla_test

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oldfritter/go-agent/v3/integrations/nrgorilla"
	oldfritter "github.com/oldfritter/go-agent/v3/oldfritter"
)

var (
	app                *oldfritter.Application
	MyCustomMiddleware mux.MiddlewareFunc
)

func makeHandler(text string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(text))
	})
}

func ExampleMiddleware() {
	r := mux.NewRouter()
	r.Use(nrgorilla.Middleware(app))

	// All handlers and custom middlewares will be instrumented.  The
	// transaction will be available in the Request's context.
	r.Use(MyCustomMiddleware)
	r.Handle("/", makeHandler("index"))

	http.ListenAndServe(":8000", r)
}

func ExampleMiddleware_specialHandlers() {
	r := mux.NewRouter()
	r.Use(nrgorilla.Middleware(app))

	// The NotFoundHandler and MethodNotAllowedHandler must be instrumented
	// separately using oldfritter.WrapHandle.  The second argument to
	// oldfritter.WrapHandle is used as the transaction name; the string returned
	// from oldfritter.WrapHandle should be ignored.
	_, r.NotFoundHandler = oldfritter.WrapHandle(app, "NotFoundHandler", makeHandler("not found"))
	_, r.MethodNotAllowedHandler = oldfritter.WrapHandle(app, "MethodNotAllowedHandler", makeHandler("method not allowed"))
}
