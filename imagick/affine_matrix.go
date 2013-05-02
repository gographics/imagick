package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

type AffineMatrix struct {
	am *AffineMatrix
}
