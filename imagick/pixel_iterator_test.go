// Copyright 2013-2016 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

import (
	"runtime"
	"testing"
	"time"
)

func init() {
	Initialize()
}

func TestPixelIteratorPrivatePixelWands(t *testing.T) {
	Initialize()
	defer func(t *testing.T) {
		checkGC(t)
		Terminate()
	}(t)

	mw := NewMagickWand()
	if !mw.IsVerified() {
		t.Fatal("MagickWand is not verified")
	}

	if err := mw.ReadImage(`logo:`); err != nil {
		t.Fatal(err.Error())
	}

	if mw.GetImageWidth() <= 0 || mw.GetImageHeight() <= 0 {
		t.Fatal("logo: source does not have a valid width/height")
	}

	pi := mw.NewPixelIterator()
	pis := []*PixelIterator{pi, pi.Clone()}

	for i, each := range pis {

		if !each.IsVerified() {
			t.Fatal("PixelIterator %d is not verified", i)
		}

		var j uint
		for j = 0; j < mw.GetImageHeight(); j++ {

			row := each.GetNextIteratorRow()
			if len(row) == 0 {
				t.Fatal("Got an empty row of PixelWands")
			}

			row = each.GetCurrentIteratorRow()
			if len(row) == 0 {
				t.Fatal("Got an empty row of PixelWands")
			}

			if j > 0 {
				row = each.GetPreviousIteratorRow()
				if len(row) == 0 {
					t.Fatal("Got an empty row of PixelWands")
				}
			}
		}

		pis[i] = nil
	}

	pi = nil

	runtime.GC()
	time.Sleep(100 * time.Millisecond)
}
