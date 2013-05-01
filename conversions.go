package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

// Convert a boolean to a integer
func b2i(boolean bool) C.MagickBooleanType {
	if boolean {
		return C.MagickBooleanType(1)
	}
	return C.MagickBooleanType(0)
}
