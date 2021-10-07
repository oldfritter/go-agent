// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/oldfritter/echo"
	"github.com/oldfritter/go-agent/v3/integrations/nrecho-v3"
	"github.com/oldfritter/go-agent/v3/oldfritter"
)

func getUser(c echo.Context) error {
	id := c.Param("id")

	txn := nrecho.FromContext(c)
	txn.AddAttribute("userId", id)

	return c.String(http.StatusOK, id)
}

func main() {
	app, err := oldfritter.NewApplication(
		oldfritter.ConfigAppName("Echo App"),
		oldfritter.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
		oldfritter.ConfigDebugLogger(os.Stdout),
	)
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
	g.GET("/:id", getUser)

	// Start server
	e.Start(":8000")
}
