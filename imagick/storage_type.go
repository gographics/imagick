// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type StorageType int

const (
	PIXEL_UNDEFINED StorageType = C.UndefinedPixel
	PIXEL_CHAR      StorageType = C.CharPixel
	PIXEL_DOUBLE    StorageType = C.DoublePixel
	PIXEL_FLOAT     StorageType = C.FloatPixel
	PIXEL_LONGLONG  StorageType = C.LongLongPixel
	PIXEL_LONG      StorageType = C.LongPixel
	PIXEL_QUANTUM   StorageType = C.QuantumPixel
	PIXEL_SHORT     StorageType = C.ShortPixel
)
