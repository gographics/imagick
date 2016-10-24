// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type FillRule int

const (
	FILL_UNDEFINED FillRule = C.UndefinedRule
	FILL_EVEN_ODD  FillRule = C.EvenOddRule
	FILL_NON_ZERO  FillRule = C.NonZeroRule
)
