// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package nrzap_test

import (
	"github.com/oldfritter/go-agent/v3/integrations/nrzap"
	"github.com/oldfritter/go-agent/v3/oldfritter"
	"go.uber.org/zap"
)

func Example() {
	// Create a new zap logger:
	z, _ := zap.NewProduction()

	oldfritter.NewApplication(
		oldfritter.ConfigAppName("Example App"),
		oldfritter.ConfigLicense("__YOUR_NEWRELIC_LICENSE_KEY__"),
		// Use nrzap to register the logger with the agent:
		nrzap.ConfigLogger(z.Named("oldfritter")),
	)
}
