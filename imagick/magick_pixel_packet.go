package imagick

/*
#cgo pkg-config: MagickCore
#include <magick/MagickCore.h>
*/
import "C"

import (
//"fmt"
//"unsafe"
)

type MagickPixelPacket struct {
	mpp *C.MagickPixelPacket
}

func newMagickPixelPacketFromCAPI(mpp *C.MagickPixelPacket) *MagickPixelPacket {
	return &MagickPixelPacket{mpp}
}
