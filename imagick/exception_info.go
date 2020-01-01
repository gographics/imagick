// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

import "fmt"

// ExceptionInfo is an error type returned by certain
// New* API calls.
type ExceptionInfo struct {
	kind        ExceptionType
	errno       int
	reason      string
	description string
}

// Create a new ExceptionInfo wrapper around a C ExceptionInfo ptr
func newExceptionInfo(errInfo *C.ExceptionInfo) *ExceptionInfo {
	if errInfo == nil {
		return nil
	}

	return &ExceptionInfo{
		kind:        ExceptionType(errInfo.severity),
		errno:       int(errInfo.error_number),
		reason:      C.GoString(errInfo.reason),
		description: C.GoString(errInfo.description),
	}
}

// Check if a given C ExceptionInfo ptr is an error.
// Returns a valid ExceptionInfo if there was an error,
// otherwise returns nil
func checkExceptionInfo(errInfo *C.ExceptionInfo) *ExceptionInfo {
	if errInfo != nil && errInfo.error_number != 0 {
		return newExceptionInfo(errInfo)
	}

	return nil
}

func (e *ExceptionInfo) Error() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("%s (%d): %s %s",
		e.kind.String(), e.Errno(), e.Reason(), e.Description())
}

// Errno returns the ExceptionInfo error number (non-zero if error)
func (e *ExceptionInfo) Errno() int {
	if e == nil {
		return 0
	}
	return e.errno
}

// Reason returns the string reason for the Exception
func (e *ExceptionInfo) Reason() string {
	if e == nil {
		return ""
	}
	return e.reason
}

// Description returns the string description for the Exception
func (e *ExceptionInfo) Description() string {
	if e == nil {
		return ""
	}
	return e.description
}
