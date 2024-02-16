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

func SetGeometry(image *Image, info *RectangleInfo) {
	var ri C.RectangleInfo

	C.SetGeometry(image.img, &ri)
	*info = *newRectangleInfo(&ri)
}

func ParseAbsoluteGeometry(geometry string, info *RectangleInfo) uint {
	var ri C.RectangleInfo

	ri.x = C.ssize_t(info.X)
	ri.y = C.ssize_t(info.Y)
	ri.width = C.size_t(info.Width)
	ri.height = C.size_t(info.Height)

	flags := C.ParseAbsoluteGeometry(C.CString(geometry), &ri)
	*info = *newRectangleInfo(&ri)

	return uint(flags)
}

func ParseMetaGeometry(geometry string, x *int, y *int, width *uint, height *uint) uint {
	var ri C.RectangleInfo

	ri.x = C.ssize_t(*x)
	ri.y = C.ssize_t(*y)
	ri.width = C.size_t(*width)
	ri.height = C.size_t(*height)

	flags := C.ParseMetaGeometry(C.CString(geometry), &ri.x, &ri.y, &ri.width, &ri.height)

	*x = int(ri.x)
	*y = int(ri.y)
	*width = uint(ri.width)
	*height = uint(ri.height)

	return uint(flags)
}

func ParseGravityGeometry(image *Image, geometry string, rect *RectangleInfo, exception *ExceptionInfo) uint {
	var ri C.RectangleInfo
	var ex C.ExceptionInfo

	flags := C.ParseGravityGeometry(image.img, C.CString(geometry), &ri, &ex)
	*rect = *newRectangleInfo(&ri)
	*exception = *newExceptionInfo(&ex)

	return uint(flags)
}

func ParsePageGeometry(image *Image, geometry string, rect *RectangleInfo, exception *ExceptionInfo) uint {
	var ri C.RectangleInfo
	var ex C.ExceptionInfo

	flags := C.ParsePageGeometry(image.img, C.CString(geometry), &ri, &ex)
	*rect = *newRectangleInfo(&ri)
	*exception = *newExceptionInfo(&ex)

	return uint(flags)
}

func ParseRegionGeometry(image *Image, geometry string, rect *RectangleInfo, exception *ExceptionInfo) uint {
	var ri C.RectangleInfo
	var ex C.ExceptionInfo

	flags := C.ParseRegionGeometry(image.img, C.CString(geometry), &ri, &ex)
	*rect = *newRectangleInfo(&ri)
	*exception = *newExceptionInfo(&ex)

	return uint(flags)
}
