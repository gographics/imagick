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
	UndefinedColorInterpolate   = UndefinedDistortion
	BarycentricColorInterpolate = AffineDistortion
	BilinearColorInterpolate    = BilinearReverseDistortion
	PolynomialColorInterpolate  = PolynomialDistortion
	ShepardsColorInterpolate    = ShepardsDistortion
	// Methods unique to SparseColor().
	VoronoiColorInterpolate = SentinelDistortion
	InverseColorInterpolate
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
