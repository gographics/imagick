package imagick

/*
#include <MagickCore/MagickCore.h>
*/
import "C"

type GeometryInfo struct {
	gi C.GeometryInfo
}
