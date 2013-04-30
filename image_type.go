package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
)

type ImageType int

const (
	UndefinedType ImageType = iota
	BilevelType
	GrayscaleType
	GrayscaleMatteType
	PaletteType
	PaletteMatteType
	TrueColorType
	TrueColorMatteType
	ColorSeparationType
	ColorSeparationMatteType
	OptimizeType
	PaletteBilevelMatteType
)

var imageTypeStrings = map[ImageType]string{
	UndefinedType:            "UndefinedType",
	BilevelType:              "BilevelType",
	GrayscaleType:            "GrayscaleType",
	GrayscaleMatteType:       "GrayscaleMatteType",
	PaletteType:              "PaletteType",
	PaletteMatteType:         "PaletteMatteType",
	TrueColorType:            "TrueColorType",
	TrueColorMatteType:       "TrueColorMatteType",
	ColorSeparationType:      "ColorSeparationType",
	ColorSeparationMatteType: "ColorSeparationMatteType",
	OptimizeType:             "OptimizeType",
	PaletteBilevelMatteType:  "PaletteBilevelMatteType",
}

func (ct *ImageType) String() string {
	if v, ok := imageTypeStrings[ImageType(*ct)]; ok {
		return v
	}
	return fmt.Sprintf("UnknownCompression[%d]", *ct)
}
