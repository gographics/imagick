package imagick

/*
#cgo pkg-config: MagickCore
#include <magick/MagickCore.h>
*/
import "C"

type RectangleInfo struct {
	info *C.RectangleInfo
}
