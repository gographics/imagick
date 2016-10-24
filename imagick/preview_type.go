// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type PreviewType int

const (
	PREVIEW_UNDEFINED        PreviewType = C.UndefinedPreview
	PREVIEW_ROTATE           PreviewType = C.RotatePreview
	PREVIEW_SHEAR            PreviewType = C.ShearPreview
	PREVIEW_ROLL             PreviewType = C.RollPreview
	PREVIEW_HUE              PreviewType = C.HuePreview
	PREVIEW_SATURATION       PreviewType = C.SaturationPreview
	PREVIEW_BRIGHTNESS       PreviewType = C.BrightnessPreview
	PREVIEW_GAMMA            PreviewType = C.GammaPreview
	PREVIEW_SPIFF            PreviewType = C.SpiffPreview
	PREVIEW_DULL             PreviewType = C.DullPreview
	PREVIEW_GRAYSCALE        PreviewType = C.GrayscalePreview
	PREVIEW_QUANTIZE         PreviewType = C.QuantizePreview
	PREVIEW_DESPECKLE        PreviewType = C.DespecklePreview
	PREVIEW_REDUCE_NOISE     PreviewType = C.ReduceNoisePreview
	PREVIEW_ADD_NOISE        PreviewType = C.AddNoisePreview
	PREVIEW_SHARPEN          PreviewType = C.SharpenPreview
	PREVIEW_BLUR             PreviewType = C.BlurPreview
	PREVIEW_THRESHOLD        PreviewType = C.ThresholdPreview
	PREVIEW_EDGE_DETECT      PreviewType = C.EdgeDetectPreview
	PREVIEW_SPREAD           PreviewType = C.SpreadPreview
	PREVIEW_SOLARIZE         PreviewType = C.SolarizePreview
	PREVIEW_SHADE            PreviewType = C.ShadePreview
	PREVIEW_RAISE            PreviewType = C.RaisePreview
	PREVIEW_SEGMENT          PreviewType = C.SegmentPreview
	PREVIEW_SWIRL            PreviewType = C.SwirlPreview
	PREVIEW_IMPLODE          PreviewType = C.ImplodePreview
	PREVIEW_WAVE             PreviewType = C.WavePreview
	PREVIEW_OIL_PAINT        PreviewType = C.OilPaintPreview
	PREVIEW_CHARCOAL_DRAWING PreviewType = C.CharcoalDrawingPreview
	PREVIEW_JPEG             PreviewType = C.JPEGPreview
)
