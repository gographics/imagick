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
