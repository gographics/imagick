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
// TODO(justinfx): Needs tests
type ExceptionInfo struct {
	kind        ExceptionType
	errno       int
	reason      string
	description string
}

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

func (e *ExceptionInfo) Error() string {
	return fmt.Sprintf("%s: %s", e.kind.String(), e.description)
}

func (e *ExceptionInfo) Errno() int {
	return e.errno
}

func (e *ExceptionInfo) Reason() string {
	return e.reason
}

func (e *ExceptionInfo) Description() string {
	return e.description
}
