// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

import (
	"unsafe"
)

func IsCoreInstantiated() bool {
	return 1 == C.int(C.IsMagickCoreInstantiated())
}

func CoreInitialize(path string) {
	cspath := C.CString(path)
	defer C.free(unsafe.Pointer(cspath))
	C.MagickCoreGenesis(cspath, 0)
}

func CoreTerminate() {
	C.MagickCoreTerminus()
}

func GetPrecision() int {
	return int(C.GetMagickPrecision())
}

func SetPrecision(precision int) {
	C.SetMagickPrecision(C.int(precision))
}
