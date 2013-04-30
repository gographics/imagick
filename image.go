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

func newImageFromCAPI(img *C.Image) *Image {
	return &Image{img}
}

//MagickDestroyImage() dereferences an image, deallocating memory associated with the image if the reference count becomes zero.
//The format of the MagickDestroyImage method is:
//Image *MagickDestroyImage(Image *image)
//A description of each parameter follows:
//image
//the image.
