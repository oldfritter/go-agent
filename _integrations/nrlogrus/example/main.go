// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	oldfritter "github.com/oldfritter/go-agent"
	"github.com/oldfritter/go-agent/_integrations/nrlogrus"
	"github.com/sirupsen/logrus"
)

func mustGetEnv(key string) string {
	if val := os.Getenv(key); "" != val {
		return val
	}
	panic(fmt.Sprintf("environment variable %s unset", key))
}

func main() {
	cfg := oldfritter.NewConfig("Logrus App", mustGetEnv("NEW_RELIC_LICENSE_KEY"))
	logrus.SetLevel(logrus.DebugLevel)
	cfg.Logger = nrlogrus.StandardLogger()

	app, err := oldfritter.NewApplication(cfg)
	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	http.HandleFunc(oldfritter.WrapHandleFunc(app, "/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello world")
	}))

	http.ListenAndServe(":8000", nil)
}
