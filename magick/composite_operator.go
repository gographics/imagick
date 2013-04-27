package magick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
)

type CompositeOperator int

const (
	UndefinedCompositeOp CompositeOperator = iota
	NoCompositeOp
	ModulusAddCompositeOp
	AtopCompositeOp
	BlendCompositeOp
	BumpmapCompositeOp
	ChangeMaskCompositeOp
	ClearCompositeOp
	ColorBurnCompositeOp
	ColorDodgeCompositeOp
	ColorizeCompositeOp
	CopyBlackCompositeOp
	CopyBlueCompositeOp
	CopyCompositeOp
	CopyCyanCompositeOp
	CopyGreenCompositeOp
	CopyMagentaCompositeOp
	CopyOpacityCompositeOp
	CopyRedCompositeOp
	CopyYellowCompositeOp
	DarkenCompositeOp
	DstAtopCompositeOp
	DstCompositeOp
	DstInCompositeOp
	DstOutCompositeOp
	DstOverCompositeOp
	DifferenceCompositeOp
	DisplaceCompositeOp
	DissolveCompositeOp
	ExclusionCompositeOp
	HardLightCompositeOp
	HueCompositeOp
	InCompositeOp
	LightenCompositeOp
	LinearLightCompositeOp
	LuminizeCompositeOp
	MinusDstCompositeOp
	ModulateCompositeOp
	MultiplyCompositeOp
	OutCompositeOp
	OverCompositeOp
	OverlayCompositeOp
	PlusCompositeOp
	ReplaceCompositeOp
	SaturateCompositeOp
	ScreenCompositeOp
	SoftLightCompositeOp
	SrcAtopCompositeOp
	SrcCompositeOp
	SrcInCompositeOp
	SrcOutCompositeOp
	SrcOverCompositeOp
	ModulusSubtractCompositeOp
	ThresholdCompositeOp
	XorCompositeOp
	DivideDstCompositeOp
	DistortCompositeOp
	BlurCompositeOp
	PegtopLightCompositeOp
	VividLightCompositeOp
	PinLightCompositeOp
	LinearDodgeCompositeOp
	LinearBurnCompositeOp
	MathematicsCompositeOp
	DivideSrcCompositeOp
	MinusSrcCompositeOp
	DarkenIntensityCompositeOp
	LightenIntensityCompositeOp
)

var compositeOperatorStrings = map[CompositeOperator]string{
	UndefinedCompositeOp:        "UndefinedCompositeOp",
	NoCompositeOp:               "NoCompositeOp",
	ModulusAddCompositeOp:       "ModulusAddCompositeOp",
	AtopCompositeOp:             "AtopCompositeOp",
	BlendCompositeOp:            "BlendCompositeOp",
	BumpmapCompositeOp:          "BumpmapCompositeOp",
	ChangeMaskCompositeOp:       "ChangeMaskCompositeOp",
	ClearCompositeOp:            "ClearCompositeOp",
	ColorBurnCompositeOp:        "ColorBurnCompositeOp",
	ColorDodgeCompositeOp:       "ColorDodgeCompositeOp",
	ColorizeCompositeOp:         "ColorizeCompositeOp",
	CopyBlackCompositeOp:        "CopyBlackCompositeOp",
	CopyBlueCompositeOp:         "CopyBlueCompositeOp",
	CopyCompositeOp:             "CopyCompositeOp",
	CopyCyanCompositeOp:         "CopyCyanCompositeOp",
	CopyGreenCompositeOp:        "CopyGreenCompositeOp",
	CopyMagentaCompositeOp:      "CopyMagentaCompositeOp",
	CopyOpacityCompositeOp:      "CopyOpacityCompositeOp",
	CopyRedCompositeOp:          "CopyRedCompositeOp",
	CopyYellowCompositeOp:       "CopyYellowCompositeOp",
	DarkenCompositeOp:           "DarkenCompositeOp",
	DstAtopCompositeOp:          "DstAtopCompositeOp",
	DstCompositeOp:              "DstCompositeOp",
	DstInCompositeOp:            "DstInCompositeOp",
	DstOutCompositeOp:           "DstOutCompositeOp",
	DstOverCompositeOp:          "DstOverCompositeOp",
	DifferenceCompositeOp:       "DifferenceCompositeOp",
	DisplaceCompositeOp:         "DisplaceCompositeOp",
	DissolveCompositeOp:         "DissolveCompositeOp",
	ExclusionCompositeOp:        "ExclusionCompositeOp",
	HardLightCompositeOp:        "HardLightCompositeOp",
	HueCompositeOp:              "HueCompositeOp",
	InCompositeOp:               "InCompositeOp",
	LightenCompositeOp:          "LightenCompositeOp",
	LinearLightCompositeOp:      "LinearLightCompositeOp",
	LuminizeCompositeOp:         "LuminizeCompositeOp",
	MinusDstCompositeOp:         "MinusDstCompositeOp",
	ModulateCompositeOp:         "ModulateCompositeOp",
	MultiplyCompositeOp:         "MultiplyCompositeOp",
	OutCompositeOp:              "OutCompositeOp",
	OverCompositeOp:             "OverCompositeOp",
	OverlayCompositeOp:          "OverlayCompositeOp",
	PlusCompositeOp:             "PlusCompositeOp",
	ReplaceCompositeOp:          "ReplaceCompositeOp",
	SaturateCompositeOp:         "SaturateCompositeOp",
	ScreenCompositeOp:           "ScreenCompositeOp",
	SoftLightCompositeOp:        "SoftLightCompositeOp",
	SrcAtopCompositeOp:          "SrcAtopCompositeOp",
	SrcCompositeOp:              "SrcCompositeOp",
	SrcInCompositeOp:            "SrcInCompositeOp",
	SrcOutCompositeOp:           "SrcOutCompositeOp",
	SrcOverCompositeOp:          "SrcOverCompositeOp",
	ModulusSubtractCompositeOp:  "ModulusSubtractCompositeOp",
	ThresholdCompositeOp:        "ThresholdCompositeOp",
	XorCompositeOp:              "XorCompositeOp",
	DivideDstCompositeOp:        "DivideDstCompositeOp",
	DistortCompositeOp:          "DistortCompositeOp",
	BlurCompositeOp:             "BlurCompositeOp",
	PegtopLightCompositeOp:      "PegtopLightCompositeOp",
	VividLightCompositeOp:       "VividLightCompositeOp",
	PinLightCompositeOp:         "PinLightCompositeOp",
	LinearDodgeCompositeOp:      "LinearDodgeCompositeOp",
	LinearBurnCompositeOp:       "LinearBurnCompositeOp",
	MathematicsCompositeOp:      "MathematicsCompositeOp",
	DivideSrcCompositeOp:        "DivideSrcCompositeOp",
	MinusSrcCompositeOp:         "MinusSrcCompositeOp",
	DarkenIntensityCompositeOp:  "DarkenIntensityCompositeOp",
	LightenIntensityCompositeOp: "LightenIntensityCompositeOp",
}

func (cst *CompositeOperator) String() string {
	if v, ok := compositeOperatorStrings[CompositeOperator(*cst)]; ok {
		return v
	}
	return fmt.Sprintf("UnknownCompositeOp[%d]", *cst)
}
