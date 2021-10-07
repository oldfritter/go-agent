// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"

	oldfritter "github.com/oldfritter/go-agent"
	"github.com/oldfritter/go-agent/_integrations/nrlambda"
)

func handler(ctx context.Context) {
	// The nrlambda handler instrumentation will add the transaction to the
	// context.  Access it using oldfritter.FromContext to add additional
	// instrumentation.
	if txn := oldfritter.FromContext(ctx); nil != txn {
		txn.AddAttribute("userLevel", "gold")
		txn.Application().RecordCustomEvent("MyEvent", map[string]interface{}{
			"zip": "zap",
		})
	}
	fmt.Println("hello world")
}

func main() {
	// nrlambda.NewConfig should be used in place of oldfritter.NewConfig
	// since it sets Lambda specific configuration settings including
	// Config.ServerlessMode.Enabled.
	cfg := nrlambda.NewConfig()
	// Here is the opportunity to change configuration settings before the
	// application is created.
	app, err := oldfritter.NewApplication(cfg)
	if nil != err {
		fmt.Println("error creating app (invalid config):", err)
	}
	// nrlambda.Start should be used in place of lambda.Start.
	// nrlambda.StartHandler should be used in place of lambda.StartHandler.
	nrlambda.Start(handler, app)
}
