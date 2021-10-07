// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"log"
	"net/http"
	"os"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/oldfritter/go-agent/v3/integrations/nrgraphgophers"
	oldfritter "github.com/oldfritter/go-agent/v3/oldfritter"
)

type query struct{}

func (*query) Hello() string { return "Hello, world!" }

func main() {
	// First create your New Relic Application:
	app, err := oldfritter.NewApplication(
		oldfritter.ConfigAppName("GraphQL App"),
		oldfritter.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
		oldfritter.ConfigDebugLogger(os.Stdout),
	)
	if nil != err {
		panic(err)
	}

	s := `type Query { hello: String! }`

	// Then add a graphql.Tracer(nrgraphgophers.NewTracer()) option to your
	// schema parsing to get field and query segment instrumentation:
	opt := graphql.Tracer(nrgraphgophers.NewTracer())
	schema := graphql.MustParseSchema(s, &query{}, opt)

	// Finally, instrument your request handler using oldfritter.WrapHandle
	// to create transactions for requests:
	http.Handle(oldfritter.WrapHandle(app, "/graphql", &relay.Handler{Schema: schema}))

	// To test, run:
	// curl -X POST -d '{"query": "query HelloOperation { hello }" }' localhost:8000/graphql
	log.Fatal(http.ListenAndServe(":8000", nil))
}
