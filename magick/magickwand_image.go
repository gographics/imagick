package magick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

import (
//"fmt"
//"unsafe"
)

// Returns the current image from the magick wand
func (mw *MagickWand) GetImage() *Image {
	return newImageFromCAPI(C.GetImageFromMagickWand(mw.wand))
}

// Adaptively blurs the image by blurring less intensely near image edges and more intensely far
// from edges. We blur the image with a Gaussian operator of the given radius and standard deviation
// (sigma). For reasonable results, radius should be larger than sigma. Use a radius of 0 and
// AdaptiveBlurImage() selects a suitable radius for you.
// radius: the radius of the Gaussian, in pixels, not counting the center pixel
// sigma: the standard deviation of the Gaussian, in pixels
func (mw *MagickWand) AdaptiveBlurImage(radius, sigma float64) error {
	C.MagickAdaptiveBlurImage(mw.wand, C.double(radius), C.double(sigma))
	return mw.GetLastError()
}

// Adaptively blurs the image by blurring less intensely near image edges and more intensely far
// from edges. We blur the image with a Gaussian operator of the given radius and standard deviation
// (sigma). For reasonable results, radius should be larger than sigma. Use a radius of 0 and
// AdaptiveBlurImage() selects a suitable radius for you.
// radius: the radius of the Gaussian, in pixels, not counting the center pixel
// sigma: the standard deviation of the Gaussian, in pixels
func (mw *MagickWand) AdaptiveBlurImageChannel(channel ChannelType, radius, sigma float64) error {
	C.MagickAdaptiveBlurImageChannel(mw.wand, C.ChannelType(channel), C.double(radius), C.double(sigma))
	return mw.GetLastError()
}

// Adaptively resize image with data dependent triangulation
func (mw *MagickWand) AdaptiveResizeImage(columns, rows uint) error {
	C.MagickAdaptiveResizeImage(mw.wand, C.size_t(columns), C.size_t(rows))
	return mw.GetLastError()
}

// Adaptively sharpens the image by sharpening more intensely near image edges and less intensely far from edges.
// We sharpen the image with a Gaussian operator of the given radius and standard deviation (sigma). For reasonable
// results, radius should be larger than sigma. Use a radius of 0 and AdaptiveSharpenImage() selects a suitable
// radius for you.
// radius: the radius of the Gaussian, in pixels, not counting the center pixel
// sigma: the standard deviation of the Gaussian, in pixels.
func (mw *MagickWand) AdaptiveSharpenImage(radius, sigma float64) error {
	C.MAgickAdaptiveSharpenImage(mw.wand, C.double(radius), C.double(sigma))
	return mw.GetLastError()
}

// Adaptively sharpens the image by sharpening more intensely near image edges and less intensely far from edges.
// We sharpen the image with a Gaussian operator of the given radius and standard deviation (sigma). For reasonable
// results, radius should be larger than sigma. Use a radius of 0 and AdaptiveSharpenImage() selects a suitable
// radius for you.
// radius: the radius of the Gaussian, in pixels, not counting the center pixel
// sigma: the standard deviation of the Gaussian, in pixels.
func (mw *MagickWand) AdaptiveSharpenImageChannel(channel ChannelType, radius, sigma float64) error {
	C.MagickAdaptiveSharpenImage(mw.wand, C.ChannelType(channel), C.double(radius), C.double(sigma))
	return mw.GetLastError()
}

// Selects an individual threshold for each pixel based on the range of intensity values in its local neighborhood.
// This allows for thresholding of an image whose global intensity histogram doesn't contain distinctive peaks.
func (mw *MagickWand) AdaptiveThresholdImage(width, height uint, offset int) error {
	C.MagickAdaptiveThresholdImage(mw.wand, C.size_t(width), C.size_t(height), C.ssize_t(offset))
	return mw.GetLastError()
}

// Adds a clone of the images from the second wand and inserts them into the first wand.
// Use SetLastIterator(), to append new images into an existing wand, current image will be set to last image so
// later adds with also be appened to end of wand.
// Use SetFirstIterator() to prepend new images into wand, any more images added will also be prepended before other
// images in the wand. However the order of a list of new images will not change.
// Otherwise the new images will be inserted just after the current image, and any later image will also be added
// after this current image but before the previously added images. Caution is advised when multiple image adds are
// inserted into the middle of the wand image list.
func (mw *MagickWand) AddImage(wand *MagickWand) error {
	C.MagickAddImage(mw.wand, wand.wand)
	return mw.GetLastError()
}

