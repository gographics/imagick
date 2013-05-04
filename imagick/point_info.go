package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

type PointInfo struct {
	X float64
	Y float64
}
