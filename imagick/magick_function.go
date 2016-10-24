// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type MagickFunction int

const (
	FUNCTION_UNDEFINED  MagickFunction = C.UndefinedFunction
	FUNCTION_POLYNOMIAL MagickFunction = C.PolynomialFunction
	FUNCTION_SINUSOID   MagickFunction = C.SinusoidFunction
	FUNCTION_ARCSIN     MagickFunction = C.ArcsinFunction
	FUNCTION_ARCTAN     MagickFunction = C.ArctanFunction
)
