// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/oldfritter/go-agent/v3/oldfritter"
)

func main() {
	app, err := oldfritter.NewApplication(
		oldfritter.ConfigAppName("Short Lived App"),
		oldfritter.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
		oldfritter.ConfigDebugLogger(os.Stdout),
	)
	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	// Wait for the application to connect.
	if err := app.WaitForConnection(5 * time.Second); nil != err {
		fmt.Println(err)
	}

	// Do the tasks at hand.  Perhaps record them using transactions and/or
	// custom events.
	tasks := []string{"white", "black", "red", "blue", "green", "yellow"}
	for _, task := range tasks {
		txn := app.StartTransaction("task")
		time.Sleep(10 * time.Millisecond)
		txn.End()
		app.RecordCustomEvent("task", map[string]interface{}{
			"color": task,
		})
	}

	// Shut down the application to flush data to New Relic.
	app.Shutdown(10 * time.Second)
}
