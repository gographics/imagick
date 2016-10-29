// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickCore/MagickCore.h>
*/
import "C"
import "runtime"

type ImageInfo struct {
	info *C.ImageInfo
}

func newImageInfo() *ImageInfo {
	ptr := C.AcquireImageInfo()
	C.GetImageInfo(ptr)
	imageInfo := &ImageInfo{ptr}
	runtime.SetFinalizer(imageInfo, Destroy)
	return imageInfo
}

// Destroy the ImageInfo immediately.
// This will also be called automatically during garbage collection.
func (ii *ImageInfo) Destroy() {
	if ii.info != nil {
		C.DestroyImageInfo(ii.info)
		ii.info = nil
	}
}
