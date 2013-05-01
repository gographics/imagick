package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
)

type SparseColorMethod int

const (
	UndefinedColorInterpolate   = SparseColorMethod(UndefinedDistortion)
	BarycentricColorInterpolate = SparseColorMethod(AffineDistortion)
	BilinearColorInterpolate    = SparseColorMethod(BilinearReverseDistortion)
	PolynomialColorInterpolate  = SparseColorMethod(PolynomialDistortion)
	ShepardsColorInterpolate    = SparseColorMethod(ShepardsDistortion)
	// Methods unique to SparseColor().
	VoronoiColorInterpolate = SparseColorMethod(SentinelDistortion)
	InverseColorInterpolate = VoronoiColorInterpolate + 1
)

var sparseColorMethodStrings = map[SparseColorMethod]string{
	UndefinedColorInterpolate:   "UndefinedColorInterpolate",
	BarycentricColorInterpolate: "BarycentricColorInterpolate",
	BilinearColorInterpolate:    "BilinearColorInterpolate",
	PolynomialColorInterpolate:  "PolynomialColorInterpolate",
	ShepardsColorInterpolate:    "ShepardsColorInterpolate",
	VoronoiColorInterpolate:     "VoronoiColorInterpolate",
	InverseColorInterpolate:     "InverseColorInterpolate",
}

func (ct *SparseColorMethod) String() string {
	if v, ok := sparseColorMethodStrings[SparseColorMethod(*ct)]; ok {
		return v
	}
	return fmt.Sprintf("UnknownCompression[%d]", *ct)
}
