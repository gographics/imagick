package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"
import (
	"unsafe"
)

type PointInfo struct {
	pi *PointInfo
}
