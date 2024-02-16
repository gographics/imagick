package imagick

/*
#include <MagickCore/MagickCore.h>
#include <MagickWand/MagickWand.h>
*/
import "C"

// ParseGeometry parses a geometry specification and returns the sigma,
// rho, xi, and psi values.  It also returns flags that indicates which
// of the four values (sigma, rho, xi, psi) were located in the string, and
// whether the xi or pi values are negative.
//
// In addition, it reports if there are any of meta characters (%, !, <, >, @,
// and ^) flags present. It does not report the location of the percentage
// relative to the values.
//
// Values may also be separated by commas, colons, or slashes, and offsets.
// Chroma subsampling definitions have to be in the form of a:b:c.  Offsets may
// be prefixed by multiple signs to make offset string substitutions easier to
// handle from shell scripts.  For example: "-10-10", "-+10-+10", or "+-10+-10"
// will generate negative offsets, while "+10+10", "++10++10", or "--10--10"
// will generate positive offsets.
//
// A description of each parameter follows:
//
//	o geometry:  The geometry string (e.g. "100x100+10+10").
//	o geometry_info:  returns the parsed width/height/x/y in this structure.
func ParseGeometry(geometry string, info *GeometryInfo) uint {
	var gi C.GeometryInfo
	flags := C.ParseGeometry(C.CString(geometry), &gi)
	*info = *newGeometryInfo(&gi)
	return uint(flags)
}

// SetGeometry sets the geometry to its default values.
//
// A description of each parameter follows:
//
//	o image: the image.
//	o geometry: the geometry.
func SetGeometry(image *Image, info *RectangleInfo) {
	var ri C.RectangleInfo

	C.SetGeometry(image.img, &ri)
	*info = *newRectangleInfo(&ri)
}

// ParseAbsoluteGeometry returns a region as defined by the geometry string,
// without any modification by percentages or gravity.
//
// It currently just a wrapper around GetGeometry(), but may be expanded in
// the future to handle other positioning information.
//
// A description of each parameter follows:
//
//	o geometry:  The geometry string (e.g. "100x100+10+10").
//	o region_info: the region as defined by the geometry string.
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

// ParseMetaGeometry is similar to GetGeometry() except the returned
// geometry is modified as determined by the meta characters:  %, !, <, >, @,
// :, and ^ in relation to image resizing.
//
// Final image dimensions are adjusted so as to preserve the aspect ratio as
// much as possible, while generating a integer (pixel) size, and fitting the
// image within the specified geometry width and height.
//
// Flags are interpreted...
//
//	%   geometry size is given percentage of original width and height given
//	!   do not try to preserve aspect ratio
//	<   only enlarge images smaller that geometry
//	>   only shrink images larger than geometry
//	@   fit image to contain at most this many pixels
//	:   width and height denotes an aspect ratio
//	^   contain the given geometry given, (minimal dimensions given)
//
// A description of each parameter follows:
//
//	o geometry:  The geometry string (e.g. "100x100+10+10").
//	o x,y:  The x and y offset, set according to the geometry specification.
//	o width,height:  The width and height of original image, modified by
//	  the given geometry specification.
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

// ParseGravityGeometry returns a region as defined by the geometry string
// with respect to the given image page (canvas) dimensions and the images
// gravity setting.
//
// This is typically used for specifying a area within a given image for
// cropping images to a smaller size, chopping out rows and or columns, or
// resizing and positioning overlay images.
//
// Percentages are relative to image size and not page size, and are set to
// nearest integer (pixel) size.
//
// The format of the ParseGravityGeometry method is:
//
//	MagickStatusType ParseGravityGeometry(Image *image,const char *geometry,
//	  RectangleInfo *region_info,ExceptionInfo *exception)
//
// A description of each parameter follows:
//
//	o geometry:  The geometry string (e.g. "100x100+10+10").
//	o region_info: the region as defined by the geometry string with respect
//	  to the image dimensions and its gravity.
//	o exception: return any errors or warnings in this structure.
func ParseGravityGeometry(image *Image, geometry string, rect *RectangleInfo, exception *ExceptionInfo) uint {
	var ri C.RectangleInfo
	var ex C.ExceptionInfo

	flags := C.ParseGravityGeometry(image.img, C.CString(geometry), &ri, &ex)
	*rect = *newRectangleInfo(&ri)
	*exception = *newExceptionInfo(&ex)

	return uint(flags)
}

// ParsePageGeometry returns a region as defined by the geometry string with
// respect to the image page (canvas) dimensions.
//
// WARNING: Percentage dimensions remain relative to the actual image
// dimensions, and not canvas dimensions.
//
// A description of each parameter follows:
//
//	o geometry:  The geometry string (e.g. "100x100+10+10").
//	o region_info: the region as defined by the geometry string with
//	  respect to the image and its gravity.
//	o exception: return any errors or warnings in this structure.
func ParsePageGeometry(image *Image, geometry string, rect *RectangleInfo, exception *ExceptionInfo) uint {
	var ri C.RectangleInfo
	var ex C.ExceptionInfo

	flags := C.ParsePageGeometry(image.img, C.CString(geometry), &ri, &ex)
	*rect = *newRectangleInfo(&ri)
	*exception = *newExceptionInfo(&ex)

	return uint(flags)
}

// ParseRegionGeometry returns a region as defined by the geometry string
// with respect to the image dimensions and aspect ratio.
//
// This is basically a wrapper around ParseMetaGeometry.  This is typically
// used to parse a geometry string to work out the final integer dimensions
// for image resizing.
//
// A description of each parameter follows:
//
//	o geometry:  The geometry string (e.g. "100x100+10+10").
//	o region_info: the region as defined by the geometry string.
//	o exception: return any errors or warnings in this structure.
func ParseRegionGeometry(image *Image, geometry string, rect *RectangleInfo, exception *ExceptionInfo) uint {
	var ri C.RectangleInfo
	var ex C.ExceptionInfo

	flags := C.ParseRegionGeometry(image.img, C.CString(geometry), &ri, &ex)
	*rect = *newRectangleInfo(&ri)
	*exception = *newExceptionInfo(&ex)

	return uint(flags)
}
