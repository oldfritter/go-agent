// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

// Package integrationsupport exists to expose functionality to integration
// packages without adding noise to the public API.
package integrationsupport

import (
	"github.com/oldfritter/go-agent/v3/internal"
	oldfritter "github.com/oldfritter/go-agent/v3/oldfritter"
)

// AddAgentAttribute allows instrumentation packages to add agent attributes.
func AddAgentAttribute(txn *oldfritter.Transaction, id string, stringVal string, otherVal interface{}) {
	if nil == txn {
		return
	}
	if aa, ok := txn.Private.(internal.AddAgentAttributer); ok {
		aa.AddAgentAttribute(id, stringVal, otherVal)
	}
}

// AddAgentSpanAttribute allows instrumentation packages to add span attributes.
func AddAgentSpanAttribute(txn *oldfritter.Transaction, key string, val string) {
	if nil == txn {
		return
	}
	internal.AddAgentSpanAttribute(txn.Private, key, val)
}

// This code below is used for testing and is based on the similar code in internal_test.go in
// the oldfritter package. That code is not exported, though, and we frequently need something similar
// for integration packages, so it is copied here.
const (
	testLicenseKey = "0123456789012345678901234567890123456789"
	SampleAppName  = "my app"
)

// ExpectApp combines Application and Expect, for use in validating data in test apps
type ExpectApp struct {
	internal.Expect
	*oldfritter.Application
}

// ConfigFullTraces enables distributed tracing and sets transaction
// trace and transaction trace segment thresholds to zero for full traces.
func ConfigFullTraces(cfg *oldfritter.Config) {
	cfg.DistributedTracer.Enabled = true
	cfg.TransactionTracer.Segments.Threshold = 0
	cfg.TransactionTracer.Threshold.IsApdexFailing = false
	cfg.TransactionTracer.Threshold.Duration = 0
}

// NewTestApp creates an ExpectApp with the given ConnectReply function and Config function
func NewTestApp(replyfn func(*internal.ConnectReply), cfgFn ...oldfritter.ConfigOption) ExpectApp {
	cfgFn = append(cfgFn,
		func(cfg *oldfritter.Config) {
			// Prevent spawning app goroutines in tests.
			if !cfg.ServerlessMode.Enabled {
				cfg.Enabled = false
			}
		},
		oldfritter.ConfigAppName(SampleAppName),
		oldfritter.ConfigLicense(testLicenseKey),
	)

	app, err := oldfritter.NewApplication(cfgFn...)
	if nil != err {
		panic(err)
	}

	internal.HarvestTesting(app.Private, replyfn)

	return ExpectApp{
		Expect:      app.Private.(internal.Expect),
		Application: app,
	}
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
	reply.SetSampleEverything()
}
