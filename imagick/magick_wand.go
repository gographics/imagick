package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

// This struct represents the MagickWand C API of ImageMagick
type MagickWand struct {
	mw *C.MagickWand
}

// Returns a wand required for all other methods in the API. A fatal exception is thrown if there is not enough memory to allocate the wand.
// Use Destroy() to dispose of the wand then it is no longer needed.
func NewMagickWand() *MagickWand {
	return &MagickWand{C.NewMagickWand()}
}

// Returns a wand with an image/
func NewMagickWandFromImage(img *Image) *MagickWand {
	return &MagickWand{C.NewMagickWandFromImage(img.img)}
}

// Clear resources associated with the wand, leaving the wand blank, and ready to be used for a new set of images.
func (mw *MagickWand) Clear() {
	C.ClearMagickWand(mw.mw)
}

// Makes an exact copy of the MagickWand object
func (mw *MagickWand) Clone() *MagickWand {
	clone := C.CloneMagickWand(mw.mw)
	return &MagickWand{clone}
}

// Deallocates memory associated with an MagickWand
func (mw *MagickWand) Destroy() {
	mw.mw = C.DestroyMagickWand(mw.mw)
	C.free(unsafe.Pointer(mw.mw))
	mw.mw = nil
}

// Returns true if the wand is a verified magick wand
func (mw *MagickWand) IsVerified() bool {
	if mw.mw != nil {
		return 1 == C.int(C.IsMagickWand(mw.mw))
	}
	return false
}

// Returns the position of the iterator in the image list
func (mw *MagickWand) GetIteratorIndex() uint {
	return uint(C.MagickGetIteratorIndex(mw.mw))
}

// Returns the value associated with the specified configure option
func (mw *MagickWand) QueryConfigureOption(option string) (string, error) {
	csoption := C.CString(option)
	defer C.free(unsafe.Pointer(csoption))
	availableOptions := mw.QueryConfigureOptions(option)
	for _, availableOption := range availableOptions {
		if availableOption == option {
			return C.GoString(C.MagickQueryConfigureOption(csoption)), nil
		}
	}
	return "", fmt.Errorf("Unknown option \"%s\"", option)
}

// Returns any configure options that match the specified pattern (e.g. "*" for all). Options include NAME, VERSION, LIB_VERSION, etc
func (mw *MagickWand) QueryConfigureOptions(pattern string) (options []string) {
	cspattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(cspattern))
	var num C.size_t
	copts := C.MagickQueryConfigureOptions(cspattern, &num)
	defer C.free(unsafe.Pointer(copts))
	options = sizedCStringArrayToStringSlice(copts, num)
	return
}

// Returns a FontMetrics struct
func (mw *MagickWand) QueryFontMetrics(dw *DrawingWand, textLine string) *FontMetrics {
	cstext := C.CString(textLine)
	defer C.free(unsafe.Pointer(cstext))
	cdoubles := C.MagickQueryFontMetrics(mw.mw, dw.dw, cstext)
	defer C.free(unsafe.Pointer(cdoubles))
	doubles := sizedDoubleArrayToFloat64Slice(cdoubles, 13)
	return NewFontMetricsFromArray(doubles)
}

// Returns a FontMetrics struct related to the multiline text
func (mw *MagickWand) QueryMultilineFontMetrics(dw *DrawingWand, textParagraph string) *FontMetrics {
	cstext := C.CString(textParagraph)
	defer C.free(unsafe.Pointer(cstext))
	cdoubles := C.MagickQueryMultilineFontMetrics(mw.mw, dw.dw, cstext)
	doubles := sizedDoubleArrayToFloat64Slice(cdoubles, 13)
	defer C.free(unsafe.Pointer(cdoubles))
	return NewFontMetricsFromArray(doubles)
}

// Returns any font that match the specified pattern (e.g. "*" for all)
func (mw *MagickWand) QueryFonts(pattern string) (fonts []string) {
	cspattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(cspattern))
	var num C.size_t
	copts := C.MagickQueryFonts(cspattern, &num)
	fonts = sizedCStringArrayToStringSlice(copts, num)
	return
}

// Returns any supported image format that match the specified pattern (e.g. "*" for all)
func (mw *MagickWand) QueryFormats(pattern string) (formats []string) {
	cspattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(cspattern))
	var num C.size_t
	copts := C.MagickQueryFormats(cspattern, &num)
	formats = sizedCStringArrayToStringSlice(copts, num)
	return
}

// Relinquishes memory resources returned by such methods as MagickIdentifyImage(), MagickGetException(), etc
func (mw *MagickWand) relinquishMemory(ptr unsafe.Pointer) {
	C.MagickRelinquishMemory(ptr)
}

// This method resets the wand iterator.
// It is typically used either before iterating though images, or before calling specific methods such as AppendImages()
// to append all images together.
// Afterward you can use NextImage() to iterate over all the images in a wand container, starting with the first image.
// Using this before AddImages() or ReadImages() will cause new images to be inserted between the first and second image.
func (mw *MagickWand) ResetIterator() {
	C.MagickResetIterator(mw.mw)
}

// This method sets the wand iterator to the first image.
// After using any images added to the wand using AddImage() or ReadImage() will be prepended before any image in the wand.
// Also the current image has been set to the first image (if any) in the MagickWand. Using NextImage() will then set the
// current image to the second image in the list (if present).
// This operation is similar to ResetIterator() but differs in how AddImage(), ReadImage(), and NextImage() behaves afterward.
func (mw *MagickWand) SetFirstIterator() {
	C.MagickSetFirstIterator(mw.mw)
}

// This method set the iterator to the given position in the image list specified with the index parameter.
// A zero index will set the first image as current, and so on. Negative indexes can be used to specify an
// image relative to the end of the images in the wand, with -1 being the last image in the wand.
// If the index is invalid (range too large for number of images in wand) the function will return false.
// In that case the current image will not change.
// After using any images added to the wand using AddImage() or ReadImage() will be added after the image indexed,
// regardless of if a zero (first image in list) or negative index (from end) is used.
// Jumping to index 0 is similar to ResetIterator() but differs in how NextImage() behaves afterward.
func (mw *MagickWand) SetIteratorIndex(index int) bool {
	return 1 == C.int(C.MagickSetIteratorIndex(mw.mw, C.ssize_t(index)))
}

// SetLastIterator() sets the wand iterator to the last image.
// The last image is actually the current image, and the next use of PreviousImage() will not change this allowing this function
// to be used to iterate over the images in the reverse direction. In this sense it is more like ResetIterator() than SetFirstIterator().
// Typically this function is used before AddImage(), ReadImage() functions to ensure new images are appended to the very end of wand's image list.
func (mw *MagickWand) SetLastIterator() {
	C.MagickSetLastIterator(mw.mw)
}
