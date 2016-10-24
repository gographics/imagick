// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type AlignType int

const (
	ALIGN_UNDEFINED AlignType = C.UndefinedAlign
	ALIGN_LEFT      AlignType = C.LeftAlign
	ALIGN_CENTER    AlignType = C.CenterAlign
	ALIGN_RIGHT     AlignType = C.RightAlign
)
