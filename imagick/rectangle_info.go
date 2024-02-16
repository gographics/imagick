// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickCore/MagickCore.h>
*/
import "C"

type RectangleInfo struct {
	X      int
	Y      int
	Width  uint
	Height uint
}

// Create a new RectangleInfo wrapper around a C RectangleInfo ptr
func newRectangleInfo(rectInfo *C.RectangleInfo) *RectangleInfo {
	if rectInfo == nil {
		return nil
	}
	return &RectangleInfo{
		X:      int(rectInfo.x),
		Y:      int(rectInfo.y),
		Width:  uint(rectInfo.width),
		Height: uint(rectInfo.height),
	}
}
