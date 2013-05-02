package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"
import (
	"unsafe"
)

type AffineMatrix struct {
	am *AffineMatrix
}
