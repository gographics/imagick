// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type FilterType int

const (
	FILTER_UNDEFINED      FilterType = C.UndefinedFilter
	FILTER_POINT          FilterType = C.PointFilter
	FILTER_BOX            FilterType = C.BoxFilter
	FILTER_TRIANGLE       FilterType = C.TriangleFilter
	FILTER_HERMITE        FilterType = C.HermiteFilter
	FILTER_HANNING        FilterType = C.HanningFilter
	FILTER_HAMMING        FilterType = C.HammingFilter
	FILTER_BLACKMAN       FilterType = C.BlackmanFilter
	FILTER_GAUSSIAN       FilterType = C.GaussianFilter
	FILTER_QUADRATIC      FilterType = C.QuadraticFilter
	FILTER_CUBIC          FilterType = C.CubicFilter
	FILTER_CATROM         FilterType = C.CatromFilter
	FILTER_MITCHELL       FilterType = C.MitchellFilter
	FILTER_JINC           FilterType = C.JincFilter
	FILTER_SINC           FilterType = C.SincFilter
	FILTER_SINC_FAST      FilterType = C.SincFastFilter
	FILTER_KAISER         FilterType = C.KaiserFilter
	FILTER_WELSH          FilterType = C.WelshFilter
	FILTER_PARZEN         FilterType = C.ParzenFilter
	FILTER_BOHMAN         FilterType = C.BohmanFilter
	FILTER_BARTLETT       FilterType = C.BartlettFilter
	FILTER_LAGRANGE       FilterType = C.LagrangeFilter
	FILTER_LANCZOS        FilterType = C.LanczosFilter
	FILTER_LANCZOS_SHARP  FilterType = C.LanczosSharpFilter
	FILTER_LANCZOS2       FilterType = C.Lanczos2Filter
	FILTER_LANCZOS2_SHARP FilterType = C.Lanczos2SharpFilter
	FILTER_ROBIDOUX       FilterType = C.RobidouxFilter
	FILTER_ROBIDOUX_SHARP FilterType = C.RobidouxSharpFilter
	FILTER_COSINE         FilterType = C.CosineFilter
	FILTER_SPLINE         FilterType = C.SplineFilter
	FILTER_SENTINEL       FilterType = C.SentinelFilter
	FILTER_LANCZOS_RADIUS FilterType = C.LanczosRadiusFilter
)
