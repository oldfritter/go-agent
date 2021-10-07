// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"os"
	"time"

	"github.com/oldfritter/go-agent/v3/integrations/nrmongo"
	oldfritter "github.com/oldfritter/go-agent/v3/oldfritter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	app, err := oldfritter.NewApplication(
		oldfritter.ConfigAppName("Basic Mongo App"),
		oldfritter.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
		oldfritter.ConfigDebugLogger(os.Stdout),
	)
	if nil != err {
		panic(err)
	}
	app.WaitForConnection(10 * time.Second)

	// If you have another CommandMonitor, you can pass it to NewCommandMonitor and it will get called along
	// with the NR monitor
	nrMon := nrmongo.NewCommandMonitor(nil)
	ctx := context.Background()

	// nrMon must be added after any other monitors are added, as previous options get overwritten.
	// This example assumes Mongo is running locally on port 27017
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017").SetMonitor(nrMon))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	txn := app.StartTransaction("Mongo txn")
	// Make sure to add the oldfritter.Transaction to the context
	nrCtx := oldfritter.NewContext(context.Background(), txn)
	collection := client.Database("testing").Collection("numbers")
	_, err = collection.InsertOne(nrCtx, bson.M{"name": "exampleName", "value": "exampleValue"})
	if err != nil {
		panic(err)
	}
	txn.End()
	app.Shutdown(10 * time.Second)

}
