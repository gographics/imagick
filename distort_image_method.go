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
	UndefinedDistortion             DistortImageMethod = C.UndefinedDistortion
	AffineDistortion                DistortImageMethod = C.AffineDistortion
	AffineProjectionDistortion      DistortImageMethod = C.AffineProjectionDistortion
	ScaleRotateTranslateDistortion  DistortImageMethod = C.ScaleRotateTranslateDistortion
	PerspectiveDistortion           DistortImageMethod = C.PerspectiveDistortion
	PerspectiveProjectionDistortion DistortImageMethod = C.PerspectiveProjectionDistortion
	BilinearForwardDistortion       DistortImageMethod = C.BilinearForwardDistortion
	BilinearDistortion              DistortImageMethod = C.BilinearDistortion
	BilinearReverseDistortion       DistortImageMethod = C.BilinearReverseDistortion
	PolynomialDistortion            DistortImageMethod = C.PolynomialDistortion
	ArcDistortion                   DistortImageMethod = C.ArcDistortion
	PolarDistortion                 DistortImageMethod = C.PolarDistortion
	DePolarDistortion               DistortImageMethod = C.DePolarDistortion
	Cylinder2PlaneDistortion        DistortImageMethod = C.Cylinder2PlaneDistortion
	Plane2CylinderDistortion        DistortImageMethod = C.Plane2CylinderDistortion
	BarrelDistortion                DistortImageMethod = C.BarrelDistortion
	BarrelInverseDistortion         DistortImageMethod = C.BarrelInverseDistortion
	ShepardsDistortion              DistortImageMethod = C.ShepardsDistortion
	ResizeDistortion                DistortImageMethod = C.ResizeDistortion
	SentinelDistortion              DistortImageMethod = C.SentinelDistortion
)

var distortImageMethodStrings = map[DistortImageMethod]string{
	UndefinedDistortion:             "UndefinedDistortion",
	AffineDistortion:                "AffineDistortion",
	AffineProjectionDistortion:      "AffineProjectionDistortion",
	ScaleRotateTranslateDistortion:  "ScaleRotateTranslateDistortion",
	PerspectiveDistortion:           "PerspectiveDistortion",
	PerspectiveProjectionDistortion: "PerspectiveProjectionDistortion",
	BilinearForwardDistortion:       "BilinearForwardDistortion",
	BilinearReverseDistortion:       "BilinearReverseDistortion",
	PolynomialDistortion:            "PolynomialDistortion",
	ArcDistortion:                   "ArcDistortion",
	PolarDistortion:                 "PolarDistortion",
	DePolarDistortion:               "DePolarDistortion",
	Cylinder2PlaneDistortion:        "Cylinder2PlaneDistortion",
	Plane2CylinderDistortion:        "Plane2CylinderDistortion",
	BarrelDistortion:                "BarrelDistortion",
	BarrelInverseDistortion:         "BarrelInverseDistortion",
	ShepardsDistortion:              "ShepardsDistortion",
	ResizeDistortion:                "ResizeDistortion",
	SentinelDistortion:              "SentinelDistortion",
}

func (cst *DistortImageMethod) String() string {
	if v, ok := distortImageMethodStrings[DistortImageMethod(*cst)]; ok {
		return v
	}
	return fmt.Sprintf("UnknownDistortImageMethod[%d]", *cst)
}
