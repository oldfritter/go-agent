// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

// Package nrlogxi supports https://github.com/mgutz/logxi.
//
// Wrap your logxi Logger using nrlogxi.New to send agent log messages through
// logxi.
package nrlogxi

import (
	log "github.com/mgutz/logxi/v1"
	"github.com/oldfritter/go-agent/v3/internal"
	oldfritter "github.com/oldfritter/go-agent/v3/oldfritter"
)

func init() { internal.TrackUsage("integration", "logging", "logxi", "v1") }

type shim struct {
	e log.Logger
}

func (l *shim) Error(msg string, context map[string]interface{}) {
	l.e.Error(msg, convert(context)...)
}
func (l *shim) Warn(msg string, context map[string]interface{}) {
	l.e.Warn(msg, convert(context)...)
}
func (l *shim) Info(msg string, context map[string]interface{}) {
	l.e.Info(msg, convert(context)...)
}
func (l *shim) Debug(msg string, context map[string]interface{}) {
	l.e.Debug(msg, convert(context)...)
}
func (l *shim) DebugEnabled() bool {
	return l.e.IsDebug()
}

func convert(c map[string]interface{}) []interface{} {
	output := make([]interface{}, 0, 2*len(c))
	for k, v := range c {
		output = append(output, k, v)
	}
	return output
}

// New returns a oldfritter.Logger which forwards agent log messages to the
// provided logxi Logger.
func New(l log.Logger) oldfritter.Logger {
	return &shim{
		e: l,
	}
}

// ConfigLogger configures the oldfritter.Application to send log messsages to the
// provided logxi logger.
func ConfigLogger(l log.Logger) oldfritter.ConfigOption {
	return oldfritter.ConfigLogger(New(l))
}
