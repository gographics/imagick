// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

import (
	"testing"
)

var (
	mw *MagickWand
)

func Init() {
	Initialize()
}

func TestNewMagickWand(t *testing.T) {
	mw = NewMagickWand()
	defer mw.Destroy()

	if !mw.IsVerified() {
		t.Fatal("MagickWand not verified")
	}
}

func TestCloningAndDestroying(t *testing.T) {
	clone := mw.Clone()
	if !clone.IsVerified() {
		t.Fatal("Unsuccessful clone")
	}
	clone.Destroy()
	if clone.IsVerified() || !mw.IsVerified() {
		t.Fatal("MagickWand not properly destroyed")
	}
}

func TestQueryConfigureOptions(t *testing.T) {
	opts := mw.QueryConfigureOptions("*")
	if len(opts) == 0 {
		t.Fatal("QueryConfigureOptions returned an empty array")
	}
	for _, opt := range opts {
		mw.QueryConfigureOption(opt)
	}
}

func TestNonExistingConfigureOption(t *testing.T) {
	_, err := mw.QueryConfigureOption("4321foobaramps1234")
	if err == nil {
		t.Fatal("Missing error when trying to get non-existing configure option")
	}
}

func TestQueryFonts(t *testing.T) {
	fonts := mw.QueryFonts("*")
	if len(fonts) == 0 {
		t.Fatal("ImageMagick have not identified a single font in this system")
	}
}

func TestQueryFormats(t *testing.T) {
	formats := mw.QueryFormats("*")
	if len(formats) == 0 {
		t.Fatal("ImageMagick have not identified a single image format in this system")
	}
}

func TestDeleteImageArtifact(t *testing.T) {
	err := mw.DeleteImageArtifact("*")
	t.Log(err.Error())
}

func TestGetImageFloats(t *testing.T) {
	Initialize()
	mw = NewMagickWand()
	defer mw.Destroy()

	if !mw.IsVerified() {
		t.Fatal("MagickWand not verified")
	}

	var err error
	if err = mw.ReadImage(`logo:`); err != nil {
		t.Fatal("Failed to read internal logo: image")
	}

	width, height := int(mw.GetImageWidth()), int(mw.GetImageHeight())

	pixels := mw.GetImageFloats(FLOAT_FORMAT_RGB)
	actual := len(pixels)
	expected := (width * height * 3)
	if actual != expected {
		t.Fatalf("Expected RGB image to have %d float vals; Got %d", expected, actual)
	}

	pixels = mw.GetImageFloats(FLOAT_FORMAT_RGBA)
	actual = len(pixels)
	expected = (width * height * 4)
	if actual != expected {
		t.Fatalf("Expected RGBA image to have %d float vals; Got %d", expected, actual)
	}

	pixels = mw.GetImageFloats(FLOAT_FORMAT_R)
	actual = len(pixels)
	expected = (width * height * 1)
	if actual != expected {
		t.Fatalf("Expected RNN image to have %d float vals; Got %d", expected, actual)
	}

	pixels = mw.GetImageFloats(FLOAT_FORMAT_G | FLOAT_FORMAT_B)
	actual = len(pixels)
	expected = (width * height * 2)
	if actual != expected {
		t.Fatalf("Expected NGB image to have %d float vals; Got %d", expected, actual)
	}

}
