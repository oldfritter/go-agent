// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"
	"os"
	"time"

	oldfritter "github.com/oldfritter/go-agent"
	"github.com/oldfritter/go-agent/_integrations/logcontext/nrlogrusplugin"
	"github.com/sirupsen/logrus"
)

func mustGetEnv(key string) string {
	if val := os.Getenv(key); "" != val {
		return val
	}
	panic(fmt.Sprintf("environment variable %s unset", key))
}

func doFunction2(txn oldfritter.Transaction, e *logrus.Entry) {
	defer oldfritter.StartSegment(txn, "doFunction2").End()
	e.Error("In doFunction2")
}

func doFunction1(txn oldfritter.Transaction, e *logrus.Entry) {
	defer oldfritter.StartSegment(txn, "doFunction1").End()
	e.Trace("In doFunction1")
	doFunction2(txn, e)
}

func main() {
	log := logrus.New()
	// To enable New Relic log decoration, use the
	// nrlogrusplugin.ContextFormatter{}
	log.SetFormatter(nrlogrusplugin.ContextFormatter{})
	log.SetLevel(logrus.TraceLevel)

	log.Debug("Logger created")

	cfg := oldfritter.NewConfig("Logrus Log Decoration", mustGetEnv("NEW_RELIC_LICENSE_KEY"))
	cfg.DistributedTracer.Enabled = true
	cfg.CrossApplicationTracer.Enabled = false

	app, err := oldfritter.NewApplication(cfg)
	if nil != err {
		log.Panic("Failed to create application", err)
	}

	log.Debug("Application created, waiting for connection")

	err = app.WaitForConnection(10 * time.Second)
	if nil != err {
		log.Panic("Failed to connect application", err)
	}
	log.Info("Application connected")
	defer app.Shutdown(10 * time.Second)

	log.Debug("Starting transaction now")
	txn := app.StartTransaction("main", nil, nil)

	// Add the transaction context to the logger. Only once this happens will
	// the logs be properly decorated with all required fields.
	e := log.WithContext(oldfritter.NewContext(context.Background(), txn))

	doFunction1(txn, e)

	e.Info("Ending transaction")
	txn.End()
}
