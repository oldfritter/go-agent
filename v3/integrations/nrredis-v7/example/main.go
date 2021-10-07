// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"
	"os"
	"time"

	redis "github.com/go-redis/redis/v7"
	nrredis "github.com/oldfritter/go-agent/v3/integrations/nrredis-v7"
	oldfritter "github.com/oldfritter/go-agent/v3/oldfritter"
)

func main() {
	app, err := oldfritter.NewApplication(
		oldfritter.ConfigAppName("Redis App"),
		oldfritter.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
		oldfritter.ConfigDebugLogger(os.Stdout),
	)
	if nil != err {
		panic(err)
	}
	app.WaitForConnection(10 * time.Second)
	txn := app.StartTransaction("ping txn")

	opts := &redis.Options{
		Addr: "localhost:6379",
	}
	client := redis.NewClient(opts)

	//
	// Step 1:  Add a nrredis.NewHook() to your redis client.
	//
	client.AddHook(nrredis.NewHook(opts))

	//
	// Step 2: Ensure that all client calls contain a context which includes
	// the transaction.
	//
	ctx := oldfritter.NewContext(context.Background(), txn)
	pipe := client.WithContext(ctx).Pipeline()
	incr := pipe.Incr("pipeline_counter")
	pipe.Expire("pipeline_counter", time.Hour)
	_, err = pipe.Exec()
	fmt.Println(incr.Val(), err)

	txn.End()
	app.Shutdown(5 * time.Second)
}
