// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

import (
	"fmt"
	"reflect"
	"runtime"
	"sync/atomic"
	"testing"
	"time"
)

var (
	mw *MagickWand
)

func init() {
	Initialize()
}

func TestNewMagickWand(t *testing.T) {
	Initialize()
	defer func(t *testing.T) {
		checkGC(t)
	}(t)
	defer Terminate()

	mw := NewMagickWand()

	if !mw.IsVerified() {
		t.Fatal("MagickWand not verified")
	}
}

func TestCloningAndDestroying(t *testing.T) {
	Initialize()
	defer func(t *testing.T) {
		checkGC(t)
	}(t)
	defer Terminate()

	mw := NewMagickWand()
	clone := mw.Clone()
	if !clone.IsVerified() {
		t.Fatal("Unsuccessful clone")
	}

	clone = nil
	runtime.GC()
	time.Sleep(100 * time.Millisecond)

	if !mw.IsVerified() {
		t.Fatal("MagickWand not properly destroyed")
	}
}

func TestQueryConfigureOptions(t *testing.T) {
	Initialize()
	defer func(t *testing.T) {
		checkGC(t)
	}(t)
	defer Terminate()

	opts := mw.QueryConfigureOptions("*")
	if len(opts) == 0 {
		t.Fatal("QueryConfigureOptions returned an empty array")
	}
	for _, opt := range opts {
		mw.QueryConfigureOption(opt)
	}
}

func TestNonExistingConfigureOption(t *testing.T) {
	Initialize()
	defer func(t *testing.T) {
		checkGC(t)
	}(t)
	defer Terminate()

	_, err := mw.QueryConfigureOption("4321foobaramps1234")
	if err == nil {
		t.Fatal("Missing error when trying to get non-existing configure option")
	}
}

func TestQueryFonts(t *testing.T) {
	Initialize()
	defer func(t *testing.T) {
		checkGC(t)
	}(t)
	defer Terminate()

	fonts := mw.QueryFonts("*")
	if len(fonts) == 0 {
		t.Fatal("ImageMagick have not identified a single font in this system")
	}
}

func TestQueryFormats(t *testing.T) {
	Initialize()
	defer func(t *testing.T) {
		checkGC(t)
	}(t)
	defer Terminate()

	formats := mw.QueryFormats("*")
	if len(formats) == 0 {
		t.Fatal("ImageMagick have not identified a single image format in this system")
	}
}

func TestDeleteImageArtifact(t *testing.T) {
	Initialize()
	defer func(t *testing.T) {
		checkGC(t)
	}(t)
	defer Terminate()

	mw := NewMagickWand()
	mw.ReadImage(`logo:`)

	if err := mw.DeleteImageArtifact("*"); err != nil {
		t.Fatalf("Error calling DeleteImageArtifact: %s", err.Error())
	}
}

func TestImageAlpha(t *testing.T) {
	mw := NewMagickWand()
	mw.ReadImage(`logo:`)

	if mw.GetImageAlphaChannel() {
		t.Fatal("Expected image not to have an activated alpha channel")
	}

	mw.SetImageAlphaChannel(ALPHA_CHANNEL_ACTIVATE)

	if !mw.GetImageAlphaChannel() {
		t.Fatal("Expected image to have an activated alpha channel")
	}

	if err := mw.SetImageAlpha(.5); err != nil {
		t.Fatal(err)
	}
}

func TestImageMask(t *testing.T) {
	mw := NewMagickWand()
	mw.ReadImage(`logo:`)

	mask := NewMagickWand()
	mask.ReadImage(`logo:`)

	if err := mw.SetImageMask(PIXEL_MASK_READ, mask); err != nil {
		t.Fatal(err)
	}
	if err := mw.SetImageMask(PIXEL_MASK_WRITE, mask); err != nil {
		t.Fatal(err)
	}

}

func TestImageChannelMask(t *testing.T) {
	mw := NewMagickWand()
	mw.ReadImage(`logo:`)

	channel := mw.SetImageChannelMask(CHANNEL_ALPHA)
	if channel != CHANNELS_ALL {
		t.Fatalf("Expected CHANNELS_ALL (%v), got %v", CHANNELS_ALL, channel)
	}

	channel = mw.SetImageChannelMask(CHANNELS_ALL)
	if channel != CHANNEL_ALPHA {
		t.Fatalf("Expected CHANNEL_ALPHA (%v), got %v", CHANNEL_ALPHA, channel)
	}
}

