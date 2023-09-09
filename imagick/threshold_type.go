// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type AutoThresholdMethodType int

const (
	AUTO_THRESHOLD_METHOD_UNDEFINED AutoThresholdMethodType = C.UndefinedThresholdMethod
	AUTO_THRESHOLD_METHOD_KAPUR     AutoThresholdMethodType = C.KapurThresholdMethod
	AUTO_THRESHOLD_METHOD_OTSU      AutoThresholdMethodType = C.OTSUThresholdMethod
	AUTO_THRESHOLD_METHOD_TRIANGLE  AutoThresholdMethodType = C.TriangleThresholdMethod
)
