// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#cgo !no_pkgconfig pkg-config: MagickWand MagickCore
#include <MagickWand/MagickWand.h>
*/
import "C"

import (
	"unsafe"
)

// relinquishes memory resources returned by such methods as MagickIdentifyImage(), MagickGetException(), etc
func relinquishMemory(ptr unsafe.Pointer) {
	if ptr != nil {
		C.MagickRelinquishMemory(ptr)
	}
}

// relinquishes memory resources, null terminated array of strings
func relinquishMemoryCStringArray(p **C.char) {
	defer relinquishMemory(unsafe.Pointer(p))
	for *p != nil {
		relinquishMemory(unsafe.Pointer(*p))
		q := uintptr(unsafe.Pointer(p))
		q += unsafe.Sizeof(q)
		p = (**C.char)(unsafe.Pointer(q))
	}
}
