package magick

/*
#cgo pkg-config: MagickCore
#include <magick/MagickCore.h>
*/
import "C"

type KernelInfo struct {
	info *C.KernelInfo
}
