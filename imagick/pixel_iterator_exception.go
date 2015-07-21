// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
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
	return 1 == C.int(C.PixelClearIteratorException(pi.pi))
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
	return nil
}
