package magick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

import (
//"fmt"
//"unsafe"
)

type DrawingWand struct {
	draw *C.DrawingWand
}

func NewDrawingWand() *DrawingWand {
	return &DrawingWand{draw: C.NewDrawingWand()}
}
