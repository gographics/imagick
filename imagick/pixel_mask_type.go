// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type PixelMaskType int

const (
	PIXEL_MASK_UNDEFINED PixelMaskType = C.UndefinedPixelMask
	PIXEL_MASK_READ      PixelMaskType = C.ReadPixelMask
	PIXEL_MASK_WRITE     PixelMaskType = C.WritePixelMask
	PIXEL_MASK_COMPOSITE PixelMaskType = C.CompositePixelMask
)
