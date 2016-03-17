// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

import (
	"reflect"
	"testing"
)

var (
	mw *MagickWand
)

func Init() {
	Initialize()
}

func TestNewMagickWand(t *testing.T) {
	mw := NewMagickWand()
	defer mw.Destroy()

	if !mw.IsVerified() {
		t.Fatal("MagickWand not verified")
	}
}

func TestCloningAndDestroying(t *testing.T) {
	mw := NewMagickWand()
	defer mw.Destroy()
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
	mw := NewMagickWand()
	defer mw.Destroy()

	mw.ReadImage(`logo:`)

	if err := mw.DeleteImageArtifact("*"); err != nil {
		t.Fatalf("Error calling DeleteImageArtifact: %s", err.Error())
	}
}

func TestReadImageBlob(t *testing.T) {
	mw := NewMagickWand()
	defer mw.Destroy()

	// Read an invalid blob
	blob := []byte{}
	if err := mw.ReadImageBlob(blob); err == nil {
		t.Fatal("Expected a failure when passing a zero length blob")
	}

	mw.ReadImage(`logo:`)
	blob = mw.GetImageBlob()

	// Read a valid blob
	if err := mw.ReadImageBlob(blob); err != nil {
		t.Fatal(err.Error())
	}
}

func TestGetImageFloats(t *testing.T) {
	Initialize()
	mw := NewMagickWand()
	defer mw.Destroy()

	var err error
	if err = mw.ReadImage(`logo:`); err != nil {
		t.Fatal("Failed to read internal logo: image")
	}

	width, height := mw.GetImageWidth(), mw.GetImageHeight()

	val, err := mw.ExportImagePixels(0, 0, width, height, "RGB", PIXEL_FLOAT)
	if err != nil {
		t.Fatal(err.Error())
	}
	pixels := val.([]float32)
	actual := len(pixels)
	expected := (width * height * 3)
	if actual != int(expected) {
		t.Fatalf("Expected RGB image to have %d float vals; Got %d", expected, actual)
	}

	val, err = mw.ExportImagePixels(0, 0, width, height, "RGBA", PIXEL_DOUBLE)
	if err != nil {
		t.Fatal(err.Error())
	}
	pixels64 := val.([]float64)
	actual = len(pixels64)
	expected = (width * height * 4)
	if actual != int(expected) {
		t.Fatalf("Expected RGBA image to have %d float vals; Got %d", expected, actual)
	}

	val, err = mw.ExportImagePixels(0, 0, width, height, "R", PIXEL_FLOAT)
	if err != nil {
		t.Fatal(err.Error())
	}
	pixels = val.([]float32)
	actual = len(pixels)
	expected = (width * height * 1)
	if actual != int(expected) {
		t.Fatalf("Expected RNN image to have %d float vals; Got %d", expected, actual)
	}

	val, err = mw.ExportImagePixels(0, 0, width, height, "GB", PIXEL_FLOAT)
	if err != nil {
		t.Fatal(err.Error())
	}
	pixels = val.([]float32)
	actual = len(pixels)
	expected = (width * height * 2)
	if actual != int(expected) {
		t.Fatalf("Expected NGB image to have %d float vals; Got %d", expected, actual)
	}
}

func TestGetQuantumDepth(t *testing.T) {
	name, depth := GetQuantumDepth()
	if name == "" {
		t.Fatal("Depth name returned was an empty string")
	}
	if depth == 0 {
		t.Fatal("Depth value returned was 0")
	}
}

func TestGetQuantumRange(t *testing.T) {
	name, r := GetQuantumRange()
	if name == "" {
		t.Fatal("Depth name returned was an empty string")
	}
	if r == 0 {
		t.Fatal("Range value returned was 0")
	}
}

func BenchmarkExportImagePixels(b *testing.B) {
	wand := NewMagickWand()
	defer wand.Destroy()

	wand.ReadImage("logo:")
	wand.ScaleImage(1024, 1024)

	var val interface{}
	var pixels []float32

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		val, _ = wand.ExportImagePixels(0, 0, 1024, 1024, "RGB", PIXEL_FLOAT)
		pixels = val.([]float32)
	}

	b.StopTimer()

	if len(pixels) == 0 {
		b.Fatal("Pixel slice is 0")
	}
}

func BenchmarkImportImagePixels(b *testing.B) {
	wand := NewMagickWand()
	defer wand.Destroy()

	wand.ReadImage("logo:")
	wand.ScaleImage(1024, 1024)

	val, _ := wand.ExportImagePixels(0, 0, 1024, 1024, "RGB", PIXEL_FLOAT)
	pixels := val.([]float32)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		wand.ImportImagePixels(0, 0, 1024, 1024, "RGB", PIXEL_UNDEFINED, pixels)
	}

	b.StopTimer()
}

func TestPixelInterfaceToPtr(t *testing.T) {
	tests := []struct {
		pixels  interface{}
		storage StorageType
	}{
		{[]byte{0}, PIXEL_CHAR},
		{[]float64{0}, PIXEL_DOUBLE},
		{[]float32{0}, PIXEL_FLOAT},
		{[]int16{0}, PIXEL_SHORT},
		{[]int32{0}, PIXEL_INTEGER},
		{[]int64{0}, PIXEL_LONG},
	}
	for _, value := range tests {
		_, storageType, err := pixelInterfaceToPtr(value.pixels)
		if err != nil {
			t.Fatal("Error when passing", reflect.TypeOf(value.pixels))
		}
		if storageType != value.storage {
			t.Fatal("Wrong storage type received for", reflect.TypeOf(value.pixels))
		}
	}

	_, _, err := pixelInterfaceToPtr(32)
	if err == nil {
		t.Fatal("Expected error when passing invalid type")
	}
}
