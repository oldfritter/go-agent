// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/oldfritter/echo"
	"github.com/oldfritter/echo/middleware"
	"github.com/oldfritter/go-agent"
	"github.com/oldfritter/go-agent/_integrations/nrecho"
)

func mustGetEnv(key string) string {
	if val := os.Getenv(key); "" != val {
		return val
	}
	panic(fmt.Sprintf("environment variable %s unset", key))
}

func getUser(c echo.Context) error {
	id := c.Param("id")

	if txn := nrecho.FromContext(c); nil != txn {
		txn.AddAttribute("userId", id)
	}

	return c.String(http.StatusOK, id)
}

func main() {
	cfg := oldfritter.NewConfig("Echo App", mustGetEnv("NEW_RELIC_LICENSE_KEY"))
	cfg.Logger = oldfritter.NewDebugLogger(os.Stdout)
	app, err := oldfritter.NewApplication(cfg)
	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	// Echo instance
	e := echo.New()

	// The New Relic Middleware should be the first middleware registered
	e.Use(nrecho.Middleware(app))

	// Routes
	e.GET("/home", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Groups
	g := e.Group("/user")
	g.Use(middleware.Gzip())
	g.GET("/:id", getUser)

	// Start server
	e.Start(":8000")
}
