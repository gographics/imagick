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
	UndefinedColorInterpolate   SparseColorMethod = C.UndefinedColorInterpolate
	BarycentricColorInterpolate SparseColorMethod = C.BarycentricColorInterpolate
	BilinearColorInterpolate    SparseColorMethod = C.BilinearColorInterpolate
	PolynomialColorInterpolate  SparseColorMethod = C.PolynomialColorInterpolate
	ShepardsColorInterpolate    SparseColorMethod = C.ShepardsColorInterpolate
	VoronoiColorInterpolate     SparseColorMethod = C.VoronoiColorInterpolate
	InverseColorInterpolate     SparseColorMethod = C.InverseColorInterpolate
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
