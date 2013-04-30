package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
)

type InterpolatePixelMethod int

const (
	UndefinedInterpolatePixel       InterpolatePixelMethod = iota
	AverageInterpolatePixel                                // Average 4 nearest neighbours
	BicubicInterpolatePixel                                // Catmull-Rom interpolation
	BilinearInterpolatePixel                               // Triangular filter interpolation
	FilterInterpolatePixel                                 // Use resize filter - (very slow)
	IntegerInterpolatePixel                                // Integer (floor) interpolation
	MeshInterpolatePixel                                   // Triangular mesh interpolation
	NearestNeighborInterpolatePixel                        // Nearest neighbour only
	SplineInterpolatePixel                                 // Cubic Spline (blurred) interpolation
	Average9InterpolatePixel                               // Average 9 nearest neighbours
	Average16InterpolatePixel                              // Average 16 nearest neighbours
	BlendInterpolatePixel                                  // blend of nearest 1, 2 or 4 pixels
	BackgroundInterpolatePixel                             // just return background color
	CatromInterpolatePixel                                 // Catmull-Rom interpolation
)

var interpolatePixelMethodStrings = map[InterpolatePixelMethod]string{
	UndefinedInterpolatePixel:       "UndefinedInterpolatePixel",
	AverageInterpolatePixel:         "AverageInterpolatePixel",
	BicubicInterpolatePixel:         "BicubicInterpolatePixel",
	BilinearInterpolatePixel:        "BilinearInterpolatePixel",
	FilterInterpolatePixel:          "FilterInterpolatePixel",
	IntegerInterpolatePixel:         "IntegerInterpolatePixel",
	MeshInterpolatePixel:            "MeshInterpolatePixel",
	NearestNeighborInterpolatePixel: "NearestNeighborInterpolatePixel",
	SplineInterpolatePixel:          "SplineInterpolatePixel",
	Average9InterpolatePixel:        "Average9InterpolatePixel",
	Average16InterpolatePixel:       "Average16InterpolatePixel",
	BlendInterpolatePixel:           "BlendInterpolatePixel",
	BackgroundInterpolatePixel:      "BackgroundInterpolatePixel",
	CatromInterpolatePixel:          "CatromInterpolatePixel",
}

func (cst *InterpolatePixelMethod) String() string {
	if v, ok := interpolatePixelMethodStrings[InterpolatePixelMethod(*cst)]; ok {
		return v
	}
	return fmt.Sprintf("UnknownInterpolatePixelMethod[%d]", *cst)
}
