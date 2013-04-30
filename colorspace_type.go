package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
)

type ColorspaceType int

const (
	UndefinedColorspace ColorspaceType = iota
	RGBColorspace
	GRAYColorspace
	TransparentColorspace
	OHTAColorspace
	LabColorspace
	XYZColorspace
	YCbCrColorspace
	YCCColorspace
	YIQColorspace
	YPbPrColorspace
	YUVColorspace
	CMYKColorspace
	SRGBColorspace
	HSBColorspace
	HSLColorspace
	HWBColorspace
	Rec601LumaColorspace
	Rec601YCbCrColorspace
	Rec709LumaColorspace
	Rec709YCbCrColorspace
	LogColorspace
	CMYColorspace
)

var colorspaceTypeStrings = map[ColorspaceType]string{
	UndefinedColorspace:   "UndefinedColorspace",
	RGBColorspace:         "RGBColorspace",
	GRAYColorspace:        "GRAYColorspace",
	TransparentColorspace: "TransparentColorspace",
	OHTAColorspace:        "OHTAColorspace",
	LabColorspace:         "LabColorspace",
	XYZColorspace:         "XYZColorspace",
	YCbCrColorspace:       "YCbCrColorspace",
	YCCColorspace:         "YCCColorspace",
	YIQColorspace:         "YIQColorspace",
	YPbPrColorspace:       "YPbPrColorspace",
	YUVColorspace:         "YUVColorspace",
	CMYKColorspace:        "CMYKColorspace",
	SRGBColorspace:        "SRGBColorspace",
	HSBColorspace:         "HSBColorspace",
	HSLColorspace:         "HSLColorspace",
	HWBColorspace:         "HWBColorspace",
	Rec601LumaColorspace:  "Rec601LumaColorspace",
	Rec601YCbCrColorspace: "Rec601YCbCrColorspace",
	Rec709LumaColorspace:  "Rec709LumaColorspace",
	Rec709YCbCrColorspace: "Rec709YCbCrColorspace",
	LogColorspace:         "LogColorspace",
	CMYColorspace:         "CMYColorspace",
}

func (cst *ColorspaceType) String() string {
	if v, ok := colorspaceTypeStrings[ColorspaceType(*cst)]; ok {
		return v
	}
	return fmt.Sprintf("UnknownColorspace[%d]", *cst)
}
