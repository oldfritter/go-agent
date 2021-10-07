// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"

	"github.com/oldfritter/go-agent/v3/integrations/nrlambda"
	oldfritter "github.com/oldfritter/go-agent/v3/oldfritter"
)

func handler(ctx context.Context) {
	// The nrlambda handler instrumentation will add the transaction to the
	// context.  Access it using oldfritter.FromContext to add additional
	// instrumentation.
	txn := oldfritter.FromContext(ctx)
	txn.AddAttribute("userLevel", "gold")
	txn.Application().RecordCustomEvent("MyEvent", map[string]interface{}{
		"zip": "zap",
	})

	fmt.Println("hello world")
}

func main() {
	// Pass nrlambda.ConfigOption() into oldfritter.NewApplication to set
	// Lambda specific configuration settings including
	// Config.ServerlessMode.Enabled.
	app, err := oldfritter.NewApplication(nrlambda.ConfigOption())
	if nil != err {
		fmt.Println("error creating app (invalid config):", err)
	}
	// nrlambda.Start should be used in place of lambda.Start.
	// nrlambda.StartHandler should be used in place of lambda.StartHandler.
	nrlambda.Start(handler, app)
}
