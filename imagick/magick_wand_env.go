// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <wand/MagickWand.h>
*/
import "C"
import (
	"sync"
)

var (
	initialized   bool
	initOnce      sync.Once
	terminateOnce sync.Once
)

// Inicializes the MagickWand environment
func Initialize() {
	initOnce.Do(func() {
		if initialized {
			return
		}
		C.MagickWandGenesis()
		initialized = true
	})
}

// Terminates the MagickWand environment
func Terminate() {
	if initialized {
		terminateOnce.Do(func() {
			C.MagickWandTerminus()
			initialized = false
		})
	}
}