func TestReadImageBlob(t *testing.T) {
	Initialize()
	defer func(t *testing.T) {
		checkGC(t)
	}(t)
	defer Terminate()

	mw := NewMagickWand()

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

func TestExportImagePixels(t *testing.T) {
	Initialize()
	defer func(t *testing.T) {
		checkGC(t)
	}(t)
	defer Terminate()

	mw := NewMagickWand()

	var err error
	if err = mw.ReadImage(`logo:`); err != nil {
		t.Fatal("Failed to read internal logo: image")
	}

	var width, height uint = 100, 100
	err = mw.ScaleImage(width, height)
	if err != nil {
		t.Fatalf("Failed scaling image: %s", err.Error())
	}

	tests := []struct {
		x, y    int
		w, h    uint
		pmap    string
		stype   StorageType
		isError bool // If we expect to fail
	}{
		{0, 0, width, height, "RGB", PIXEL_FLOAT, false},   // 0
		{0, 0, width, height, "RGBA", PIXEL_DOUBLE, false}, // 1
		{0, 0, width, height, "R", PIXEL_FLOAT, false},     // 2
		{0, 0, width, height, "GB", PIXEL_FLOAT, false},    // 3
		{0, 1, width, 1, "RGB", PIXEL_DOUBLE, false},       // 4

		{0, 0, width, 0, "RGB", PIXEL_DOUBLE, true},                    // 5
		{0, 0, 0, height, "RGB", PIXEL_DOUBLE, true},                   // 6
		{0, 0, 0, 0, "RGB", PIXEL_DOUBLE, true},                        // 7
		{int(width) + 1, 0, width, height, "RGB", PIXEL_DOUBLE, true},  // 8
		{0, int(height) + 1, width, height, "RGB", PIXEL_DOUBLE, true}, // 9
		{0, int(height) + 1, width, 1, "RGB", PIXEL_DOUBLE, true},      // 10
	}

	for i, tt := range tests {

		val, err := mw.ExportImagePixels(tt.x, tt.y, tt.w, tt.h, tt.pmap, tt.stype)

		if tt.isError {
			if err == nil {
				t.Errorf("[#%d] Expected test to fail, but it did not", i)
			}
			continue
		}

		if err != nil {
			t.Errorf("[#%d] %s", i, err.Error())
			continue
		}

		typ := reflect.TypeOf(val)
		elem := typ.Elem()
		actualKind := elem.Kind()

		var (
			expectKind reflect.Kind
			size       int
		)

		switch tt.stype {

		case PIXEL_FLOAT:
			expectKind = reflect.Float32
			size = len(val.([]float32))

		case PIXEL_DOUBLE:
			expectKind = reflect.Float64
			size = len(val.([]float64))

		default:
			t.Fatalf("[#%d] Unhandled testcase StorageType: %v", i, tt.stype)
		}

		if actualKind != expectKind {
			t.Errorf("[#%d] Expected slice element type %v, got %v",
				i, expectKind, actualKind)
			continue
		}

		expected := int(tt.w) * int(tt.h) * len(tt.pmap)

		if size != expected {
			t.Errorf("[#%d] Expected %q image to have %d pixel vals; Got %d",
				i, tt.pmap, expected, size)
		}

	}
}

func TestGetQuantumDepth(t *testing.T) {
	Initialize()
	defer func(t *testing.T) {
		checkGC(t)
	}(t)
	defer Terminate()

	name, depth := GetQuantumDepth()
	if name == "" {
		t.Fatal("Depth name returned was an empty string")
	}
	if depth == 0 {
		t.Fatal("Depth value returned was 0")
	}
}

func TestGetQuantumRange(t *testing.T) {
	Initialize()
	defer func(t *testing.T) {
		checkGC(t)
	}(t)
	defer Terminate()

	name, r := GetQuantumRange()
	if name == "" {
		t.Fatal("Depth name returned was an empty string")
	}
	if r == 0 {
		t.Fatal("Range value returned was 0")
	}
}

func checkGC(t *testing.T) {
	if !isImageMagickCleaned() {
		t.Fatal("Some ImageMagick objects are not destroyed", getObjectCountersString())
	}
}

func getObjectCountersString() string {
	str := fmt.Sprintf("\nmagickWandCounter %d\n", atomic.LoadInt64(&magickWandCounter))
	str += fmt.Sprintf("drawingWandCounter %d\n", atomic.LoadInt64(&drawingWandCounter))
	str += fmt.Sprintf("pixelIteratorCounter %d\n", atomic.LoadInt64(&pixelIteratorCounter))
	str += fmt.Sprintf("pixelWandCounter %d\n", atomic.LoadInt64(&pixelWandCounter))

	return str
}

func BenchmarkExportImagePixels(b *testing.B) {
	wand := NewMagickWand()

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
		{[]int64{0}, PIXEL_LONGLONG},
		{[]int32{0}, PIXEL_LONG},
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
