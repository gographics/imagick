package imagick

/*
#cgo pkg-config: MagickCore
#include <magick/MagickCore.h>
*/
import "C"

type ImageInfo struct {
	info *C.ImageInfo
}
