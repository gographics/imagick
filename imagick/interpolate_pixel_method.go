// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type PixelInterpolateMethod int

const (
	INTERPOLATE_PIXEL_UNDEFINED           PixelInterpolateMethod = C.UndefinedInterpolatePixel
	INTERPOLATE_PIXEL_AVERAGE             PixelInterpolateMethod = C.AverageInterpolatePixel    // Average 4 nearest neighbours
	INTERPOLATE_PIXEL_AVERAGE16           PixelInterpolateMethod = C.Average16InterpolatePixel  // Average 16 nearest neighbours
	INTERPOLATE_PIXEL_AVERAGE9            PixelInterpolateMethod = C.Average9InterpolatePixel   // Average 9 nearest neighbours
	INTERPOLATE_PIXEL_BACKGROUND          PixelInterpolateMethod = C.BackgroundInterpolatePixel // just return background color
	INTERPOLATE_PIXEL_BILINEAR            PixelInterpolateMethod = C.BilinearInterpolatePixel   // Triangular filter interpolation
	INTERPOLATE_PIXEL_BLEND               PixelInterpolateMethod = C.BlendInterpolatePixel      // blend of nearest 1, 2 or 4 pixels
	INTERPOLATE_PIXEL_CATROM              PixelInterpolateMethod = C.CatromInterpolatePixel     // Catmull-Rom interpolation
	INTERPOLATE_PIXEL_INTEGER             PixelInterpolateMethod = C.IntegerInterpolatePixel    // Integer (floor) interpolation
	INTERPOLATE_PIXEL_MESH                PixelInterpolateMethod = C.MeshInterpolatePixel       // Triangular mesh interpolation
	INTERPOLATE_PIXEL_NEAREST_INTERPOLATE PixelInterpolateMethod = C.NearestInterpolatePixel    // Nearest neighbour only
	INTERPOLATE_PIXEL_SPLINE              PixelInterpolateMethod = C.SplineInterpolatePixel     // Cubic Spline (blurred) interpolation
)
