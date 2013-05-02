package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
)

type FillRule int

const (
	FILL_UNDEFINED FillRule = C.UndefinedRule
	FILL_EVEN_ODD  FillRule = C.EvenOddRule
	FILL_NON_ZERO  FillRule = C.NonZeroRule
)
