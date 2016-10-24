// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type ResolutionType int

const (
	RESOLUTION_UNDEFINED             ResolutionType = C.UndefinedResolution
	RESOLUTION_PIXELS_PER_INCH       ResolutionType = C.PixelsPerInchResolution
	RESOLUTION_PIXELS_PER_CENTIMETER ResolutionType = C.PixelsPerCentimeterResolution
)
