package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
)

type EndianType int

const (
	ENDIAN_UNDEFINED EndianType = C.UndefinedEndian
	ENDIAN_LSB       EndianType = C.LSBEndian
	ENDIAN_MSB       EndianType = C.MSBEndian
)
