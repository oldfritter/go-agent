// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package nrnats_test

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/oldfritter/go-agent/v3/integrations/nrnats"
	"github.com/oldfritter/go-agent/v3/oldfritter"
)

func currentTransaction() *oldfritter.Transaction { return nil }

func ExampleStartPublishSegment() {
	nc, _ := nats.Connect(nats.DefaultURL)
	txn := currentTransaction()
	subject := "testing.subject"

	// Start the Publish segment
	seg := nrnats.StartPublishSegment(txn, nc, subject)
	err := nc.Publish(subject, []byte("Hello World"))
	if nil != err {
		panic(err)
	}
	// Manually end the segment
	seg.End()
}

func ExampleStartPublishSegment_defer() {
	nc, _ := nats.Connect(nats.DefaultURL)
	txn := currentTransaction()
	subject := "testing.subject"

	// Start the Publish segment and defer End till the func returns
	defer nrnats.StartPublishSegment(txn, nc, subject).End()
	m, err := nc.Request(subject, []byte("request"), time.Second)
	if nil != err {
		panic(err)
	}
	fmt.Println("Received reply message:", string(m.Data))
}
