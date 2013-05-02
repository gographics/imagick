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
	UndefinedInterpolatePixel       InterpolatePixelMethod = C.UndefinedInterpolatePixel
	AverageInterpolatePixel         InterpolatePixelMethod = C.AverageInterpolatePixel         // Average 4 nearest neighbours
	BicubicInterpolatePixel         InterpolatePixelMethod = C.BicubicInterpolatePixel         // Catmull-Rom interpolation
	BilinearInterpolatePixel        InterpolatePixelMethod = C.BilinearInterpolatePixel        // Triangular filter interpolation
	FilterInterpolatePixel          InterpolatePixelMethod = C.FilterInterpolatePixel          // Use resize filter - (very slow)
	IntegerInterpolatePixel         InterpolatePixelMethod = C.IntegerInterpolatePixel         // Integer (floor) interpolation
	MeshInterpolatePixel            InterpolatePixelMethod = C.MeshInterpolatePixel            // Triangular mesh interpolation
	NearestNeighborInterpolatePixel InterpolatePixelMethod = C.NearestNeighborInterpolatePixel // Nearest neighbour only
	SplineInterpolatePixel          InterpolatePixelMethod = C.SplineInterpolatePixel          // Cubic Spline (blurred) interpolation
	Average9InterpolatePixel        InterpolatePixelMethod = C.Average9InterpolatePixel        // Average 9 nearest neighbours
	Average16InterpolatePixel       InterpolatePixelMethod = C.Average16InterpolatePixel       // Average 16 nearest neighbours
	BlendInterpolatePixel           InterpolatePixelMethod = C.BlendInterpolatePixel           // blend of nearest 1, 2 or 4 pixels
	BackgroundInterpolatePixel      InterpolatePixelMethod = C.BackgroundInterpolatePixel      // just return background color
	CatromInterpolatePixel          InterpolatePixelMethod = C.CatromInterpolatePixel          // Catmull-Rom interpolation
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
