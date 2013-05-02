package imagick

/*
#cgo pkg-config: MagickWand
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
	defer C.free(unsafe.Pointer(csdescription))
	if ExceptionType(et) != UndefinedException {
		pi.clearException()
		return &PixelIteratorException{ExceptionType(C.int(et)), C.GoString(csdescription)}
	}
	return nil
}
