// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

import (
	"fmt"
	"runtime"
	"unsafe"
)

type PixelIteratorException struct {
	kind        ExceptionType
	description string
}

func (pie *PixelIteratorException) Error() string {
	return fmt.Sprintf("%s: %s", pie.kind.String(), pie.description)
}

// Clears any exceptions associated with the iterator
func (pi *PixelIterator) clearException() bool {
	ret := 1 == C.int(C.PixelClearIteratorException(pi.pi))
	runtime.KeepAlive(pi)
	return ret
}

// Returns the kind, reason and description of any error that occurs when using
// other methods in this API
func (pi *PixelIterator) GetLastError() error {
	var et C.ExceptionType
	csdescription := C.PixelGetIteratorException(pi.pi, &et)
	defer relinquishMemory(unsafe.Pointer(csdescription))
	if ExceptionType(et) != EXCEPTION_UNDEFINED {
		pi.clearException()
		return &PixelIteratorException{ExceptionType(C.int(et)), C.GoString(csdescription)}
	}
	runtime.KeepAlive(pi)
	return nil
}

func (pi *PixelIterator) getLastErrorIfFailed(ok C.MagickBooleanType) error {
	if C.int(ok) == 0 {
		return pi.GetLastError()
	} else {
		return nil
	}
}
