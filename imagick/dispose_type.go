// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type DisposeType int

const (
	DISPOSE_UNRECOGNIZED DisposeType = C.UnrecognizedDispose
	DISPOSE_UNDEFINED    DisposeType = C.UndefinedDispose
	DISPOSE_NONE         DisposeType = C.NoneDispose
	DISPOSE_BACKGROUND   DisposeType = C.BackgroundDispose
	DISPOSE_PREVIOUS     DisposeType = C.PreviousDispose
)
