package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

type DrawInfo struct {
	di *C.DrawInfo
}

func NewDrawInfo() *DrawInfo {
	return &DrawInfo{C.AcquireDrawInfo()}
}
