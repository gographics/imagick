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
	UndefinedDitherMethod DitherMethod = iota
	NoDitherMethod
	RiemersmaDitherMethod
	FloydSteinbergDitherMethod
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
