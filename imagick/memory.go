// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#cgo !no_pkgconfig pkg-config: MagickWand MagickCore
#include <wand/MagickWand.h>
*/
import "C"

import (
	"unsafe"
)

// Relinquishes memory resources returned by such methods as MagickIdentifyImage(), MagickGetException(), etc
func RelinquishMemory(ptr unsafe.Pointer) {
	if ptr != nil {
		C.MagickRelinquishMemory(ptr)
	}
}

// Relinquishes memory resources, null terminated array of strings
func RelinquishMemoryCStringArray(p **C.char) {
	defer RelinquishMemory(unsafe.Pointer(p))
	for *p != nil {
		RelinquishMemory(unsafe.Pointer(*p))
		q := uintptr(unsafe.Pointer(p))
		q += unsafe.Sizeof(q)
		p = (**C.char)(unsafe.Pointer(q))
	}
}
