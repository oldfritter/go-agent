// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/oldfritter/go-agent/v3/integrations/nrlogrus"
	oldfritter "github.com/oldfritter/go-agent/v3/oldfritter"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	app, err := oldfritter.NewApplication(
		oldfritter.ConfigAppName("Logrus App"),
		oldfritter.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
		nrlogrus.ConfigStandardLogger(),
	)

	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	http.HandleFunc(oldfritter.WrapHandleFunc(app, "/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello world")
	}))

	http.ListenAndServe(":8000", nil)
}
