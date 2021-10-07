// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

// Package nrpkgerrors introduces support for https://github.com/pkg/errors.
//
// This package improves the class and stack-trace fields of pkg/error errors
// when they are recorded with Transaction.NoticeError.
//
package nrpkgerrors

import (
	"fmt"

	oldfritter "github.com/oldfritter/go-agent"
	"github.com/oldfritter/go-agent/internal"
	"github.com/pkg/errors"
)

func init() { internal.TrackUsage("integration", "pkg-errors") }

type nrpkgerror struct {
	error
}

// stackTracer is an error that also knows about its StackTrace.
// All wrapped errors from github.com/pkg/errors implement this interface.
type stackTracer interface {
	StackTrace() errors.StackTrace
}

func deepestStackTrace(err error) errors.StackTrace {
	var last stackTracer
	for err != nil {
		if err, ok := err.(stackTracer); ok {
			last = err
		}
		cause, ok := err.(interface {
			Cause() error
		})
		if !ok {
			break
		}
		err = cause.Cause()
	}

	if last == nil {
		return nil
	}
	return last.StackTrace()
}

func transformStackTrace(orig errors.StackTrace) []uintptr {
	st := make([]uintptr, len(orig))
	for i, frame := range orig {
		st[i] = uintptr(frame)
	}
	return st
}

func (e nrpkgerror) StackTrace() []uintptr {
	st := deepestStackTrace(e.error)
	if nil == st {
		return nil
	}
	return transformStackTrace(st)
}

func (e nrpkgerror) ErrorClass() string {
	if ec, ok := e.error.(oldfritter.ErrorClasser); ok {
		return ec.ErrorClass()
	}
	cause := errors.Cause(e.error)
	if ec, ok := cause.(oldfritter.ErrorClasser); ok {
		return ec.ErrorClass()
	}
	return fmt.Sprintf("%T", cause)
}

// Wrap wraps a pkg/errors error so that when noticed by
// oldfritter.Transaction.NoticeError it gives an improved stacktrace and class
// type.
func Wrap(e error) error {
	return nrpkgerror{e}
}
