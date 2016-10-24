// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type DistortImageMethod int

const (
	DISTORTION_UNDEFINED              DistortImageMethod = C.UndefinedDistortion
	DISTORTION_AFFINE                 DistortImageMethod = C.AffineDistortion
	DISTORTION_AFFINE_PROJECTION      DistortImageMethod = C.AffineProjectionDistortion
	DISTORTION_ARC                    DistortImageMethod = C.ArcDistortion
	DISTORTION_BARREL                 DistortImageMethod = C.BarrelDistortion
	DISTORTION_BARREL_INVERSE         DistortImageMethod = C.BarrelInverseDistortion
	DISTORTION_BILINEAR               DistortImageMethod = C.BilinearDistortion
	DISTORTION_BILINEAR_FORWARD       DistortImageMethod = C.BilinearForwardDistortion
	DISTORTION_BILINEAR_REVERSE       DistortImageMethod = C.BilinearReverseDistortion
	DISTORTION_CYLINDER_2_PLANE       DistortImageMethod = C.Cylinder2PlaneDistortion
	DISTORTION_DE_POLAR               DistortImageMethod = C.DePolarDistortion
	DISTORTION_PERSPECTIVE            DistortImageMethod = C.PerspectiveDistortion
	DISTORTION_PERSPECTIVE_PROJECTION DistortImageMethod = C.PerspectiveProjectionDistortion
	DISTORTION_PLANE_2_CYLINDER       DistortImageMethod = C.Plane2CylinderDistortion
	DISTORTION_POLAR                  DistortImageMethod = C.PolarDistortion
	DISTORTION_POLYNOMIAL             DistortImageMethod = C.PolynomialDistortion
	DISTORTION_RESIZE                 DistortImageMethod = C.ResizeDistortion
	DISTORTION_SCALE_ROTATE_TRANSLATE DistortImageMethod = C.ScaleRotateTranslateDistortion
	DISTORTION_SENTINEL               DistortImageMethod = C.SentinelDistortion
	DISTORTION_SHEPARDS               DistortImageMethod = C.ShepardsDistortion
)
