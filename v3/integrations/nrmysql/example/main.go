// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/oldfritter/go-agent/v3/integrations/nrmysql"
	"github.com/oldfritter/go-agent/v3/oldfritter"
)

func main() {
	// Set up a local mysql docker container with:
	// docker run -it -p 3306:3306 --net "bridge" -e MYSQL_ALLOW_EMPTY_PASSWORD=true mysql

	db, err := sql.Open("nrmysql", "root@/information_schema")
	if nil != err {
		panic(err)
	}

	app, err := oldfritter.NewApplication(
		oldfritter.ConfigAppName("MySQL App"),
		oldfritter.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
		oldfritter.ConfigDebugLogger(os.Stdout),
	)
	if nil != err {
		panic(err)
	}
	app.WaitForConnection(5 * time.Second)
	txn := app.StartTransaction("mysqlQuery")

	ctx := oldfritter.NewContext(context.Background(), txn)
	row := db.QueryRowContext(ctx, "SELECT count(*) from tables")
	var count int
	row.Scan(&count)

	txn.End()
	app.Shutdown(5 * time.Second)

	fmt.Println("number of tables in information_schema", count)
}
