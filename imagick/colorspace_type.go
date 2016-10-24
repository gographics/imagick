// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type ColorspaceType int

const (
	COLORSPACE_UNDEFINED   ColorspaceType = C.UndefinedColorspace
	COLORSPACE_CMY         ColorspaceType = C.CMYColorspace
	COLORSPACE_CMYK        ColorspaceType = C.CMYKColorspace
	COLORSPACE_GRAY        ColorspaceType = C.GRAYColorspace
	COLORSPACE_HCL         ColorspaceType = C.HCLColorspace
	COLORSPACE_HCLP        ColorspaceType = C.HCLpColorspace
	COLORSPACE_HSB         ColorspaceType = C.HSBColorspace
	COLORSPACE_HSI         ColorspaceType = C.HSIColorspace
	COLORSPACE_HSL         ColorspaceType = C.HSLColorspace
	COLORSPACE_HSV         ColorspaceType = C.HSVColorspace
	COLORSPACE_HWB         ColorspaceType = C.HWBColorspace
	COLORSPACE_LAB         ColorspaceType = C.LabColorspace
	COLORSPACE_LCH         ColorspaceType = C.LCHColorspace
	COLORSPACE_LCHAB       ColorspaceType = C.LCHabColorspace
	COLORSPACE_LCHUV       ColorspaceType = C.LCHuvColorspace
	COLORSPACE_LMS         ColorspaceType = C.LMSColorspace
	COLORSPACE_LOG         ColorspaceType = C.LogColorspace
	COLORSPACE_LUV         ColorspaceType = C.LuvColorspace
	COLORSPACE_OHTA        ColorspaceType = C.OHTAColorspace
	COLORSPACE_REC601YCBCR ColorspaceType = C.Rec601YCbCrColorspace
	COLORSPACE_REC709YCBCR ColorspaceType = C.Rec709YCbCrColorspace
	COLORSPACE_RGB         ColorspaceType = C.RGBColorspace
	COLORSPACE_SCRGB       ColorspaceType = C.scRGBColorspace
	COLORSPACE_SRGB        ColorspaceType = C.sRGBColorspace
	COLORSPACE_TRANSPARENT ColorspaceType = C.TransparentColorspace
	COLORSPACE_XYY         ColorspaceType = C.xyYColorspace
	COLORSPACE_XYZ         ColorspaceType = C.XYZColorspace
	COLORSPACE_YCBCR       ColorspaceType = C.YCbCrColorspace
	COLORSPACE_YCC         ColorspaceType = C.YCCColorspace
	COLORSPACE_YDDDR       ColorspaceType = C.YDbDrColorspace
	COLORSPACE_YIQ         ColorspaceType = C.YIQColorspace
	COLORSPACE_YPBPR       ColorspaceType = C.YPbPrColorspace
	COLORSPACE_YUV         ColorspaceType = C.YUVColorspace
)
