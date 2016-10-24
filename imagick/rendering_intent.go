// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type RenderingIntent int

const (
	RENDERING_INTENT_UNDEFINED  RenderingIntent = C.UndefinedIntent
	RENDERING_INTENT_SATURATION RenderingIntent = C.SaturationIntent
	RENDERING_INTENT_PERCEPTUAL RenderingIntent = C.PerceptualIntent
	RENDERING_INTENT_ABSOLUTE   RenderingIntent = C.AbsoluteIntent
	RENDERING_INTENT_RELATIVE   RenderingIntent = C.RelativeIntent
)
