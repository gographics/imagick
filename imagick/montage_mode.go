// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type MontageMode int

const (
	MONTAGE_MODE_UNDEFINED   MontageMode = C.UndefinedMode
	MONTAGE_MODE_FRAME       MontageMode = C.FrameMode
	MONTAGE_MODE_UNFRAME     MontageMode = C.UnframeMode
	MONTAGE_MODE_CONCATENATE MontageMode = C.ConcatenateMode
)
