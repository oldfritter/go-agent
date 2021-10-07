// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package nrstan

import (
	stan "github.com/nats-io/stan.go"
	oldfritter "github.com/oldfritter/go-agent"
	"github.com/oldfritter/go-agent/internal"
	"github.com/oldfritter/go-agent/internal/integrationsupport"
)

// StreamingSubWrapper can be used to wrap the function for STREAMING stan.Subscribe and stan.QueueSubscribe
// (https://godoc.org/github.com/nats-io/stan.go#Conn)
// If the `oldfritter.Application` parameter is non-nil, it will create a `oldfritter.Transaction` and end the transaction
// when the passed function is complete.
func StreamingSubWrapper(app oldfritter.Application, f func(msg *stan.Msg)) func(msg *stan.Msg) {
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
		txn := app.StartTransaction(namer.Name(), nil, nil)
		defer txn.End()

		integrationsupport.AddAgentAttribute(txn, internal.AttributeMessageRoutingKey, msg.MsgProto.Subject, nil)
		integrationsupport.AddAgentAttribute(txn, internal.AttributeMessageReplyTo, msg.MsgProto.Reply, nil)

		f(msg)
	}
}
