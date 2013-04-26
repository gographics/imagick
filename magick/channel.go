package magick

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
	0:  "UndefinedColorspace",
	1:  "RGBColorspace",
	2:  "GRAYColorspace",
	3:  "TransparentColorspace",
	4:  "OHTAColorspace",
	5:  "LabColorspace",
	6:  "XYZColorspace",
	7:  "YCbCrColorspace",
	8:  "YCCColorspace",
	9:  "YIQColorspace",
	10: "YPbPrColorspace",
	11: "YUVColorspace",
	12: "CMYKColorspace",
	13: "SRGBColorspace",
	14: "HSBColorspace",
	15: "HSLColorspace",
	16: "HWBColorspace",
	17: "Rec601LumaColorspace",
	18: "Rec601YCbCrColorspace",
	19: "Rec709LumaColorspace",
	20: "Rec709YCbCrColorspace",
	21: "LogColorspace",
	22: "CMYColorspace",
}

func (cst *ColorspaceType) String() string {
	if v, ok := colorspaceTypeStrings[ColorspaceType(*cst)]; ok {
		return v
	}
	return fmt.Sprintf("UnknownColorspace[%d]", *cst)
}
