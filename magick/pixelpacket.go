package magick

/*
#cgo pkg-config: MagickCore
#include <magick/MagickCore.h>
*/
import "C"

import (
//"fmt"
//"unsafe"
)

type PixelPacket struct {
	pp *C.PixelPacket
}

func newPixelPacketFromCAPI(mpp *C.PixelPacket) *PixelPacket {
	return &PixelPacket{mpp}
}
