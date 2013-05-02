package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

type OrientationType int

const (
	ORIENTATION_UNDEFINED    OrientationType = C.UndefinedOrientation
	ORIENTATION_TOP_LEFT     OrientationType = C.TopLeftOrientation
	ORIENTATION_TOP_RIGHT    OrientationType = C.TopRightOrientation
	ORIENTATION_BOTTOM_RIGHT OrientationType = C.BottomRightOrientation
	ORIENTATION_BOTTOM_LEFT  OrientationType = C.BottomLeftOrientation
	ORIENTATION_LEFT_TOP     OrientationType = C.LeftTopOrientation
	ORIENTATION_RIGHT_TOP    OrientationType = C.RightTopOrientation
	ORIENTATION_RIGHT_BOTTOM OrientationType = C.RightBottomOrientation
	ORIENTATION_LEFT_BOTTOM  OrientationType = C.LeftBottomOrientation
)
