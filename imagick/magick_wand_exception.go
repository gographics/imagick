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

type MagickWandException struct {
	kind        ExceptionType
	description string
}

func (mwe *MagickWandException) Error() string {
	return fmt.Sprintf("%s: %s", mwe.kind.String(), mwe.description)
}

// Clears any exceptions associated with the wand
func (mw *MagickWand) clearException() bool {
	return 1 == C.int(C.MagickClearException(mw.mw))
}

// Returns the kind, reason and description of any error that occurs when using other methods in this API
func (mw *MagickWand) GetLastError() error {
	var et C.ExceptionType
	csdescription := C.MagickGetException(mw.mw, &et)
	defer mw.relinquishMemory(unsafe.Pointer(csdescription))
	if ExceptionType(et) != EXCEPTION_UNDEFINED {
		mw.clearException()
		return &MagickWandException{ExceptionType(C.int(et)), C.GoString(csdescription)}
	}
	return nil
}
