package imagick

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
	UndefinedCompositeOp        CompositeOperator = C.UndefinedCompositeOp
	NoCompositeOp               CompositeOperator = C.NoCompositeOp
	ModulusAddCompositeOp       CompositeOperator = C.ModulusAddCompositeOp
	AtopCompositeOp             CompositeOperator = C.AtopCompositeOp
	BlendCompositeOp            CompositeOperator = C.BlendCompositeOp
	BumpmapCompositeOp          CompositeOperator = C.BumpmapCompositeOp
	ChangeMaskCompositeOp       CompositeOperator = C.ChangeMaskCompositeOp
	ClearCompositeOp            CompositeOperator = C.ClearCompositeOp
	ColorBurnCompositeOp        CompositeOperator = C.ColorBurnCompositeOp
	ColorDodgeCompositeOp       CompositeOperator = C.ColorDodgeCompositeOp
	ColorizeCompositeOp         CompositeOperator = C.ColorizeCompositeOp
	CopyBlackCompositeOp        CompositeOperator = C.CopyBlackCompositeOp
	CopyBlueCompositeOp         CompositeOperator = C.CopyBlueCompositeOp
	CopyCompositeOp             CompositeOperator = C.CopyCompositeOp
	CopyCyanCompositeOp         CompositeOperator = C.CopyCyanCompositeOp
	CopyGreenCompositeOp        CompositeOperator = C.CopyGreenCompositeOp
	CopyMagentaCompositeOp      CompositeOperator = C.CopyMagentaCompositeOp
	CopyOpacityCompositeOp      CompositeOperator = C.CopyOpacityCompositeOp
	CopyRedCompositeOp          CompositeOperator = C.CopyRedCompositeOp
	CopyYellowCompositeOp       CompositeOperator = C.CopyYellowCompositeOp
	DarkenCompositeOp           CompositeOperator = C.DarkenCompositeOp
	DstAtopCompositeOp          CompositeOperator = C.DstAtopCompositeOp
	DstCompositeOp              CompositeOperator = C.DstCompositeOp
	DstInCompositeOp            CompositeOperator = C.DstInCompositeOp
	DstOutCompositeOp           CompositeOperator = C.DstOutCompositeOp
	DstOverCompositeOp          CompositeOperator = C.DstOverCompositeOp
	DifferenceCompositeOp       CompositeOperator = C.DifferenceCompositeOp
	DisplaceCompositeOp         CompositeOperator = C.DisplaceCompositeOp
	DissolveCompositeOp         CompositeOperator = C.DissolveCompositeOp
	ExclusionCompositeOp        CompositeOperator = C.ExclusionCompositeOp
	HardLightCompositeOp        CompositeOperator = C.HardLightCompositeOp
	HueCompositeOp              CompositeOperator = C.HueCompositeOp
	InCompositeOp               CompositeOperator = C.InCompositeOp
	LightenCompositeOp          CompositeOperator = C.LightenCompositeOp
	LinearLightCompositeOp      CompositeOperator = C.LinearLightCompositeOp
	LuminizeCompositeOp         CompositeOperator = C.LuminizeCompositeOp
	MinusDstCompositeOp         CompositeOperator = C.MinusDstCompositeOp
	ModulateCompositeOp         CompositeOperator = C.ModulateCompositeOp
	MultiplyCompositeOp         CompositeOperator = C.MultiplyCompositeOp
	OutCompositeOp              CompositeOperator = C.OutCompositeOp
	OverCompositeOp             CompositeOperator = C.OverCompositeOp
	OverlayCompositeOp          CompositeOperator = C.OverlayCompositeOp
	PlusCompositeOp             CompositeOperator = C.PlusCompositeOp
	ReplaceCompositeOp          CompositeOperator = C.ReplaceCompositeOp
	SaturateCompositeOp         CompositeOperator = C.SaturateCompositeOp
	ScreenCompositeOp           CompositeOperator = C.ScreenCompositeOp
	SoftLightCompositeOp        CompositeOperator = C.SoftLightCompositeOp
	SrcAtopCompositeOp          CompositeOperator = C.SrcAtopCompositeOp
	SrcCompositeOp              CompositeOperator = C.SrcCompositeOp
	SrcInCompositeOp            CompositeOperator = C.SrcInCompositeOp
	SrcOutCompositeOp           CompositeOperator = C.SrcOutCompositeOp
	SrcOverCompositeOp          CompositeOperator = C.SrcOverCompositeOp
	ModulusSubtractCompositeOp  CompositeOperator = C.ModulusSubtractCompositeOp
	ThresholdCompositeOp        CompositeOperator = C.ThresholdCompositeOp
	XorCompositeOp              CompositeOperator = C.XorCompositeOp
	DivideDstCompositeOp        CompositeOperator = C.DivideDstCompositeOp
	DistortCompositeOp          CompositeOperator = C.DistortCompositeOp
	BlurCompositeOp             CompositeOperator = C.BlurCompositeOp
	PegtopLightCompositeOp      CompositeOperator = C.PegtopLightCompositeOp
	VividLightCompositeOp       CompositeOperator = C.VividLightCompositeOp
	PinLightCompositeOp         CompositeOperator = C.PinLightCompositeOp
	LinearDodgeCompositeOp      CompositeOperator = C.LinearDodgeCompositeOp
	LinearBurnCompositeOp       CompositeOperator = C.LinearBurnCompositeOp
	MathematicsCompositeOp      CompositeOperator = C.MathematicsCompositeOp
	DivideSrcCompositeOp        CompositeOperator = C.DivideSrcCompositeOp
	MinusSrcCompositeOp         CompositeOperator = C.MinusSrcCompositeOp
	DarkenIntensityCompositeOp  CompositeOperator = C.DarkenIntensityCompositeOp
	LightenIntensityCompositeOp CompositeOperator = C.LightenIntensityCompositeOp
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
