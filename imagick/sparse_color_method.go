// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type SparseColorMethod int

const (
	INTERPOLATE_UNDEFINED_COLOR   SparseColorMethod = C.UndefinedColorInterpolate
	INTERPOLATE_BARYCENTRIC_COLOR SparseColorMethod = C.BarycentricColorInterpolate
	INTERPOLATE_BILINEAR_COLOR    SparseColorMethod = C.BilinearColorInterpolate
	INTERPOLATE_POLYNOMIAL_COLOR  SparseColorMethod = C.PolynomialColorInterpolate
	INTERPOLATE_SHEPARDS_COLOR    SparseColorMethod = C.ShepardsColorInterpolate
	INTERPOLATE_VORONOI_COLOR     SparseColorMethod = C.VoronoiColorInterpolate
	INTERPOLATE_INVERSE_COLOR     SparseColorMethod = C.InverseColorInterpolate
)
