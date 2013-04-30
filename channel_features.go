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

type ChannelFeatures struct {
	cf *C.ChannelFeatures
}
