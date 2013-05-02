package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
)

type PaintMethod int

const (
	PAINT_METHOD_UNDEFINED    PaintMethod = C.UndefinedMethod
	PAINT_METHOD_POINT        PaintMethod = C.PointMethod
	PAINT_METHOD_REPLACE      PaintMethod = C.ReplaceMethod
	PAINT_METHOD_FLOODFILL    PaintMethod = C.FloodfillMethod
	PAINT_METHOD_FILLTOBORDER PaintMethod = C.FillToBorderMethod
	PAINT_METHOD_RESET        PaintMethod = C.ResetMethod
)
