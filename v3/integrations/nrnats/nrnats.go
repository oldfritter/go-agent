// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package nrnats

import (
	"strings"

	nats "github.com/nats-io/nats.go"
	"github.com/oldfritter/go-agent/v3/internal"
	"github.com/oldfritter/go-agent/v3/internal/integrationsupport"
	oldfritter "github.com/oldfritter/go-agent/v3/oldfritter"
)

// StartPublishSegment creates and starts a `oldfritter.MessageProducerSegment`
// (https://godoc.org/github.com/oldfritter/go-agent#MessageProducerSegment) for NATS
// publishers.  Call this function before calling any method that publishes or
// responds to a NATS message.  Call `End()`
// (https://godoc.org/github.com/oldfritter/go-agent#MessageProducerSegment.End) on the
// returned oldfritter.MessageProducerSegment when the publish is complete.  The
// `oldfritter.Transaction` and `nats.Conn` parameters are required.  The subject
// parameter is the subject of the publish call and is used in metric and span
// names.
func StartPublishSegment(txn *oldfritter.Transaction, nc *nats.Conn, subject string) *oldfritter.MessageProducerSegment {
	if nil == txn {
		return nil
	}
	if nil == nc {
		return nil
	}
	return &oldfritter.MessageProducerSegment{
		StartTime:            txn.StartSegmentNow(),
		Library:              "NATS",
		DestinationType:      oldfritter.MessageTopic,
		DestinationName:      subject,
		DestinationTemporary: strings.HasPrefix(subject, "_INBOX"),
	}
}

// SubWrapper can be used to wrap the function for nats.Subscribe (https://godoc.org/github.com/nats-io/go-nats#Conn.Subscribe
// or https://godoc.org/github.com/nats-io/go-nats#EncodedConn.Subscribe)
// and nats.QueueSubscribe (https://godoc.org/github.com/nats-io/go-nats#Conn.QueueSubscribe or
// https://godoc.org/github.com/nats-io/go-nats#EncodedConn.QueueSubscribe)
// If the `oldfritter.Application` parameter is non-nil, it will create a `oldfritter.Transaction` and end the transaction
// when the passed function is complete.
func SubWrapper(app *oldfritter.Application, f func(msg *nats.Msg)) func(msg *nats.Msg) {
	if app == nil {
		return f
	}
	return func(msg *nats.Msg) {
		namer := internal.MessageMetricKey{
			Library:         "NATS",
			DestinationType: string(oldfritter.MessageTopic),
			DestinationName: msg.Subject,
			Consumer:        true,
		}
		txn := app.StartTransaction(namer.Name())
		defer txn.End()

		integrationsupport.AddAgentAttribute(txn, oldfritter.AttributeMessageRoutingKey, msg.Sub.Subject, nil)
		integrationsupport.AddAgentAttribute(txn, oldfritter.AttributeMessageQueueName, msg.Sub.Queue, nil)
		integrationsupport.AddAgentAttribute(txn, oldfritter.AttributeMessageReplyTo, msg.Reply, nil)

		f(msg)
	}
}
