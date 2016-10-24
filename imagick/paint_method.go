// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type PaintMethod int

const (
	PAINT_METHOD_UNDEFINED    PaintMethod = C.UndefinedMethod
	PAINT_METHOD_POINT        PaintMethod = C.PointMethod
	PAINT_METHOD_REPLACE      PaintMethod = C.ReplaceMethod
	PAINT_METHOD_FLOODFILL    PaintMethod = C.FloodfillMethod
	PAINT_METHOD_FILLTOBORDER PaintMethod = C.FillToBorderMethod
	PAINT_METHOD_RESET        PaintMethod = C.ResetMethod
)
