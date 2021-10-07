// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/oldfritter/go-agent/v3/integrations/nrsqlite3"
	oldfritter "github.com/oldfritter/go-agent/v3/oldfritter"
)

func main() {
	db, err := sql.Open("nrsqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Exec("CREATE TABLE zaps ( zap_num INTEGER )")
	db.Exec("INSERT INTO zaps (zap_num) VALUES (22)")

	app, err := oldfritter.NewApplication(
		oldfritter.ConfigAppName("SQLite App"),
		oldfritter.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
		oldfritter.ConfigDebugLogger(os.Stdout),
	)
	if nil != err {
		panic(err)
	}
	app.WaitForConnection(5 * time.Second)
	txn := app.StartTransaction("sqliteQuery")

	ctx := oldfritter.NewContext(context.Background(), txn)
	row := db.QueryRowContext(ctx, "SELECT count(*) from zaps")
	var count int
	row.Scan(&count)

	txn.End()
	app.Shutdown(5 * time.Second)

	fmt.Println("number of entries in table", count)
}