// Adds random noise to the image
func (mw *MagickWand) AddNoiseImage(noiseType NoiseType) error {
	C.MagickAddNoiseImage(mw.wand, C.NoiseType(noiseType))
	return mw.GetLastError()
}

// Adds random noise to the image's channel
func (mw *MagickWand) AddNoiseImage(channel ChannelType, noiseType NoiseType) error {
	C.MagickAddNoiseImageChannel(mw.wand, C.ChannelType(channel), C.NoiseType(noiseType))
	return mw.GetLastError()
}

// Transforms an image as dictaded by the affine matrix of the drawing wand
func (mw *MagickWand) AffineTransformImage(drawingWand *DrawingWand) error {
	C.MagickAffineTransformImage(mw.wand, drawingWand.draw)
	return mw.GetLastError()
}

// Annotates an image with text
// x: ordinate to left of text
// y: ordinate to text baseline
// angle: rotate text relative to this angle
func (mw *MagickWand) AnnotateImage(drawingWand *DrawingWand, x, y, angle float64, text string) error {
	cstext := C.CString(text)
	defer C.free(unsafe.Pointer(cstext))
	C.MagickAnnotateImage(mw.wand, drawingWand.draw, C.double(x), C.double(y), C.double(angle), cstext)
	return mw.GetLastError()
}

// Animates an image or image sequence in X11
func (mw *MagickWand) X11AnimateImages(server string) error {
	csserver := C.CString(server)
	defer C.free(unsafe.Pointer(csserver))
	C.MagickAnimateImages(mw.wand, csserver)
	return mw.GetLastError()
}

// Append the images in a wand from the current image onwards, creating a new wand with the single image result.
// This is affected by the gravity and background setting of the first image.
// Typically you would call either ResetIterator() or SetFirstImage() before calling this function to ensure that
// all the images in the wand's image list will be appended together.
// By default, images are stacked left-to-right. Set topToBottom to true to stack them top-to-bottom.
func (mw *MagickWand) AppendImages(topToBottom bool) *MagickWand {
	stack := C.MagickBooleanType(0)
	if topToBottom {
		stack = C.MagickBooleanType(1)
	}
	return &MagickWand{wand: C.MagickAppendImages(mw.wand, stack)}
}

// Extracts the 'mean' from the image and adjust the image to try make set it's gamma appropriatally
func (mw *MagickWand) AutoGammaImage() error {
	C.MagickAutoGammaImage(mw.wand)
	return mw.GetLastError()
}

// Extracts the 'mean' from the image's channel and adjust the image to try make set it's gamma appropriatally
func (mw *MagickWand) AutoGammaImageChannel(channel ChannelType) error {
	C.MagickAutoGammaImageChannel(mw.wand, C.ChannelType(channel))
	return mw.GetLastError()
}

// Adjust the levels of a particular image by scaling the minimum and maximum values to the full quantum range.
func (mw *MagickWand) AutoLevelImage() error {
	C.MagickAutoLevelImage(mw.wand)
	return mw.GetLastError()
}

// Adjust the levels of a particular image channel by scaling the minimum and maximum values to the full quantum range.
func (mw *MagickWand) AutoLevelImage(channel ChannelType) error {
	C.MagickAutoLevelImageChannel(mw.wand, C.ChannelType(channel))
	return mw.GetLastError()
}

// This is like ThresholdImage() but forces all pixels below the threshold into black while leaving all
// pixels above the threshold unchanged.
func (mw *MagickWand) BlackThresholdImage(threshold *PixelWand) error {
	C.MagickBlackThresholdImage(mw.wand, threshold.pixel)
	return mw.GetLastError()
}

// Mutes the colors of the image to simulate a scene at nighttime in the moonlight.
func (mw *MagickWand) BlueShiftImage(factor float64) error {
	C.MagickBlueShiftImage(mw.wand, C.double(factor))
	return mw.GetLastError()
}

