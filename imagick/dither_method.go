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
	DITHER_METHOD_UNDEFINED       DitherMethod = C.UndefinedDitherMethod
	DITHER_METHOD_NO              DitherMethod = C.NoDitherMethod
	DITHER_METHOD_RIEMERSMA       DitherMethod = C.RiemersmaDitherMethod
	DITHER_METHOD_FLOYD_STEINBERG DitherMethod = C.FloydSteinbergDitherMethod
)
