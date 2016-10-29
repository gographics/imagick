package imagick

/*
#include <magick/MagickCore.h>
*/
import "C"

type GeometryInfo struct {
	gi C.GeometryInfo
}
