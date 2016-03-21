// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <wand/MagickWand.h>
*/
import "C"
import "sync"

var (
	initMutex sync.Mutex
	initialized bool
)

// Inicializes the MagickWand environment
func Initialize() {
	initMutex.Lock()
	defer initMutex.Unlock()
	if initialized {
		return
	}
	C.MagickWandGenesis()
	initialized = true
}

// Terminates the MagickWand environment
func Terminate() {
	initMutex.Lock()
	defer initMutex.Unlock()
	if initialized {
		C.MagickWandTerminus()
		initialized = false
	}
}
