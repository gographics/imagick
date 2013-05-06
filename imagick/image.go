// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#cgo pkg-config: MagickCore
#include <magick/MagickCore.h>
*/
import "C"

type Image struct {
	img *C.Image
}

func NewMagickImage(info *ImageInfo, width, height uint, background *MagickPixelPacket) *Image {
	return &Image{img: C.NewMagickImage(info.info, C.size_t(width), C.size_t(height), background.mpp)}
}
