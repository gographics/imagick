// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

import "testing"

func init() {
	Initialize()
}

func TestKernelInfo(t *testing.T) {
	mw := NewMagickWand()
	mw.SetSize(640, 480)
	mw.ReadImage("magick:logo")
	mw.SetFormat("png")

	table := []struct {
		kernel string
		size   int
	}{
		{"Laplacian:2", 3},                                      // Laplacian
		{"3x3: 0.0, 0.5, 0.0, 0.5, 1.0, 0.5, 0.0, 0.5, 0.0", 3}, // spread
		{"0.0 0.5 0.0 0.5 1.0 0.5 0.0 0.5 0.0", 3},              // old
	}

	var err error
	var kernel_info *KernelInfo

	for i, tt := range table {
		kernel_info = NewKernelInfo(tt.kernel)
		if kernel_info == nil {
			t.Fatalf("NewKernelInfo failed to init (test #%d)", i)
		}

		kernel_values := kernel_info.ToArray()

		expected := tt.size
		actual := len(kernel_values)
		if actual != expected {
			t.Fatalf("Expected %d rows of values, got %d", expected, actual)
		}

		for i, row := range kernel_values {
			actual = len(row)
			if actual != expected {
				t.Fatalf("Row %d of %d: Expected %d value, got %d",
					i, expected, expected, actual)
			}
		}

		if err = mw.MorphologyImage(MORPHOLOGY_CONVOLVE, 1, kernel_info); err != nil {
			t.Errorf("Convolve failed for %q: %s", tt.kernel, err.Error())
		}
	}
}

func TestKernelInfoBuiltIn(t *testing.T) {
	mw := NewMagickWand()
	mw.SetSize(640, 480)
	mw.ReadImage("magick:logo")
	mw.SetFormat("png")

	kernel_info := NewKernelInfoBuiltIn(KERNEL_RING, "2,1")
	if kernel_info == nil {
		t.Fatal("NewKernelInfoBuiltIn failed to init")
	}

	kernel_info.Scale(1.0, KERNEL_NORMALIZE_VALUE)
	kernel_values := kernel_info.ToArray()

	expected := 5
	actual := len(kernel_values)
	if actual != expected {
		t.Fatalf("Expected %d rows of values, got %d", expected, actual)
	}

	for i, row := range kernel_values {
		actual = len(row)
		if actual != expected {
			t.Fatalf("Row %d of %d: Expected %d value, got %d",
				i, expected, expected, actual)
		}
	}

	if err2 := mw.MorphologyImage(MORPHOLOGY_CONVOLVE, 1, kernel_info); err2 != nil {
		t.Fatalf("Convolve failed: %s", err2.Error())
	}
}
