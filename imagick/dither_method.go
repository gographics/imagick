package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
)

type DitherMethod int

const (
	UndefinedDitherMethod      DitherMethod = C.UndefinedDitherMethod
	NoDitherMethod             DitherMethod = C.NoDitherMethod
	RiemersmaDitherMethod      DitherMethod = C.RiemersmaDitherMethod
	FloydSteinbergDitherMethod DitherMethod = C.FloydSteinbergDitherMethod
)

var ditherMethodStrings = map[DitherMethod]string{
	UndefinedDitherMethod:      "UndefinedDitherMethod",
	NoDitherMethod:             "NoDitherMethod",
	RiemersmaDitherMethod:      "RiemersmaDitherMethod",
	FloydSteinbergDitherMethod: "FloydSteinbergDitherMethod",
}

func (ct *DitherMethod) String() string {
	if v, ok := ditherMethodStrings[DitherMethod(*ct)]; ok {
		return v
	}
	return fmt.Sprintf("UnknownCompression[%d]", *ct)
}
