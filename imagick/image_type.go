// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type ImageType int

const (
	IMAGE_TYPE_UNDEFINED              ImageType = C.UndefinedType
	IMAGE_TYPE_BILEVEL                ImageType = C.BilevelType
	IMAGE_TYPE_COLOR_SEPARATION       ImageType = C.ColorSeparationType
	IMAGE_TYPE_COLOR_SEPARATION_ALPHA ImageType = C.ColorSeparationAlphaType
	IMAGE_TYPE_GRAYSCALE              ImageType = C.GrayscaleType
	IMAGE_TYPE_GRAYSCALE_ALPHA        ImageType = C.GrayscaleAlphaType
	IMAGE_TYPE_OPTIMIZE               ImageType = C.OptimizeType
	IMAGE_TYPE_PALETTE                ImageType = C.PaletteType
	IMAGE_TYPE_PALETTE_ALPHA          ImageType = C.PaletteAlphaType
	IMAGE_TYPE_PALETTE_BILEVEL_ALPHA  ImageType = C.PaletteBilevelAlphaType
	IMAGE_TYPE_TRUE_COLOR             ImageType = C.TrueColorType
	IMAGE_TYPE_TRUE_COLOR_ALPHA       ImageType = C.TrueColorAlphaType
)
