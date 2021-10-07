// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/micro/go-micro"
	"github.com/oldfritter/go-agent/v3/integrations/nrmicro"
	proto "github.com/oldfritter/go-agent/v3/integrations/nrmicro/example/proto"
	oldfritter "github.com/oldfritter/go-agent/v3/oldfritter"
)

func main() {
	app, err := oldfritter.NewApplication(
		oldfritter.ConfigAppName("Micro Client"),
		oldfritter.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
		oldfritter.ConfigDebugLogger(os.Stdout),
	)
	if nil != err {
		panic(err)
	}
	err = app.WaitForConnection(10 * time.Second)
	if nil != err {
		panic(err)
	}
	defer app.Shutdown(10 * time.Second)

	txn := app.StartTransaction("client")
	defer txn.End()

	service := micro.NewService(
		// Add the New Relic wrapper to the client which will create External
		// segments for each out going call.
		micro.WrapClient(nrmicro.ClientWrapper()),
	)
	service.Init()
	ctx := oldfritter.NewContext(context.Background(), txn)
	c := proto.NewGreeterService("greeter", service.Client())

	rsp, err := c.Hello(ctx, &proto.HelloRequest{
		Name: "John",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp.Greeting)
}
