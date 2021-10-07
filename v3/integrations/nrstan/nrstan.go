// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package nrstan

import (
	stan "github.com/nats-io/stan.go"
	"github.com/oldfritter/go-agent/v3/internal"
	"github.com/oldfritter/go-agent/v3/internal/integrationsupport"
	oldfritter "github.com/oldfritter/go-agent/v3/oldfritter"
)

// StreamingSubWrapper can be used to wrap the function for STREAMING stan.Subscribe and stan.QueueSubscribe
// (https://godoc.org/github.com/nats-io/stan.go#Conn)
// If the `oldfritter.Application` parameter is non-nil, it will create a `oldfritter.Transaction` and end the transaction
// when the passed function is complete.
func StreamingSubWrapper(app *oldfritter.Application, f func(msg *stan.Msg)) func(msg *stan.Msg) {
	if app == nil {
		return f
	}
	return func(msg *stan.Msg) {
		namer := internal.MessageMetricKey{
			Library:         "STAN",
			DestinationType: string(oldfritter.MessageTopic),
			DestinationName: msg.MsgProto.Subject,
			Consumer:        true,
		}
		txn := app.StartTransaction(namer.Name())
		defer txn.End()

		integrationsupport.AddAgentAttribute(txn, oldfritter.AttributeMessageRoutingKey, msg.MsgProto.Subject, nil)
		integrationsupport.AddAgentAttribute(txn, oldfritter.AttributeMessageReplyTo, msg.MsgProto.Reply, nil)

		f(msg)
	}
}
