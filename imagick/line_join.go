// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type LineJoin int

const (
	LINE_JOIN_UNDEFINED LineJoin = C.UndefinedJoin
	LINE_JOIN_MITER     LineJoin = C.MiterJoin
	LINE_JOIN_ROUND     LineJoin = C.RoundJoin
	LINE_JOIN_BEVEL     LineJoin = C.BevelJoin
)
