package imagick

/*
#cgo pkg-config: MagickCore
#include <magick/MagickCore.h>
*/
import "C"

import (
//"fmt"
//"unsafe"
)

// TODO how to map the different types of Quantum?
type Quantum uint16
