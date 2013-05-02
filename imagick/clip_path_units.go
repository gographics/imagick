package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
)

type ClipPathUnits int

const (
	CLIP_UNDEFINED_PATH      ClipPathUnits = C.UndefinedPathUnits
	CLIP_USER_SPACE          ClipPathUnits = C.UserSpace
	CLIP_USER_SPACE_ON_USE   ClipPathUnits = C.UserSpaceOnUse
	CLIP_OBJECT_BOUNDING_BOX ClipPathUnits = C.ObjectBoundingBox
)
