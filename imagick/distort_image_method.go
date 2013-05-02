package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
)

type DistortImageMethod int

const (
	DISTORTION_UNDEFINED              DistortImageMethod = C.UndefinedDistortion
	DISTORTION_AFFINE                 DistortImageMethod = C.AffineDistortion
	DISTORTION_AFFINE_PROJECTION      DistortImageMethod = C.AffineProjectionDistortion
	DISTORTION_SCALE_ROTATE_TRANSLATE DistortImageMethod = C.ScaleRotateTranslateDistortion
	DISTORTION_PERSPECTIVE            DistortImageMethod = C.PerspectiveDistortion
	DISTORTION_PERSPECTIVE_PROJECTION DistortImageMethod = C.PerspectiveProjectionDistortion
	DISTORTION_BILINEAR_FORWARD       DistortImageMethod = C.BilinearForwardDistortion
	DISTORTION_BILINEAR               DistortImageMethod = C.BilinearDistortion
	DISTORTION_BILINEAR_REVERSE       DistortImageMethod = C.BilinearReverseDistortion
	DISTORTION_POLYNOMIAL             DistortImageMethod = C.PolynomialDistortion
	DISTORTION_ARC                    DistortImageMethod = C.ArcDistortion
	DISTORTION_POLAR                  DistortImageMethod = C.PolarDistortion
	DISTORTION_DE_POLAR               DistortImageMethod = C.DePolarDistortion
	DISTORTION_CYLINDER_2_PLANE       DistortImageMethod = C.Cylinder2PlaneDistortion
	DISTORTION_PLANE_2_CYLINDER       DistortImageMethod = C.Plane2CylinderDistortion
	DISTORTION_BARREL                 DistortImageMethod = C.BarrelDistortion
	DISTORTION_BARREL_INVERSE         DistortImageMethod = C.BarrelInverseDistortion
	DISTORTION_SHEPARDS               DistortImageMethod = C.ShepardsDistortion
	DISTORTION_RESIZE                 DistortImageMethod = C.ResizeDistortion
	DISTORTION_SENTINEL               DistortImageMethod = C.SentinelDistortion
)
