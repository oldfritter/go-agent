// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package nrlogxi_test

import (
	log "github.com/mgutz/logxi/v1"
	nrlogxi "github.com/oldfritter/go-agent/v3/integrations/nrlogxi"
	oldfritter "github.com/oldfritter/go-agent/v3/oldfritter"
)

func Example() {
	// Create a new logxi logger:
	l := log.New("oldfritter")
	l.SetLevel(log.LevelInfo)

	oldfritter.NewApplication(
		oldfritter.ConfigAppName("Example App"),
		oldfritter.ConfigLicense("__YOUR_NEWRELIC_LICENSE_KEY__"),
		// Use nrlogxi to register the logger with the agent:
		nrlogxi.ConfigLogger(l),
	)
}
