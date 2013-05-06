// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

var (
	initialized bool
)

// Inicializes the MagickWand environment
func Initialize() {
	if initialized {
		return
	}
	C.MagickWandGenesis()
	initialized = true
}

// Terminates the MagickWand environment
func Terminate() {
	if initialized {
		C.MagickWandTerminus()
		initialized = false
	}
}
