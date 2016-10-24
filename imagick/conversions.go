// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"
import "unsafe"

// Convert a boolean to a integer
func b2i(boolean bool) C.MagickBooleanType {
	if boolean {
		return C.MagickBooleanType(1)
	}
	return C.MagickBooleanType(0)
}

func cStringArrayToStringSlice(p **C.char) []string {
	var strings []string
	q := uintptr(unsafe.Pointer(p))
	for {
		p = (**C.char)(unsafe.Pointer(q))
		if *p == nil {
			break
		}
		strings = append(strings, C.GoString(*p))
		q += unsafe.Sizeof(q)
	}
	return strings
}

func sizedCStringArrayToStringSlice(p **C.char, num C.size_t) []string {
	var strings []string
	q := uintptr(unsafe.Pointer(p))
	for i := 0; i < int(num); i++ {
		p = (**C.char)(unsafe.Pointer(q))
		if *p == nil {
			break
		}
		strings = append(strings, C.GoString(*p))
		q += unsafe.Sizeof(q)
	}
	return strings
}

func sizedDoubleArrayToFloat64Slice(p *C.double, num C.size_t) []float64 {
	var nums []float64
	q := uintptr(unsafe.Pointer(p))
	for i := 0; i < int(num); i++ {
		p = (*C.double)(unsafe.Pointer(q))
		nums = append(nums, float64(*p))
		q += unsafe.Sizeof(q)
	}
	return nums
}

func ConvertRGBToHSL(qr, qg, qb float64) (h, s, l float64) {
	var dh, ds, dl C.double
	C.ConvertRGBToHSL(C.double(qr), C.double(qg), C.double(qb), &dh, &ds, &dl)
	return float64(dh), float64(ds), float64(dl)
}

func ConvertHSLToRGB(fh, fs, fl float64) (r, g, b float64) {
	var qr, qg, qb C.double
	C.ConvertHSLToRGB(C.double(fh), C.double(fs), C.double(fl), &qr, &qg, &qb)
	return float64(qr), float64(qg), float64(qb)
}
