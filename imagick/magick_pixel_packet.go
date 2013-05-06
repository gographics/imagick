// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#cgo pkg-config: MagickCore
#include <magick/MagickCore.h>
*/
import "C"

type MagickPixelPacket struct {
	mpp *C.MagickPixelPacket
}

func newMagickPixelPacketFromCAPI(mpp *C.MagickPixelPacket) *MagickPixelPacket {
	return &MagickPixelPacket{mpp}
}
