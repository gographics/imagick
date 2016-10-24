// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type StyleType int

const (
	STYLE_UNDEFINED StyleType = C.UndefinedStyle
	STYLE_NORMAL    StyleType = C.NormalStyle
	STYLE_ITALIC    StyleType = C.ItalicStyle
	STYLE_OBLIQUE   StyleType = C.ObliqueStyle
	STYLE_ANYSTYLE  StyleType = C.AnyStyle
)
