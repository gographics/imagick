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
	"unsafe"
)

type PixelWandException struct {
	kind        ExceptionType
	description string
}

func (pwe *PixelWandException) Error() string {
	return fmt.Sprintf("%s: %s", pwe.kind.String(), pwe.description)
}

// Clears any exceptions associated with the wand
func (pw *PixelWand) clearException() bool {
	return 1 == C.int(C.PixelClearException(pw.pw))
}

// Returns the kind, reason and description of any error that occurs when using other methods in this API
func (pw *PixelWand) GetLastError() error {
	var et C.ExceptionType
	csdescription := C.PixelGetException(pw.pw, &et)
	defer relinquishMemory(unsafe.Pointer(csdescription))
	if ExceptionType(et) != EXCEPTION_UNDEFINED {
		pw.clearException()
		return &PixelWandException{ExceptionType(C.int(et)), C.GoString(csdescription)}
	}
	return nil
}
