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
	UndefinedType            ImageType = C.UndefinedType
	BilevelType              ImageType = C.BilevelType
	GrayscaleType            ImageType = C.GrayscaleType
	GrayscaleMatteType       ImageType = C.GrayscaleMatteType
	PaletteType              ImageType = C.PaletteType
	PaletteMatteType         ImageType = C.PaletteMatteType
	TrueColorType            ImageType = C.TrueColorType
	TrueColorMatteType       ImageType = C.TrueColorMatteType
	ColorSeparationType      ImageType = C.ColorSeparationType
	ColorSeparationMatteType ImageType = C.ColorSeparationMatteType
	OptimizeType             ImageType = C.OptimizeType
	PaletteBilevelMatteType  ImageType = C.PaletteBilevelMatteType
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
