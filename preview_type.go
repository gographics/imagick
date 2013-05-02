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
	UndefinedPreview       PreviewType = C.UndefinedPreview
	RotatePreview          PreviewType = C.RotatePreview
	ShearPreview           PreviewType = C.ShearPreview
	RollPreview            PreviewType = C.RollPreview
	HuePreview             PreviewType = C.HuePreview
	SaturationPreview      PreviewType = C.SaturationPreview
	BrightnessPreview      PreviewType = C.BrightnessPreview
	GammaPreview           PreviewType = C.GammaPreview
	SpiffPreview           PreviewType = C.SpiffPreview
	DullPreview            PreviewType = C.DullPreview
	GrayscalePreview       PreviewType = C.GrayscalePreview
	QuantizePreview        PreviewType = C.QuantizePreview
	DespecklePreview       PreviewType = C.DespecklePreview
	ReduceNoisePreview     PreviewType = C.ReduceNoisePreview
	AddNoisePreview        PreviewType = C.AddNoisePreview
	SharpenPreview         PreviewType = C.SharpenPreview
	BlurPreview            PreviewType = C.BlurPreview
	ThresholdPreview       PreviewType = C.ThresholdPreview
	EdgeDetectPreview      PreviewType = C.EdgeDetectPreview
	SpreadPreview          PreviewType = C.SpreadPreview
	SolarizePreview        PreviewType = C.SolarizePreview
	ShadePreview           PreviewType = C.ShadePreview
	RaisePreview           PreviewType = C.RaisePreview
	SegmentPreview         PreviewType = C.SegmentPreview
	SwirlPreview           PreviewType = C.SwirlPreview
	ImplodePreview         PreviewType = C.ImplodePreview
	WavePreview            PreviewType = C.WavePreview
	OilPaintPreview        PreviewType = C.OilPaintPreview
	CharcoalDrawingPreview PreviewType = C.CharcoalDrawingPreview
	JPEGPreview            PreviewType = C.JPEGPreview
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
