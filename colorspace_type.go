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
	UndefinedColorspace   ColorspaceType = C.UndefinedColorspace
	RGBColorspace         ColorspaceType = C.RGBColorspace
	GRAYColorspace        ColorspaceType = C.GRAYColorspace
	TransparentColorspace ColorspaceType = C.TransparentColorspace
	OHTAColorspace        ColorspaceType = C.OHTAColorspace
	LabColorspace         ColorspaceType = C.LabColorspace
	XYZColorspace         ColorspaceType = C.XYZColorspace
	YCbCrColorspace       ColorspaceType = C.YCbCrColorspace
	YCCColorspace         ColorspaceType = C.YCCColorspace
	YIQColorspace         ColorspaceType = C.YIQColorspace
	YPbPrColorspace       ColorspaceType = C.YPbPrColorspace
	YUVColorspace         ColorspaceType = C.YUVColorspace
	CMYKColorspace        ColorspaceType = C.CMYKColorspace
	SRGBColorspace        ColorspaceType = C.sRGBColorspace
	HSBColorspace         ColorspaceType = C.HSBColorspace
	HSLColorspace         ColorspaceType = C.HSLColorspace
	HWBColorspace         ColorspaceType = C.HWBColorspace
	Rec601LumaColorspace  ColorspaceType = C.Rec601LumaColorspace
	Rec601YCbCrColorspace ColorspaceType = C.Rec601YCbCrColorspace
	Rec709LumaColorspace  ColorspaceType = C.Rec709LumaColorspace
	Rec709YCbCrColorspace ColorspaceType = C.Rec709YCbCrColorspace
	LogColorspace         ColorspaceType = C.LogColorspace
	CMYColorspace         ColorspaceType = C.CMYColorspace
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
