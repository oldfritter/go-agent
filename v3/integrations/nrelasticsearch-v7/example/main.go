// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	elasticsearch "github.com/elastic/go-elasticsearch/v7"
	nrelasticsearch "github.com/oldfritter/go-agent/v3/integrations/nrelasticsearch-v7"
	"github.com/oldfritter/go-agent/v3/oldfritter"
)

func main() {
	// Step 1: Use nrelasticsearch.NewRoundTripper to assign the
	// elasticsearch.Config's Transport field.
	cfg := elasticsearch.Config{
		Transport: nrelasticsearch.NewRoundTripper(nil),
	}

	app, err := oldfritter.NewApplication(
		oldfritter.ConfigAppName("Elastic App"),
		oldfritter.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
		oldfritter.ConfigDebugLogger(os.Stdout),
	)
	if nil != err {
		panic(err)
	}
	app.WaitForConnection(5 * time.Second)
	txn := app.StartTransaction("elastic")

	// Step 2: Ensure that all calls using the elasticsearch client have a
	// context which includes the oldfritter.Transaction.
	ctx := oldfritter.NewContext(context.Background(), txn)
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	res, err := es.Info(es.Info.WithContext(ctx))
	if err != nil {
		panic(err)
	}
	if res.IsError() {
		panic(err)
	}
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		panic(err)
	}
	fmt.Println("ELASTIC SEARCH INFO", elasticsearch.Version, r["version"].(map[string]interface{})["number"])

	txn.End()
	app.Shutdown(5 * time.Second)
}
