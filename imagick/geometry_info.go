package imagick

/*
#include <magick/MagickCore.h>
#include <magick/geometry.h>
#include <magick/morphology.h>
*/
import "C"

type GeometryInfo struct {
	gi C.GeometryInfo
}
