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
	UndefinedOrientation OrientationType = iota
	TopLeftOrientation
	TopRightOrientation
	BottomRightOrientation
	BottomLeftOrientation
	LeftTopOrientation
	RightTopOrientation
	RightBottomOrientation
	LeftBottomOrientation
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