// Blurs an image. We convolve the image with a gaussian operator of the given radius and standard deviation (sigma).
// For reasonable results, the radius should be larger than sigma. Use a radius of 0 and BlurImage() selects a suitable
// radius for you.
// radius: the radius of the, in pixels, not counting the center pixel.
// sigma: the standard deviation of the, in pixels
func (mw *MagickWand) BlurImage(radius, sigma float64) error {
	C.MagickBlurImage(mw.wand, C.double(radius), C.double(sigma))
	return mw.GetLastError()
}

// Blurs an image's channel. We convolve the image with a gaussian operator of the given radius and standard deviation (sigma).
// For reasonable results,
// the radius should be larger than sigma. Use a radius of 0 and BlurImage() selects a suitable radius for you.
// radius: the radius of the, in pixels, not counting the center pixel.
// sigma: the standard deviation of the, in pixels
func (mw *MagickWand) BlurImageChannel(channel ChannelType, radius, sigma float64) error {
	C.MagickBlurImageChannel(mw.wand, C.ChannelType(channel), C.double(radius), C.double(sigma))
	return mw.GetLastError()
}

// Surrounds the image with a border of the color defined by the bordercolor pixel wand.
func (mw *MagickWand) BorderImage(borderColor *PixelWand, width, height uint) error {
	C.MagickBorderImage(mw.wand, borderColor.pixel, C.size_t(width), C.size_t(height))
	return mw.GetLastError()
}

// Use this to change the brightness and/or contrast of an image. It converts the brightness and contrast
// brighness: the brightness percent (-100 .. 100)
// contrast: the brightness percent (-100 .. 100)
func (mw *MagickWand) BrightnessContrastImage(brightness, contrast float64) error {
	C.MagickBrightnessContrastImage(mw.wand, C.double(brightness), C.double(contrast))
	return mw.GetLastError()
}

// Use this to change the brightness and/or contrast of an image's channel. It converts the brightness and contrast
// brighness: the brightness percent (-100 .. 100)
// contrast: the brightness percent (-100 .. 100)
func (mw *MagickWand) BrightnessContrastImageChannel(channel ChannelType, brightness, contrast float64) error {
	C.MagickBrightnessContrastImageChannel(mw.wand, C.ChannelType(channel), C.double(brightness), C.double(contrast))
	return mw.GetLastError()
}

// Simulates a charcoal drawing
// radius: the radius of the Gaussian, in pixels, not counting the center pixel
// sigma: the standard deviation of the Gaussian, in pixels
func (mw *MagickWand) CharcoalImage(radius, sigma float64) error {
	C.MagickCharcoalImage(mw.wand, C.double(radius), C.double(sigma))
	return mw.GetLastError()
}

// Removes a region of an image and collapses the image to occupy the removed portion
// width, height: the region width and height
// x, y: the region x and y offsets
func (mw *MagickWand) ChopImage(width, height uint, x, y int) error {
	C.MagickChopImage(mw.wand, C.size_t(width), C.size_t(height), C.ssize_t(x), C.ssize_t(y))
	return mw.GetLastError()
}

// Restricts the color range from 0 to the quantum depth
func (mw *MagickWand) ClampImage() error {
	C.MagickClampImage(mw.wand)
	return mw.GetLastError()
}

// Restricts the color range from 0 to the quantum depth
func (mw *MagickWand) ClampImageChannel(channel ChannelType) error {
	C.MagickClampImageChannel(mw.wand, C.ChannelType(channel))
	return mw.GetLastError()
}

// Clips along the first path from the 8BIM profile, if present
func (mw *MagickWand) ClipImage() error {
	C.MagickClipImage(mw.wand)
	return mw.GetLastError()
}

// Clips along the named paths from the 8BOM profile, if present. Later operations take
// effect inside the path. Id may be a number if preceded with #, to work on a numbered
// path, e.g. "#1" to use the first path.
// pathname: name of clipping path resource. If name is preceded by #, use clipping path numbered by name
// inside: if true, later operations take effect inside clipping path. Otherwise later operations take effect outside clipping path
func (mw *MagickWand) ClipImagePath(pathname string, inside bool) error {
	cspathname := C.CString(pathname)
	defer C.free(cspathname)
	csinside := 0
	if inside {
		csinside = 1
	}
	C.MagickClipImagePath(mw.wand, cspathname, C.MagickBooleanType(csinside))
	return mw.GetLastError()
}
