package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
)

type StorageType int

const (
	PIXEL_UNDEFINED StorageType = C.UndefinedPixel
	PIXEL_CHAR      StorageType = C.CharPixel
	PIXEL_DOUBLE    StorageType = C.DoublePixel
	PIXEL_FLOAT     StorageType = C.FloatPixel
	PIXEL_INTEGER   StorageType = C.IntegerPixel
	PIXEL_LONG      StorageType = C.LongPixel
	PIXEL_QUANTUM   StorageType = C.QuantumPixel
	PIXEL_SHORT     StorageType = C.ShortPixel
)
