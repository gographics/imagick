// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickCore/MagickCore.h>
*/
import "C"

type ImageInfo struct {
	info *C.ImageInfo
}

func newImageInfo() *ImageInfo {
	ii := C.AcquireImageInfo()
	C.GetImageInfo(ii)
	return &ImageInfo{ii}
}

func (ii *ImageInfo) Destroy() {
	if ii.info != nil {
		C.DestroyImageInfo(ii.info)
		ii.info = nil
	}
}
