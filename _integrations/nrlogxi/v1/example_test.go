// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package nrlogxi_test

import (
	log "github.com/mgutz/logxi/v1"
	oldfritter "github.com/oldfritter/go-agent"
	nrlogxi "github.com/oldfritter/go-agent/_integrations/nrlogxi/v1"
)

func Example() {
	cfg := oldfritter.NewConfig("Example App", "__YOUR_NEWRELIC_LICENSE_KEY__")

	// Create a new logxi logger:
	l := log.New("oldfritter")
	l.SetLevel(log.LevelInfo)

	// Use nrlogxi to register the logger with the agent:
	cfg.Logger = nrlogxi.New(l)

	oldfritter.NewApplication(cfg)
}
