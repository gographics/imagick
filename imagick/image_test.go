// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

import (
	"testing"
)

func init() {
	Initialize()
}

// TODO(justinfx): Need to expand test to try and
// produce and catch an error from NewPixelWand
func TestNewMagickImage(t *testing.T) {
	info := newImageInfo()
	defer info.Destroy()

	pw := NewPixelWand()
	pw.SetColor("blue")

	pi := pw.GetMagickColor()

	_, err := NewMagickImage(info, 100, 100, pi)
	if err != nil {
		t.Fatal(err.Error())
	}
}
