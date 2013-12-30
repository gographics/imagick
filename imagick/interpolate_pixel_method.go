// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <wand/MagickWand.h>
*/
import "C"

type InterpolatePixelMethod int

const (
	INTERPOLATE_PIXEL_UNDEFINED        InterpolatePixelMethod = C.UndefinedInterpolatePixel
	INTERPOLATE_PIXEL_AVERAGE          InterpolatePixelMethod = C.AverageInterpolatePixel         // Average 4 nearest neighbours
	INTERPOLATE_PIXEL_BICUBIC          InterpolatePixelMethod = C.BicubicInterpolatePixel         // Catmull-Rom interpolation
	INTERPOLATE_PIXEL_BILINEAR         InterpolatePixelMethod = C.BilinearInterpolatePixel        // Triangular filter interpolation
	INTERPOLATE_PIXEL_FILTER           InterpolatePixelMethod = C.FilterInterpolatePixel          // Use resize filter - (very slow)
	INTERPOLATE_PIXEL_INTEGER          InterpolatePixelMethod = C.IntegerInterpolatePixel         // Integer (floor) interpolation
	INTERPOLATE_PIXEL_MESH             InterpolatePixelMethod = C.MeshInterpolatePixel            // Triangular mesh interpolation
	INTERPOLATE_PIXEL_NEAREST_NEIGHBOR InterpolatePixelMethod = C.NearestNeighborInterpolatePixel // Nearest neighbour only
	INTERPOLATE_PIXEL_SPLINE           InterpolatePixelMethod = C.SplineInterpolatePixel          // Cubic Spline (blurred) interpolation
	INTERPOLATE_PIXEL_AVERAGE9         InterpolatePixelMethod = C.Average9InterpolatePixel        // Average 9 nearest neighbours
	INTERPOLATE_PIXEL_AVERAGE16        InterpolatePixelMethod = C.Average16InterpolatePixel       // Average 16 nearest neighbours
	INTERPOLATE_PIXEL_BLEND            InterpolatePixelMethod = C.BlendInterpolatePixel           // blend of nearest 1, 2 or 4 pixels
	INTERPOLATE_PIXEL_BACKGROUND       InterpolatePixelMethod = C.BackgroundInterpolatePixel      // just return background color
	INTERPOLATE_PIXEL_CATROM           InterpolatePixelMethod = C.CatromInterpolatePixel          // Catmull-Rom interpolation
)
