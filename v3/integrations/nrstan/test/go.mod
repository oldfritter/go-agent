module github.com/oldfritter/go-agent/v3/integrations/nrstan/test

// This module exists to avoid a dependency on
// github.com/nats-io/nats-streaming-server in nrstan.

go 1.13

require (
	github.com/nats-io/nats-streaming-server v0.16.2
	github.com/nats-io/stan.go v0.5.0
	github.com/oldfritter/go-agent/v3 v3.4.0
	github.com/oldfritter/go-agent/v3/integrations/nrstan v0.0.0
)

replace github.com/oldfritter/go-agent/v3 => ../../../

replace github.com/oldfritter/go-agent/v3/integrations/nrstan => ../
