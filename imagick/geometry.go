package imagick

/*
#include <MagickCore/MagickCore.h>
#include <MagickWand/MagickWand.h>
*/
import "C"

func ParseGeometry(geometry string, info *GeometryInfo) uint {
	var gi C.GeometryInfo
	flags := C.ParseGeometry(C.CString(geometry), &gi)
	*info = *newGeometryInfo(&gi)
	return uint(flags)
}

func ParseMetaGeometry(geometry string, x *int, y *int, width *uint, height *uint) uint {
	var info C.RectangleInfo

	flags := C.ParseMetaGeometry(C.CString(geometry), &info.x, &info.y, &info.width, &info.height)

	*x = int(info.x)
	*y = int(info.y)
	*width = uint(info.width)
	*height = uint(info.height)

	return uint(flags)
}

func (mw *MagickWand) ParseGravityGeometry(geometry string, rect *RectangleInfo, exception *ExceptionInfo) uint {
	var info C.RectangleInfo
	var ex C.ExceptionInfo

	image := C.GetImageFromMagickWand(mw.mw)
	flags := C.ParseGravityGeometry(image, C.CString(geometry), &info, &ex)
	*rect = *newRectangleInfo(&info)
	*exception = *newExceptionInfo(&ex)
	return uint(flags)
}
