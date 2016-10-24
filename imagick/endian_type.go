// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type EndianType int

const (
	ENDIAN_UNDEFINED EndianType = C.UndefinedEndian
	ENDIAN_LSB       EndianType = C.LSBEndian
	ENDIAN_MSB       EndianType = C.MSBEndian
)
