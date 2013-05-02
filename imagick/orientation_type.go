package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
)

type OrientationType int

const (
	UndefinedOrientation   OrientationType = C.UndefinedOrientation
	TopLeftOrientation     OrientationType = C.TopLeftOrientation
	TopRightOrientation    OrientationType = C.TopRightOrientation
	BottomRightOrientation OrientationType = C.BottomRightOrientation
	BottomLeftOrientation  OrientationType = C.BottomLeftOrientation
	LeftTopOrientation     OrientationType = C.LeftTopOrientation
	RightTopOrientation    OrientationType = C.RightTopOrientation
	RightBottomOrientation OrientationType = C.RightBottomOrientation
	LeftBottomOrientation  OrientationType = C.LeftBottomOrientation
)

var orientationTypeStrings = map[OrientationType]string{
	UndefinedOrientation:   "UndefinedOrientation",
	TopLeftOrientation:     "TopLeftOrientation",
	TopRightOrientation:    "TopRightOrientation",
	BottomRightOrientation: "BottomRightOrientation",
	BottomLeftOrientation:  "BottomLeftOrientation",
	LeftTopOrientation:     "LeftTopOrientation",
	RightTopOrientation:    "RightTopOrientation",
	RightBottomOrientation: "RightBottomOrientation",
	LeftBottomOrientation:  "LeftBottomOrientation",
}

func (ct *OrientationType) String() string {
	if v, ok := orientationTypeStrings[OrientationType(*ct)]; ok {
		return v
	}
	return fmt.Sprintf("UnknownCompression[%d]", *ct)
}
