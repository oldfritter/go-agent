// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

// Package integrationsupport exists to expose functionality to integration
// packages without adding noise to the public API.
package integrationsupport

import (
	oldfritter "github.com/oldfritter/go-agent"
	"github.com/oldfritter/go-agent/internal"
)

// AddAgentAttribute allows instrumentation packages to add agent attributes.
func AddAgentAttribute(txn oldfritter.Transaction, id internal.AgentAttributeID, stringVal string, otherVal interface{}) {
	if aa, ok := txn.(internal.AddAgentAttributer); ok {
		aa.AddAgentAttribute(id, stringVal, otherVal)
	}
}

// AddAgentSpanAttribute allows instrumentation packages to add span attributes.
func AddAgentSpanAttribute(txn oldfritter.Transaction, key internal.SpanAttribute, val string) {
	internal.AddAgentSpanAttribute(txn, key, val)
}

// This code below is used for testing and is based on the similar code in internal_test.go in
// the oldfritter package. That code is not exported, though, and we frequently need something similar
// for integration packages, so it is copied here.
const (
	testLicenseKey = "0123456789012345678901234567890123456789"
	SampleAppName  = "my app"
)

// ExpectApp combines Application and Expect, for use in validating data in test apps
type ExpectApp interface {
	internal.Expect
	oldfritter.Application
}

// NewTestApp creates an ExpectApp with the given ConnectReply function and Config function
func NewTestApp(replyfn func(*internal.ConnectReply), cfgFn func(*oldfritter.Config)) ExpectApp {

	cfg := oldfritter.NewConfig(SampleAppName, testLicenseKey)

	if nil != cfgFn {
		cfgFn(&cfg)
	}

	// Prevent spawning app goroutines in tests.
	if !cfg.ServerlessMode.Enabled {
		cfg.Enabled = false
	}

	app, err := oldfritter.NewApplication(cfg)
	if nil != err {
		panic(err)
	}

	internal.HarvestTesting(app, replyfn)

	return app.(ExpectApp)
}

// NewBasicTestApp creates an ExpectApp with the standard testing connect reply function and config
func NewBasicTestApp() ExpectApp {
	return NewTestApp(nil, BasicConfigFn)
}

// BasicConfigFn is a default config function to be used when no special settings are needed for a test app
var BasicConfigFn = func(cfg *oldfritter.Config) {
	cfg.Enabled = false
}

// DTEnabledCfgFn is a reusable Config function that sets Distributed Tracing to enabled
var DTEnabledCfgFn = func(cfg *oldfritter.Config) {
	cfg.Enabled = false
	cfg.DistributedTracer.Enabled = true
}

// SampleEverythingReplyFn is a reusable ConnectReply function that samples everything
var SampleEverythingReplyFn = func(reply *internal.ConnectReply) {
	reply.AdaptiveSampler = internal.SampleEverything{}
}
