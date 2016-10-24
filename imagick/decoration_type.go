// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type DecorationType int

const (
	DECORATION_UNDEFINED    DecorationType = C.UndefinedDecoration
	DECORATION_NONE         DecorationType = C.NoDecoration
	DECORATION_UNDERLINE    DecorationType = C.UnderlineDecoration
	DECORATION_OVERLINE     DecorationType = C.OverlineDecoration
	DECORATION_LINE_THROUGH DecorationType = C.LineThroughDecoration
)
