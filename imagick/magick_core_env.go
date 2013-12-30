// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <wand/MagickWand.h>
*/
import "C"

func IsCoreInstantiated() bool {
	return 1 == C.int(C.IsMagickInstantiated())
}

func CoreInitialize(path string) {
	C.MagickCoreGenesis(C.CString(path), 0)
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
