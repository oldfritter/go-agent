// Copyright 2020, 2021 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

//
// Example of using nrpgx to instrument a Postgres database application
// using the jackc/pgx driver with database/sql.
//
// To run this example, be sure the environment variable NEW_RELIC_LICENSE_KEY
// is set to your license key. Postgres must be running on the default port
// 5432 on localhost, and have a password "docker". An easy (albeit insecure)
// way to test this is to issue the following command to run a postgres database
// in a docker container:
//    docker run --rm -e POSTGRES_PASSWORD=docker -p 5432:5432 postgres
//
// Run that in the background or in a separate window, and then run this program
// to access that database.
//

package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/oldfritter/go-agent/v3/integrations/nrpgx"
	"github.com/oldfritter/go-agent/v3/oldfritter"
)

func main() {
	// docker run --rm -e POSTGRES_PASSWORD=docker -p 5432:5432 postgres
	db, err := sql.Open("nrpgx", "host=localhost port=5432 user=postgres dbname=postgres password=docker sslmode=disable")
	if err != nil {
		panic(err)
	}

	app, err := oldfritter.NewApplication(
		oldfritter.ConfigAppName("PostgreSQL App"),
		oldfritter.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
		oldfritter.ConfigDebugLogger(os.Stdout),
	)
	if err != nil {
		panic(err)
	}
	//
	// N.B.: We do not recommend using app.WaitForConnection in production code.
	//
	app.WaitForConnection(5 * time.Second)
	txn := app.StartTransaction("postgresQuery")

	ctx := oldfritter.NewContext(context.Background(), txn)
	row := db.QueryRowContext(ctx, "SELECT count(*) FROM pg_catalog.pg_tables")
	var count int
	row.Scan(&count)

	txn.End()
	app.Shutdown(5 * time.Second)

	fmt.Println("number of entries in pg_catalog.pg_tables", count)
}
