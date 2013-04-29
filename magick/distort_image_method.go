package magick

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
	UndefinedDistortion DistortImageMethod = iota
	AffineDistortion
	AffineProjectionDistortion
	ScaleRotateTranslateDistortion
	PerspectiveDistortion
	PerspectiveProjectionDistortion
	BilinearForwardDistortion
	BilinearDistortion DistortImageMethod = iota - 1
	BilinearReverseDistortion
	PolynomialDistortion
	ArcDistortion
	PolarDistortion
	DePolarDistortion
	Cylinder2PlaneDistortion
	Plane2CylinderDistortion
	BarrelDistortion
	BarrelInverseDistortion
	ShepardsDistortion
	ResizeDistortion
	SentinelDistortion
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
