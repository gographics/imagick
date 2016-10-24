// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type ClipPathUnits int

const (
	CLIP_UNDEFINED_PATH      ClipPathUnits = C.UndefinedPathUnits
	CLIP_USER_SPACE          ClipPathUnits = C.UserSpace
	CLIP_USER_SPACE_ON_USE   ClipPathUnits = C.UserSpaceOnUse
	CLIP_OBJECT_BOUNDING_BOX ClipPathUnits = C.ObjectBoundingBox
)
