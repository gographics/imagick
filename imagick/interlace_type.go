// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type InterlaceType int

const (
	INTERLACE_UNDEFINED InterlaceType = C.UndefinedInterlace
	INTERLACE_NO        InterlaceType = C.NoInterlace
	INTERLACE_LINE      InterlaceType = C.LineInterlace
	INTERLACE_PLANE     InterlaceType = C.PlaneInterlace
	INTERLACE_PARTITION InterlaceType = C.PartitionInterlace
	INTERLACE_GIF       InterlaceType = C.GIFInterlace
	INTERLACE_JPEG      InterlaceType = C.JPEGInterlace
	INTERLACE_PNG       InterlaceType = C.PNGInterlace
)
