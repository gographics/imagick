package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

var (
	initialized bool
)

// Inicializes the MagickWand environment
func Initialize() {
	if initialized {
		return
	}
	C.MagickWandGenesis()
	initialized = true
}

// Terminates the MagickWand environment
func Terminate() {
	if initialized {
		C.MagickWandTerminus()
		initialized = false
	}
}
