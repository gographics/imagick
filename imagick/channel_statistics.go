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

type ChannelStatistics struct {
	cs *C.ChannelStatistics
}
