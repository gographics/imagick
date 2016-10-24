// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickCore/MagickCore.h>
*/
import "C"

type Image struct {
	img *C.Image
}

// TODO(justinfx): Needs tests, using ExceptionInfo
func NewMagickImage(info *ImageInfo, width, height uint,
	background *PixelInfo) (*Image, *ExceptionInfo) {

	var exc C.ExceptionInfo

	img := C.NewMagickImage(
		info.info,
		C.size_t(width), C.size_t(height),
		background.pi,
		&exc)

	var err *ExceptionInfo = nil

	if exc.error_number != 0 {
		err = newExceptionInfo(&exc)
	}

	return &Image{img: img}, err
}
