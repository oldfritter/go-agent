// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package nrzap

import (
	oldfritter "github.com/oldfritter/go-agent"
	"go.uber.org/zap"
)

func Example() {
	cfg := oldfritter.NewConfig("Example App", "__YOUR_NEWRELIC_LICENSE_KEY__")

	// Create a new zap logger:
	z, _ := zap.NewProduction()

	// Use nrzap to register the logger with the agent:
	cfg.Logger = Transform(z.Named("oldfritter"))

	oldfritter.NewApplication(cfg)
}
