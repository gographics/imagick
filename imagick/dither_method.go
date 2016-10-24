// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type DitherMethod int

const (
	DITHER_METHOD_UNDEFINED       DitherMethod = C.UndefinedDitherMethod
	DITHER_METHOD_NO              DitherMethod = C.NoDitherMethod
	DITHER_METHOD_RIEMERSMA       DitherMethod = C.RiemersmaDitherMethod
	DITHER_METHOD_FLOYD_STEINBERG DitherMethod = C.FloydSteinbergDitherMethod
)
