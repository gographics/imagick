package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
)

type PreviewType int

const (
	UndefinedPreview PreviewType = iota
	RotatePreview
	ShearPreview
	RollPreview
	HuePreview
	SaturationPreview
	BrightnessPreview
	GammaPreview
	SpiffPreview
	DullPreview
	GrayscalePreview
	QuantizePreview
	DespecklePreview
	ReduceNoisePreview
	AddNoisePreview
	SharpenPreview
	BlurPreview
	ThresholdPreview
	EdgeDetectPreview
	SpreadPreview
	SolarizePreview
	ShadePreview
	RaisePreview
	SegmentPreview
	SwirlPreview
	ImplodePreview
	WavePreview
	OilPaintPreview
	CharcoalDrawingPreview
	JPEGPreview
)

var previewTypeStrings = map[PreviewType]string{
	UndefinedPreview:       "UndefinedPreview",
	RotatePreview:          "RotatePreview",
	ShearPreview:           "ShearPreview",
	RollPreview:            "RollPreview",
	HuePreview:             "HuePreview",
	SaturationPreview:      "SaturationPreview",
	BrightnessPreview:      "BrightnessPreview",
	GammaPreview:           "GammaPreview",
	SpiffPreview:           "SpiffPreview",
	DullPreview:            "DullPreview",
	GrayscalePreview:       "GrayscalePreview",
	QuantizePreview:        "QuantizePreview",
	DespecklePreview:       "DespecklePreview",
	ReduceNoisePreview:     "ReduceNoisePreview",
	AddNoisePreview:        "AddNoisePreview",
	SharpenPreview:         "SharpenPreview",
	BlurPreview:            "BlurPreview",
	ThresholdPreview:       "ThresholdPreview",
	EdgeDetectPreview:      "EdgeDetectPreview",
	SpreadPreview:          "SpreadPreview",
	SolarizePreview:        "SolarizePreview",
	ShadePreview:           "ShadePreview",
	RaisePreview:           "RaisePreview",
	SegmentPreview:         "SegmentPreview",
	SwirlPreview:           "SwirlPreview",
	ImplodePreview:         "ImplodePreview",
	WavePreview:            "WavePreview",
	OilPaintPreview:        "OilPaintPreview",
	CharcoalDrawingPreview: "CharcoalDrawingPreview",
	JPEGPreview:            "JPEGPreview",
}

func (cst *PreviewType) String() string {
	if v, ok := previewTypeStrings[PreviewType(*cst)]; ok {
		return v
	}
	return fmt.Sprintf("UnknownPreviewType[%d]", *cst)
}
