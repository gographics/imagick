// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type LineCap int

const (
	LINE_CAP_UNDEFINED LineCap = C.UndefinedCap
	LINE_CAP_BUTT      LineCap = C.ButtCap
	LINE_CAP_ROUND     LineCap = C.RoundCap
	LINE_CAP_SQUARE    LineCap = C.SquareCap
)
