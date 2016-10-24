// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickCore/MagickCore.h>
*/
import "C"

// TODO(justinfx)
// KernelInfo contains fields to describe a convolution kernel
type KernelInfo struct {
	info *C.KernelInfo
}

func newKernelInfoFromCAPI(info *C.KernelInfo) *KernelInfo {
	return &KernelInfo{info}
}
