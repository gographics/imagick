// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <unistd.h>
#include <MagickWand/MagickWand.h>
*/
import "C"

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"syscall"
	"unsafe"
)

// Returns the current image from the magick wand
func (mw *MagickWand) GetImageFromMagickWand() *Image {
	return &Image{C.GetImageFromMagickWand(mw.mw)}
}

// Adaptively blurs the image by blurring less intensely near image edges and
// more intensely far from edges. We blur the image with a Gaussian operator of
// the given radius and standard deviation (sigma). For reasonable results,
// radius should be larger than sigma. Use a radius of 0 and
// AdaptiveBlurImage() selects a suitable radius for you.
//
// radius: the radius of the Gaussian, in pixels, not counting the center pixel
//
// sigma: the standard deviation of the Gaussian, in pixels
//
func (mw *MagickWand) AdaptiveBlurImage(radius, sigma float64) error {
	ok := C.MagickAdaptiveBlurImage(mw.mw, C.double(radius), C.double(sigma))
	return mw.getLastErrorIfFailed(ok)
}

// Adaptively resize image with data dependent triangulation
func (mw *MagickWand) AdaptiveResizeImage(cols, rows uint) error {
	ok := C.MagickAdaptiveResizeImage(mw.mw, C.size_t(cols), C.size_t(rows))
	return mw.getLastErrorIfFailed(ok)
}

// Adaptively sharpens the image by sharpening more intensely near image edges
// and less intensely far from edges. We sharpen the image with a Gaussian
// operator of the given radius and standard deviation (sigma). For reasonable
// results, radius should be larger than sigma. Use a radius of 0 and
// AdaptiveSharpenImage() selects a suitable radius for you.
//
// radius: the radius of the Gaussian, in pixels, not counting the center pixel
//
// sigma: the standard deviation of the Gaussian, in pixels.
//
func (mw *MagickWand) AdaptiveSharpenImage(radius, sigma float64) error {
	ok := C.MagickAdaptiveSharpenImage(mw.mw, C.double(radius), C.double(sigma))
	runtime.KeepAlive(mw)
	return mw.getLastErrorIfFailed(ok)
}

// Selects an individual threshold for each pixel based on the range of
// intensity values in its local neighborhood. This allows for thresholding
// of an image whose global intensity histogram doesn't contain distinctive
// peaks.
func (mw *MagickWand) AdaptiveThresholdImage(width, height uint, offset float64) error {
	ok := C.MagickAdaptiveThresholdImage(mw.mw, C.size_t(width), C.size_t(height), C.double(offset))
	return mw.getLastErrorIfFailed(ok)
}

// Adds a clone of the images from the second wand and inserts them into the
// first wand. Use SetLastIterator(), to append new images into an existing
// wand, current image will be set to last image so later adds with also be
// appened to end of wand. Use SetFirstIterator() to prepend new images into
// wand, any more images added will also be prepended before other images in
// the wand. However the order of a list of new images will not change.
// Otherwise the new images will be inserted just after the current image, and
// any later image will also be added after this current image but before the
// previously added images. Caution is advised when multiple image adds are
// inserted into the middle of the wand image list.
func (mw *MagickWand) AddImage(wand *MagickWand) error {
	ok := C.MagickAddImage(mw.mw, wand.mw)
	runtime.KeepAlive(wand)
	return mw.getLastErrorIfFailed(ok)
}

// Adds random noise to the image
func (mw *MagickWand) AddNoiseImage(noiseType NoiseType, offset float64) error {
	ok := C.MagickAddNoiseImage(mw.mw, C.NoiseType(noiseType), C.double(offset))
	return mw.getLastErrorIfFailed(ok)
}

// Transforms an image as dictaded by the affine matrix of the drawing wand
func (mw *MagickWand) AffineTransformImage(drawingWand *DrawingWand) error {
	ok := C.MagickAffineTransformImage(mw.mw, drawingWand.dw)
	runtime.KeepAlive(drawingWand)
	return mw.getLastErrorIfFailed(ok)
}

// Annotates an image with text
//
// x: ordinate to left of text
//
// y: ordinate to text baseline
//
// angle: rotate text relative to this angle
//
func (mw *MagickWand) AnnotateImage(drawingWand *DrawingWand, x, y, angle float64, text string) error {
	cstext := C.CString(text)
	defer C.free(unsafe.Pointer(cstext))
	ok := C.MagickAnnotateImage(mw.mw, drawingWand.dw, C.double(x), C.double(y), C.double(angle), cstext)
	runtime.KeepAlive(mw)
	runtime.KeepAlive(drawingWand)
	return mw.getLastErrorIfFailed(ok)
}

// Animates an image or image sequence
func (mw *MagickWand) AnimateImages(server string) error {
	csserver := C.CString(server)
	defer C.free(unsafe.Pointer(csserver))
	ok := C.MagickAnimateImages(mw.mw, csserver)
	return mw.getLastErrorIfFailed(ok)
}

// Append the images in a wand from the current image onwards, creating a new
// wand with the single image result. This is affected by the gravity and
// background setting of the first image. Typically you would call either
// ResetIterator() or SetFirstImage() before calling this function to ensure
// that all the images in the wand's image list will be appended together.
// By default, images are stacked left-to-right. Set topToBottom to true to
// stack them top-to-bottom.
func (mw *MagickWand) AppendImages(topToBottom bool) *MagickWand {
	ret := newMagickWand(C.MagickAppendImages(mw.mw, b2i(topToBottom)))
	runtime.KeepAlive(mw)
	return ret
}

// Extracts the 'mean' from the image and adjust the image to try make set
// it's gamma appropriatally
func (mw *MagickWand) AutoGammaImage() error {
	ok := C.MagickAutoGammaImage(mw.mw)
	return mw.getLastErrorIfFailed(ok)
}

// Adjust the levels of a particular image by scaling the minimum and maximum
// values to the full quantum range.
func (mw *MagickWand) AutoLevelImage() error {
	ok := C.MagickAutoLevelImage(mw.mw)
	return mw.getLastErrorIfFailed(ok)
}

// This is like ThresholdImage() but forces all pixels below the threshold
// into black while leaving all pixels above the threshold unchanged.
func (mw *MagickWand) BlackThresholdImage(threshold *PixelWand) error {
	ok := C.MagickBlackThresholdImage(mw.mw, threshold.pw)
	runtime.KeepAlive(threshold)
	return mw.getLastErrorIfFailed(ok)
}

// Mutes the colors of the image to simulate a scene at nighttime in the
// moonlight.
func (mw *MagickWand) BlueShiftImage(factor float64) error {
	ok := C.MagickBlueShiftImage(mw.mw, C.double(factor))
	return mw.getLastErrorIfFailed(ok)
}

// Blurs an image. We convolve the image with a gaussian operator of the
// given radius and standard deviation (sigma). For reasonable results, the
// radius should be larger than sigma. Use a radius of 0 and BlurImage()
// selects a suitable radius for you.
//
// radius: the radius of the, in pixels, not counting the center pixel.
//
// sigma: the standard deviation of the, in pixels
//
func (mw *MagickWand) BlurImage(radius, sigma float64) error {
	ok := C.MagickBlurImage(mw.mw, C.double(radius), C.double(sigma))
	return mw.getLastErrorIfFailed(ok)
}

// Surrounds the image with a border of the color defined by the bordercolor
// pixel wand.
func (mw *MagickWand) BorderImage(borderColor *PixelWand, width, height uint,
	compose CompositeOperator) error {

	ok := C.MagickBorderImage(mw.mw, borderColor.pw,
		C.size_t(width), C.size_t(height), C.CompositeOperator(compose))

	runtime.KeepAlive(borderColor)
	return mw.getLastErrorIfFailed(ok)
}

// Use this to change the brightness and/or contrast of an image. It converts
// the brightness and contrast.
//
// brighness: the brightness percent (-100 .. 100)
//
// contrast: the brightness percent (-100 .. 100)
//
func (mw *MagickWand) BrightnessContrastImage(brightness, contrast float64) error {
	ok := C.MagickBrightnessContrastImage(mw.mw, C.double(brightness), C.double(contrast))
	return mw.getLastErrorIfFailed(ok)
}

// Simulates a charcoal drawing
//
// radius: the radius of the Gaussian, in pixels, not counting the center pixel
//
// sigma: the standard deviation of the Gaussian, in pixels
//
func (mw *MagickWand) CharcoalImage(radius, sigma float64) error {
	ok := C.MagickCharcoalImage(mw.mw, C.double(radius), C.double(sigma))
	return mw.getLastErrorIfFailed(ok)
}

// Removes a region of an image and collapses the image to occupy the removed
// portion.
//
// width, height: the region width and height
//
// x, y: the region x and y offsets
//
func (mw *MagickWand) ChopImage(width, height uint, x, y int) error {
	ok := C.MagickChopImage(mw.mw, C.size_t(width), C.size_t(height), C.ssize_t(x), C.ssize_t(y))
	return mw.getLastErrorIfFailed(ok)
}

// Restricts the color range from 0 to the quantum depth
func (mw *MagickWand) ClampImage() error {
	ok := C.MagickClampImage(mw.mw)
	return mw.getLastErrorIfFailed(ok)
}

// Clips along the first path from the 8BIM profile, if present
func (mw *MagickWand) ClipImage() error {
	ok := C.MagickClipImage(mw.mw)
	return mw.getLastErrorIfFailed(ok)
}

// Clips along the named paths from the 8BOM profile, if present. Later
// operations take effect inside the path. Id may be a number if preceded with
// #, to work on a numbered path, e.g. "#1" to use the first path.
// pathname: name of clipping path resource. If name is preceded by #, use
// clipping path numbered by name.
//
// inside: if true, later operations take effect inside clipping path. Otherwise
// later operations take effect outside clipping path.
func (mw *MagickWand) ClipImagePath(pathname string, inside bool) error {
	cspathname := C.CString(pathname)
	defer C.free(unsafe.Pointer(cspathname))
	ok := C.MagickClipImagePath(mw.mw, cspathname, b2i(inside))
	return mw.getLastErrorIfFailed(ok)
}

// Replaces colors in the image from a color lookup table
func (mw *MagickWand) ClutImage(clut *MagickWand, interp PixelInterpolateMethod) error {
	ok := C.MagickClutImage(mw.mw, clut.mw, C.PixelInterpolateMethod(interp))
	runtime.KeepAlive(clut)
	return mw.getLastErrorIfFailed(ok)
}

// Composites a set of images while respecting any page offsets and disposal
// methods. GIF, MIFF, and MNG animation sequences typically start with an
// image background and each subsequent image varies in size and offset.
// CoalesceImages() returns a new sequence where each image in the sequence
// is the same size as the first and composited with the next image in the
// sequence.
func (mw *MagickWand) CoalesceImages() *MagickWand {
	return newMagickWand(C.MagickCoalesceImages(mw.mw))
}

// Accepts a lightweight Color Correction Collection (CCC) file which solely
// contains one or more color corrections and applies the color correction to
// the image. Here is a sample CCC file content:
//
//    <colorcorrectioncollection xmlns="urn:ASC:CDL:v1.2">
//      <colorcorrection id="cc03345">
//        <sopnode>
//          <slope> 0.9 1.2 0.5 </slope>
//          <offset> 0.4 -0.5 0.6 </offset>
//          <power> 1.0 0.8 1.5 </power>
//        </sopnode>
//        <satnode>
//          <saturation> 0.85 </saturation>
//        </satnode>
//      </colorcorrection>
//    </colorcorrectioncollection>
//
func (mw *MagickWand) ColorDecisionListImage(cccXML string) error {
	cscccXML := C.CString(cccXML)
	defer C.free(unsafe.Pointer(cscccXML))
	ok := C.MagickColorDecisionListImage(mw.mw, cscccXML)
	return mw.getLastErrorIfFailed(ok)
}

// Blends the fill color with each pixel in the image
func (mw *MagickWand) ColorizeImage(colorize, opacity *PixelWand) error {
	ok := C.MagickColorizeImage(mw.mw, colorize.pw, opacity.pw)
	runtime.KeepAlive(colorize)
	runtime.KeepAlive(opacity)
	return mw.getLastErrorIfFailed(ok)
}

// Apply color transformation to an image. The method permits saturation
// changes, hue rotation, luminance to alpha, and various other effects.
// Although variable-sized transformation matrices can be used, typically one
// uses a 5x5 matrix for an RGBA image and a 6x6 for CMYKA (or RGBA with
// offsets). The matrix is similar to those used by Adobe Flash except offsets
// are in column 6 rather than 5 (in support of CMYKA images) and offsets are
// normalized (divide Flash offset by 255).
func (mw *MagickWand) ColorMatrixImage(colorMatrix *KernelInfo) error {
	ok := C.MagickColorMatrixImage(mw.mw, colorMatrix.info)
	runtime.KeepAlive(colorMatrix)
	return mw.getLastErrorIfFailed(ok)
}

// Combines one or more images into a single image. The grayscale value of
// the pixels of each image in the sequence is assigned in order to the
// specified channels of the combined image. The typical ordering would be
// image 1 => Red, 2 => Green, 3 => Blue, etc.
func (mw *MagickWand) CombineImages(cs ColorspaceType) *MagickWand {
	ret := newMagickWand(C.MagickCombineImages(mw.mw, C.ColorspaceType(cs)))
	runtime.KeepAlive(mw)
	return ret
}

// Adds a comment to your image
func (mw *MagickWand) CommentImage(comment string) error {
	cscomment := C.CString(comment)
	defer C.free(unsafe.Pointer(cscomment))
	ok := C.MagickCommentImage(mw.mw, cscomment)
	return mw.getLastErrorIfFailed(ok)
}

// Compares each image with the next in a sequence and returns the maximum
// bounding region of any pixel differences it discovers.
func (mw *MagickWand) CompareImagesLayers(method LayerMethod) *MagickWand {
	return newMagickWand(C.MagickCompareImagesLayers(mw.mw, C.LayerMethod(method)))
}

// CompareImages() compares an image to a reconstructed image and returns the
// specified difference image. Returns the new MagickWand and the computed
// distortion between the images
func (mw *MagickWand) CompareImages(reference *MagickWand, metric MetricType) (wand *MagickWand, distortion float64) {
	cmw := C.MagickCompareImages(mw.mw, reference.mw, C.MetricType(metric), (*C.double)(&distortion))
	wand = newMagickWand(cmw)
	runtime.KeepAlive(mw)
	runtime.KeepAlive(reference)
	return
}

// Composite one image onto another at the specified offset.
// source: The magick wand holding source image.
// compose: This operator affects how the composite is applied to the image.
// The default is Over.
//
// x: the column offset of the composited image.
//
// y: the row offset of the composited image.
//
func (mw *MagickWand) CompositeImage(source *MagickWand,
	compose CompositeOperator, clipToSelf bool, x, y int) error {

	var clip C.MagickBooleanType = 0
	if clipToSelf {
		clip = 1
	}

	ok := C.MagickCompositeImage(mw.mw, source.mw,
		C.CompositeOperator(compose), clip,
		C.ssize_t(x), C.ssize_t(y))

	runtime.KeepAlive(source)
	return mw.getLastErrorIfFailed(ok)
}

// Composite one image onto another at the specified offset.
// source: The magick wand holding source image.
// compose: This operator affects how the composite is applied to the image.
// The default is Over.
// gravity : positioning gravity.
//
func (mw *MagickWand) CompositeImageGravity(source *MagickWand, compose CompositeOperator, gravity GravityType) error {
	ok := C.MagickCompositeImageGravity(mw.mw, source.mw, C.CompositeOperator(compose), C.GravityType(gravity))
	runtime.KeepAlive(source)
	return mw.getLastErrorIfFailed(ok)
}

// Composite the images in the source wand over the images in the destination
// wand in sequence, starting with the current image in both lists. Each layer
// from the two image lists are composted together until the end of one of the
// image lists is reached. The offset of each composition is also adjusted to
// match the virtual canvas offsets of each layer. As such the given offset is
// relative to the virtual canvas, and not the actual image.
// Composition uses given x and y offsets, as the 'origin' location of the
// source images virtual canvas (not the real image) allowing you to compose a
// list of 'layer images' into the destination images. This makes it well
// suitable for directly composing 'Clears Frame Animations' or 'Coaleased
// Animations' onto a static or other 'Coaleased Animation' destination image
// list. GIF disposal handling is not looked at. Special case: If one of the
// image sequences is the last image (just a single image remaining), that
// image is repeatally composed with all the images in the other image list.
// Either the source or destination lists may be the single image, for this
// situation. In the case of a single destination image (or last image given),
// that image will ve cloned to match the number of images remaining in the
// source image list. This is equivelent to the "-layer Composite" Shell API
// operator.
// source: the wand holding the source images
//
// compose, x, y: composition arguments
//
func (mw *MagickWand) CompositeLayers(source *MagickWand, compose CompositeOperator, x, y int) error {
	ok := C.MagickCompositeLayers(mw.mw, source.mw, C.CompositeOperator(compose), C.ssize_t(x), C.ssize_t(y))
	runtime.KeepAlive(source)
	return mw.getLastErrorIfFailed(ok)
}

// Enhances the intensity differences between the lighter and darker elements
// of the image. Set sharpen to a value other than 0 to increase the image
// contrast otherwise the contrast is reduced.
//
// sharpen: increase or decrease image contrast
//
func (mw *MagickWand) ContrastImage(sharpen bool) error {
	ok := C.MagickContrastImage(mw.mw, b2i(sharpen))
	return mw.getLastErrorIfFailed(ok)
}

// Enhances the contrast of a color image by adjusting the pixels color to
// span the entire range of colors available. You can also reduce the
// influence of a particular channel with a gamma value of 0.
func (mw *MagickWand) ContrastStretchImage(blackPoint, whitePoint float64) error {
	ok := C.MagickContrastStretchImage(mw.mw, C.double(blackPoint), C.double(whitePoint))
	return mw.getLastErrorIfFailed(ok)
}

// Applies a custom convolution kernel to the image.
//
// kernel: KernelInfo
//
func (mw *MagickWand) ConvolveImage(kernel *KernelInfo) error {
	ok := C.MagickConvolveImage(mw.mw, kernel.info)
	runtime.KeepAlive(kernel)
	return mw.getLastErrorIfFailed(ok)
}

// Extracts a region of the image
func (mw *MagickWand) CropImage(width, height uint, x, y int) error {
	ok := C.MagickCropImage(mw.mw, C.size_t(width), C.size_t(height), C.ssize_t(x), C.ssize_t(y))
	return mw.getLastErrorIfFailed(ok)
}

// Displaces an Image's colormap by a given number of positions. If you cycle
// the colormap a number of times you can produce a psychodelic effect.
func (mw *MagickWand) CycleColormapImage(displace int) error {
	ok := C.MagickCycleColormapImage(mw.mw, C.ssize_t(displace))
	return mw.getLastErrorIfFailed(ok)
}

// Adds an image to the wand comprised of the pixel data you supply. The pixel
// data must be in scanline order top-to-bottom.
//
// stype: Define the data type of the pixels. Float and double types are
// expected to be normalized [0..1] otherwise [0..QuantumRange].
//
// pmap: This string reflects the expected ordering of the pixel array. It can
// be any combination or order of R = red, G = green, B = blue, A = alpha
// (0 is transparent), O = opacity (0 is opaque), C = cyan, Y = yellow,
// M = magenta, K = black, I = intensity (for grayscale), P = pad.
//
// pixels: This array of values contain the pixel components as defined by the
// type.
//
func (mw *MagickWand) ConstituteImage(cols, rows uint, pmap string, stype StorageType, pixels interface{}) error {
	cspmap := C.CString(pmap)
	defer C.free(unsafe.Pointer(cspmap))
	ptr, calculatedStype, err := pixelInterfaceToPtr(pixels)
	if err != nil {
		return err
	}
	if stype == PIXEL_UNDEFINED {
		stype = calculatedStype
	}
	ok := C.MagickConstituteImage(mw.mw, C.size_t(cols), C.size_t(rows), cspmap, C.StorageType(stype), ptr)
	return mw.getLastErrorIfFailed(ok)
}

// Converts cipher pixels to plain pixels
func (mw *MagickWand) DecipherImage(passphrase string) error {
	cspassphrase := C.CString(passphrase)
	defer C.free(unsafe.Pointer(cspassphrase))
	ok := C.MagickDecipherImage(mw.mw, cspassphrase)
	return mw.getLastErrorIfFailed(ok)
}

// Compares each image with the next in a sequence and returns the maximum
// bouding region of any pixel differences it discovers.
func (mw *MagickWand) DeconstructImages() *MagickWand {
	ret := newMagickWand(C.MagickDeconstructImages(mw.mw))
	runtime.KeepAlive(mw)
	return ret
}

// Removes skew from the image. Skew is an artifact that occurs in scanned
// images because of the camera being misaligned, imperfections in the
// scanning or surface, or simply because the paper was not placed completely
// flat when scanned.
// threshold: separate background from foreground
func (mw *MagickWand) DeskewImage(threshold float64) error {
	ok := C.MagickDeskewImage(mw.mw, C.double(threshold))
	return mw.getLastErrorIfFailed(ok)
}

// Reduces the speckle noise in an image while perserving the edges of the
// original image.
func (mw *MagickWand) DespeckleImage() error {
	ok := C.MagickDespeckleImage(mw.mw)
	return mw.getLastErrorIfFailed(ok)
}

// Dereferences an image, deallocating memory associated with the image if the
// reference count becomes zero.
func (mw *MagickWand) DestroyImage(img *Image) *Image {
	ret := &Image{C.MagickDestroyImage(img.img)}
	runtime.KeepAlive(mw)
	runtime.KeepAlive(img)
	return ret
}

// Displays and image
func (mw *MagickWand) DisplayImage(server string) error {
	cstring := C.CString(server)
	defer C.free(unsafe.Pointer(cstring))
	ok := C.MagickDisplayImage(mw.mw, cstring)
	return mw.getLastErrorIfFailed(ok)
}

// Displays and image or image sequence
func (mw *MagickWand) DisplayImages(server string) error {
	cstring := C.CString(server)
	defer C.free(unsafe.Pointer(cstring))
	ok := C.MagickDisplayImages(mw.mw, cstring)
	return mw.getLastErrorIfFailed(ok)
}

// DistortImage() distorts an image using various distortion methods, by
// mapping color lookups of the source image to a new destination image usally
// of the same size as the source image, unless 'bestfit' is set to true. If
// 'bestfit' is enabled, and distortion allows it, the destination image is
// adjusted to ensure the whole source 'image' will just fit within the final
// destination image, which will be sized and offset accordingly. Also in many
// cases the virtual offset of the source image will be taken into account in
// the mapping.
//
// method: the method of image distortion. ArcDistortion always ignores the
// source image offset, and always 'bestfit' the destination image with the
// top left corner offset relative to the polar mapping center. Bilinear has
// no simple inverse mapping so it does not allow 'bestfit' style of image
// distortion. Affine, Perspective, and Bilinear, do least squares fitting of
// the distortion when more than the minimum number of control point pairs are
// provided. Perspective, and Bilinear, falls back to a Affine distortion when
// less that 4 control point pairs are provided. While Affine distortions let
// you use any number of control point pairs, that is Zero pairs is a no-Op
// (viewport only) distortion, one pair is a translation and two pairs of
// control points do a scale-rotate-translate, without any shearing.
//
// args: the arguments for this distortion method.
//
// bestfit: Attempt to resize destination to fit distorted source.
//
func (mw *MagickWand) DistortImage(method DistortImageMethod, args []float64, bestfit bool) error {
	ok := C.MagickDistortImage(mw.mw, C.DistortMethod(method), C.size_t(len(args)), (*C.double)(&args[0]), b2i(bestfit))
	return mw.getLastErrorIfFailed(ok)
}

// Renders the drawing wand on the current image
func (mw *MagickWand) DrawImage(drawingWand *DrawingWand) error {
	ok := C.MagickDrawImage(mw.mw, drawingWand.dw)
	runtime.KeepAlive(drawingWand)
	return mw.getLastErrorIfFailed(ok)
}

// Enhance edges within the image with a convolution filter of the given
// radius. Use a radius of 0 and Edge() selects a suitable radius for you.
//
// radius: the radius of the pixel neighborhood
//
func (mw *MagickWand) EdgeImage(radius float64) error {
	ok := C.MagickEdgeImage(mw.mw, C.double(radius))
	return mw.getLastErrorIfFailed(ok)
}

// Returns a grayscale image with a three-dimensional effect. We convolve the
// image with a Gaussian operator of the given radius and standard deviation
// (sigma). For reasonable results, radius should be larger than sigma. Use a
// radius of 0 and Emboss() selects a suitable radius for you.
//
// radius: the radius of the Gaussian, in pixels, not counting the center pixel
//
// sigma: the standard deviation of the Gaussian, in pixels
//
func (mw *MagickWand) EmbossImage(radius, sigma float64) error {
	ok := C.MagickEmbossImage(mw.mw, C.double(radius), C.double(sigma))
	return mw.getLastErrorIfFailed(ok)
}

// Converts plain pixels to cipher pixels
func (mw *MagickWand) EncipherImage(passphrase string) error {
	cspassphrase := C.CString(passphrase)
	defer C.free(unsafe.Pointer(cspassphrase))
	ok := C.MagickEncipherImage(mw.mw, cspassphrase)
	return mw.getLastErrorIfFailed(ok)
}

// Applies a digital filter that improves the quality of a noisy image
func (mw *MagickWand) EnhanceImage() error {
	ok := C.MagickEnhanceImage(mw.mw)
	return mw.getLastErrorIfFailed(ok)
}

// Equalizes the image histogram.
func (mw *MagickWand) EqualizeImage() error {
	ok := C.MagickEqualizeImage(mw.mw)
	return mw.getLastErrorIfFailed(ok)
}

// Applys an arithmetic, relational, or logical expression to an image.
// Use these operators to lighten or darken an image, to increase or
// decrease contrast in an image, or to produce the "negative" of an image.
func (mw *MagickWand) EvaluateImage(op EvaluateOperator, value float64) error {
	ok := C.MagickEvaluateImage(mw.mw, C.MagickEvaluateOperator(op), C.double(value))
	return mw.getLastErrorIfFailed(ok)
}

// Applys an arithmetic, relational, or logical expression to an image.
// Use these operators to lighten or darken an image, to increase or
// decrease contrast in an image, or to produce the "negative" of an image.
func (mw *MagickWand) EvaluateImages(op EvaluateOperator) error {
	C.MagickEvaluateImages(mw.mw, C.MagickEvaluateOperator(op))
	return mw.GetLastError()
}

// Extracts pixel data from an image and returns it to you.
//
// x, y, cols, rows: These values define the perimeter of a region of
// pixels you want to extract.
//
// map: This string reflects the expected ordering of the pixel array. It can
// be any combination or order of R = red, G = green, B = blue, A = alpha
// (0 is transparent), O = opacity (0 is opaque), C = cyan, Y = yellow,
// M = magenta, K = black, I = intensity (for grayscale), P = pad.
//
// stype: Define the data type of the pixels. Float and double types are
// expected to be normalized [0..1] otherwise [0..QuantumRange]. Choose from
// these types: CharPixel, DoublePixel, FloatPixel, IntegerPixel, LongPixel,
// QuantumPixel, or ShortPixel.
//
//
// StorageType defines the underlying slice type of the returned interface{}:
//     PIXEL_CHAR     => []byte
//     PIXEL_DOUBLE   => []float64
//     PIXEL_FLOAT    => []float32
//     PIXEL_SHORT    => []int16
//     PIXEL_LONG     => []int32
//     PIXEL_LONGLONG => []int64
//     PIXEL_QUANTUM  => []int64
//
// Example:
//
//     val, err := wand.ExportImagePixels(0, 0, 512, 512, "RGB", PIXEL_FLOAT)
//     if err != nil {
//         panic(err.Error())
//     }
//     floatPixels := val.([]float32)
//
func (mw *MagickWand) ExportImagePixels(x, y int, cols, rows uint,
	pmap string, stype StorageType) (interface{}, error) {
	if len(pmap) == 0 {
		return nil, errors.New("zero-length pmap not permitted")
	}

	if x < 0 || uint(x) > mw.GetImageWidth() ||
		y < 0 || uint(y) > mw.GetImageHeight() ||
		cols == 0 || rows == 0 {

		return nil, errors.New("Args x, y, cols, and rows produces an invalid region <= 0")
	}

	maplen := int(cols) * int(rows) * len(pmap)
	if maplen <= 0 {
		return nil, errors.New("Args cols and rows produces an invalid region <= 0")
	}

	var (
		pixel_iface interface{}
		ptr         unsafe.Pointer
	)

	switch stype {

	case PIXEL_CHAR:
		pixels := make([]byte, maplen)
		pixel_iface = reflect.ValueOf(pixels).Interface()
		ptr = unsafe.Pointer(&pixels[0])

	case PIXEL_DOUBLE:
		pixels := make([]float64, maplen)
		pixel_iface = reflect.ValueOf(pixels).Interface()
		ptr = unsafe.Pointer(&pixels[0])

	case PIXEL_FLOAT:
		pixels := make([]float32, maplen)
		pixel_iface = reflect.ValueOf(pixels).Interface()
		ptr = unsafe.Pointer(&pixels[0])

	case PIXEL_SHORT:
		pixels := make([]int16, maplen)
		pixel_iface = reflect.ValueOf(pixels).Interface()
		ptr = unsafe.Pointer(&pixels[0])

	case PIXEL_LONG:
		pixels := make([]int32, maplen)
		pixel_iface = reflect.ValueOf(pixels).Interface()
		ptr = unsafe.Pointer(&pixels[0])

	case PIXEL_LONGLONG, PIXEL_QUANTUM:
		pixels := make([]int64, maplen)
		pixel_iface = reflect.ValueOf(pixels).Interface()
		ptr = unsafe.Pointer(&pixels[0])

	default:
		return nil, errors.New("StorageType is not valid for this operation")

	}

	cspmap := C.CString(pmap)
	defer C.free(unsafe.Pointer(cspmap))

	ok := C.MagickExportImagePixels(mw.mw,
		C.ssize_t(x), C.ssize_t(y),
		C.size_t(cols), C.size_t(rows),
		cspmap,
		C.StorageType(stype),
		ptr)

	return pixel_iface, mw.getLastErrorIfFailed(ok)
}

// Extends the image as defined by the geometry, gravitt, and wand background
// color. Set the (x,y) offset of the geometry to move the original wand
// relative to the extended wand.
//
// width: the region width.
//
// height: the region height.
//
// x: the region x offset.
//
// y: the region y offset.
//
func (mw *MagickWand) ExtentImage(width, height uint, x, y int) error {
	ok := C.MagickExtentImage(mw.mw, C.size_t(width), C.size_t(height), C.ssize_t(x), C.ssize_t(y))
	return mw.getLastErrorIfFailed(ok)
}

// Creates a vertical mirror image by reflecting the pixels around the central
// x-axis.
func (mw *MagickWand) FlipImage() error {
	ok := C.MagickFlipImage(mw.mw)
	return mw.getLastErrorIfFailed(ok)
}

// Changes the color value of any pixel that matches target and is an immediate
// neighbor. If the method FillToBorderMethod is specified, the color value is
// changed for any neighbor pixel that does not match the bordercolor member
// of image.
//
// fill: the floodfill color pixel wand.
//
// fuzz: By default target must match a particular pixel color exactly.
//
// However, in many cases two colors may differ by a small amount. The fuzz
// member of image defines how much tolerance is acceptable to consider two
// colors as the same. For example, set fuzz to 10 and the color red at
// intensities of 100 and 102 respectively are now interpreted as the same
// color for the purposes of the floodfill.
//
// bordercolor: the border color pixel wand.
//
// x, y: the starting location of the operation.
//
// invert: paint any pixel that does not match the target color.
//
func (mw *MagickWand) FloodfillPaintImage(fill *PixelWand, fuzz float64, borderColor *PixelWand, x, y int, invert bool) error {
	ok := C.MagickFloodfillPaintImage(mw.mw, fill.pw, C.double(fuzz), borderColor.pw, C.ssize_t(x), C.ssize_t(y), b2i(invert))
	runtime.KeepAlive(fill)
	runtime.KeepAlive(borderColor)
	return mw.getLastErrorIfFailed(ok)
}

// Creates a horizontal mirror image by reflecting the pixels around the
// central y-axis.
func (mw *MagickWand) FlopImage() error {
	ok := C.MagickFlopImage(mw.mw)
	return mw.getLastErrorIfFailed(ok)
}

// Implements the discrete Fourier transform (DFT) of the image either as a
// magnitude/phase or real/imaginary image pair.
//
// magnitude: if true, return as magnitude/phase pair otherwise a
//
// real/imaginary image pair.
func (mw *MagickWand) ForwardFourierTransformImage(magnitude bool) error {
	ok := C.MagickForwardFourierTransformImage(mw.mw, b2i(magnitude))
	return mw.getLastErrorIfFailed(ok)
}

// Adds a simulated three-dimensional border around the image. The width and
// height specify the border width of the vertical and horizontal sides of the
// frame. The inner and outer bevels indicate the width of the inner and outer
// shadows of the frame.
//
// matteColor: the frame color pixel wand.
//
// width: the border width.
//
// height: the border height.
//
// innerBevel: the inner bevel width.
//
// outerBevel: the outer bevel width.
//
// compose: the composite operator.
//
func (mw *MagickWand) FrameImage(matteColor *PixelWand, width, height uint,
	innerBevel, outerBevel int, compose CompositeOperator) error {

	ok := C.MagickFrameImage(mw.mw, matteColor.pw,
		C.size_t(width), C.size_t(height),
		C.ssize_t(innerBevel), C.ssize_t(outerBevel),
		C.CompositeOperator(compose))

	runtime.KeepAlive(matteColor)
	return mw.getLastErrorIfFailed(ok)
}

// Applys an arithmetic, relational, or logical expression to an image. Use
// these operators to lighten or darken an image, to increase or decrease
// contrast in an image, or to produce the "negative" of an image.
func (mw *MagickWand) FunctionImage(function MagickFunction, args []float64) error {
	ok := C.MagickFunctionImage(mw.mw, C.MagickFunction(function), C.size_t(len(args)), (*C.double)(&args[0]))
	return mw.getLastErrorIfFailed(ok)
}

// Evaluate expression for each pixel in the image.
func (mw *MagickWand) FxImage(expression string) (fxmw *MagickWand, err error) {
	csexpression := C.CString(expression)
	defer C.free(unsafe.Pointer(csexpression))
	fxmw = newMagickWand(C.MagickFxImage(mw.mw, csexpression))
	err = mw.GetLastError()
	return
}

// Gamma-corrects an image. The same image viewed on different devices will
// have perceptual differences in the way the image's intensities are
// represented on the screen. Specify individual gamma levels for the red,
// green, and blue channels, or adjust all three with the gamma parameter.
// Values typically range from 0.8 to 2.3. You can also reduce the influence
// of a particular channel with a gamma value of 0.
func (mw *MagickWand) GammaImage(gamma float64) error {
	ok := C.MagickGammaImage(mw.mw, C.double(gamma))
	return mw.getLastErrorIfFailed(ok)
}

// Blurs an image. We convolve the image with a Gaussian operator of the given
// radius and standard deviation (sigma). For reasonable results, the radius
// should be larger than sigma. Use a radius of 0 and GaussianBlurImage()
// selects a suitable radius for you.
//
// radius: the radius of the Gaussian, in pixels, not counting the center pixel.
//
// sigma: the standard deviation of the Gaussian, in pixels.
//
func (mw *MagickWand) GaussianBlurImage(radius, sigma float64) error {
	ok := C.MagickGaussianBlurImage(mw.mw, C.double(radius), C.double(sigma))
	return mw.getLastErrorIfFailed(ok)
}

// Gets the image at the current image index.
func (mw *MagickWand) GetImage() *MagickWand {
	ret := newMagickWand(C.MagickGetImage(mw.mw))
	runtime.KeepAlive(mw)
	return ret
}

// GetImageAlphaChannel returns MagickFalse if the image alpha channel is not
// activated. That is, the image is RGB rather than RGBA or CMYK rather than CMYKA.
func (mw *MagickWand) GetImageAlphaChannel() bool {
	ret := 1 == C.MagickGetImageAlphaChannel(mw.mw)
	runtime.KeepAlive(mw)
	return ret
}

// Returns the image background color.
func (mw *MagickWand) GetImageBackgroundColor() (bgColor *PixelWand, err error) {
	cbgcolor := NewPixelWand()
	ok := C.MagickGetImageBackgroundColor(mw.mw, cbgcolor.pw)
	return cbgcolor, mw.getLastErrorIfFailed(ok)
}

// Implements direct to memory image formats. It returns the image as a blob
// (a formatted "file" in memory) and its length, starting from the current
// position in the image sequence. Use SetImageFormat() to set the format to
// write to the blob (GIF, JPEG, PNG, etc.). Utilize ResetIterator() to ensure
// the write is from the beginning of the image sequence.
func (mw *MagickWand) GetImageBlob() []byte {
	clen := C.size_t(0)
	csblob := C.MagickGetImageBlob(mw.mw, &clen)
	defer relinquishMemory(unsafe.Pointer(csblob))
	ret := C.GoBytes(unsafe.Pointer(csblob), C.int(clen))
	runtime.KeepAlive(mw)
	return ret
}

// Implements direct to memory image formats. It returns the image sequence
// as a blob and its length. The format of the image determines the format of
// the returned blob (GIF, JPEG, PNG, etc.). To return a different image
// format, use SetImageFormat(). Note, some image formats do not permit
// multiple images to the same image stream (e.g. JPEG). in this instance,
// just the first image of the sequence is returned as a blob.
func (mw *MagickWand) GetImagesBlob() []byte {
	clen := C.size_t(0)
	csblob := C.MagickGetImagesBlob(mw.mw, &clen)
	defer relinquishMemory(unsafe.Pointer(csblob))
	runtime.KeepAlive(mw)
	return C.GoBytes(unsafe.Pointer(csblob), C.int(clen))
}

// Returns the chromaticy blue primary point for the image.
//
// x: the chromaticity blue primary x-point.
//
// y: the chromaticity blue primary y-point.
//
// z: the chromaticity blue primary z-point.
//
func (mw *MagickWand) GetImageBluePrimary() (x, y, z float64, err error) {
	ok := C.MagickGetImageBluePrimary(mw.mw,
		(*C.double)(&x), (*C.double)(&y), (*C.double)(&z))
	err = mw.getLastErrorIfFailed(ok)
	return
}

// Returns the image border color.
func (mw *MagickWand) GetImageBorderColor() (borderColor *PixelWand, err error) {
	cbc := NewPixelWand()
	ok := C.MagickGetImageBorderColor(mw.mw, cbc.pw)
	return cbc, mw.getLastErrorIfFailed(ok)
}

// Returns the color of the specified colormap index.
func (mw *MagickWand) GetImageColormapColor(index uint) (color *PixelWand, err error) {
	cpw := NewPixelWand()
	ok := C.MagickGetImageColormapColor(mw.mw, C.size_t(index), cpw.pw)
	return cpw, mw.getLastErrorIfFailed(ok)
}

// Gets the number of unique colors in the image.
func (mw *MagickWand) GetImageColors() uint {
	ret := uint(C.MagickGetImageColors(mw.mw))
	runtime.KeepAlive(mw)
	return ret
}

// Gets the image colorspace.
func (mw *MagickWand) GetImageColorspace() ColorspaceType {
	ret := ColorspaceType(C.MagickGetImageColorspace(mw.mw))
	runtime.KeepAlive(mw)
	return ret
}

// Returns the composite operator associated with the image.
func (mw *MagickWand) GetImageCompose() CompositeOperator {
	ret := CompositeOperator(C.MagickGetImageCompose(mw.mw))
	runtime.KeepAlive(mw)
	return ret
}

// Gets the image compression.
func (mw *MagickWand) GetImageCompression() CompressionType {
	ret := CompressionType(C.MagickGetImageCompression(mw.mw))
	runtime.KeepAlive(mw)
	return ret
}

// Gets the image compression quality.
func (mw *MagickWand) GetImageCompressionQuality() uint {
	ret := uint(C.MagickGetImageCompressionQuality(mw.mw))
	runtime.KeepAlive(mw)
	return ret
}

// Gets the image delay.
func (mw *MagickWand) GetImageDelay() uint {
	ret := uint(C.MagickGetImageDelay(mw.mw))
	runtime.KeepAlive(mw)
	return ret
}

// Gets the image depth.
func (mw *MagickWand) GetImageDepth() uint {
	ret := uint(C.MagickGetImageDepth(mw.mw))
	runtime.KeepAlive(mw)
	return ret
}

// Compares an image to a reconstructed image and returns the specified
// distortion metric.
func (mw *MagickWand) GetImageDistortion(reference *MagickWand, metric MetricType) (distortion float64, err error) {
	ok := C.MagickGetImageDistortion(mw.mw, reference.mw, C.MetricType(metric), (*C.double)(&distortion))
	runtime.KeepAlive(reference)
	err = mw.getLastErrorIfFailed(ok)
	return
}

// Gets the image disposal method.
func (mw *MagickWand) GetImageDispose() DisposeType {
	ret := DisposeType(C.MagickGetImageDispose(mw.mw))
	runtime.KeepAlive(mw)
	return ret
}

// Gets the image endian.
func (mw *MagickWand) GetImageEndian() EndianType {
	ret := EndianType(C.MagickGetImageEndian(mw.mw))
	runtime.KeepAlive(mw)
	return ret
}

// Returns the filename of a particular image in a sequence.
func (mw *MagickWand) GetImageFilename() string {
	p := C.MagickGetImageFilename(mw.mw)
	defer relinquishMemory(unsafe.Pointer(p))
	return C.GoString(p)
}

// Returns the format of a particular image in a sequence.
func (mw *MagickWand) GetImageFormat() string {
	p := C.MagickGetImageFormat(mw.mw)
	runtime.KeepAlive(mw)
	defer relinquishMemory(unsafe.Pointer(p))
	return C.GoString(p)
}

// Gets the image fuzz.
func (mw *MagickWand) GetImageFuzz() float64 {
	ret := float64(C.MagickGetImageFuzz(mw.mw))
	runtime.KeepAlive(mw)
	return ret
}

// Gets the image gamma.
func (mw *MagickWand) GetImageGamma() float64 {
	ret := float64(C.MagickGetImageGamma(mw.mw))
	runtime.KeepAlive(mw)
	return ret
}

// Gets the image gravity.
func (mw *MagickWand) GetImageGravity() GravityType {
	ret := GravityType(C.MagickGetImageGravity(mw.mw))
	runtime.KeepAlive(mw)
	return ret
}

// Returns the chromaticy green primary point.
//
// x: the chromaticity green primary x-point.
//
// y: the chromaticity green primary y-point.
//
// z: the chromaticity green primary z-point.
//
func (mw *MagickWand) GetImageGreenPrimary() (x, y, z float64, err error) {
	ok := C.MagickGetImageGreenPrimary(mw.mw,
		(*C.double)(&x), (*C.double)(&y), (*C.double)(&z))
	err = mw.getLastErrorIfFailed(ok)
	return
}

// Returns the image height.
func (mw *MagickWand) GetImageHeight() uint {
	ret := uint(C.MagickGetImageHeight(mw.mw))
	runtime.KeepAlive(mw)
	return ret
}

// Returns the image histogram as an array of PixelWand wands.
//
// numberColors: the number of unique colors in the image and the number of
// pixel wands returned.
func (mw *MagickWand) GetImageHistogram() (numberColors uint, pws []*PixelWand) {
	cnc := C.size_t(0)
	p := C.MagickGetImageHistogram(mw.mw, &cnc)
	defer relinquishMemory(unsafe.Pointer(p))
	q := uintptr(unsafe.Pointer(p))
	numberColors = uint(cnc)
	for i := 0; i < int(numberColors); i++ {
		p = (**C.PixelWand)(unsafe.Pointer(q))
		if *p == nil {
			break
		}
		pws = append(pws, newPixelWand(*p))
		q += unsafe.Sizeof(q)
	}
	runtime.KeepAlive(mw)
	return
}

// Gets the image interlace scheme.
func (mw *MagickWand) GetImageInterlaceScheme() InterlaceType {
	ret := InterlaceType(C.MagickGetImageInterlaceScheme(mw.mw))
	runtime.KeepAlive(mw)
	return ret
}

// Returns the interpolation method for the sepcified image.
func (mw *MagickWand) GetImageInterpolateMethod() PixelInterpolateMethod {
	ret := PixelInterpolateMethod(C.MagickGetImageInterpolateMethod(mw.mw))
	runtime.KeepAlive(mw)
	return ret
}

// Gets the image iterations.
func (mw *MagickWand) GetImageIterations() uint {
	ret := uint(C.MagickGetImageIterations(mw.mw))
	runtime.KeepAlive(mw)
	return ret
}

// Returns the image length in bytes.
func (mw *MagickWand) GetImageLength() (length uint, err error) {
	cl := C.MagickSizeType(0)
	ok := C.MagickGetImageLength(mw.mw, &cl)
	return uint(cl), mw.getLastErrorIfFailed(ok)
}

// Returns the image orientation.
func (mw *MagickWand) GetImageOrientation() OrientationType {
	ret := OrientationType(C.MagickGetImageOrientation(mw.mw))
	runtime.KeepAlive(mw)
	return ret
}

// Returns the page geometry associated with the image.
//
// w, h: the page width and height
//
// x, h: the page x-offset and y-offset.
//
func (mw *MagickWand) GetImagePage() (w, h uint, x, y int, err error) {
	var cw, ch C.size_t
	var cx, cy C.ssize_t
	ok := C.MagickGetImagePage(mw.mw, &cw, &ch, &cx, &cy)
	return uint(cw), uint(ch), int(cx), int(cy), mw.getLastErrorIfFailed(ok)
}

// Returns the color of the specified pixel.
func (mw *MagickWand) GetImagePixelColor(x, y int) (color *PixelWand, err error) {
	pw := NewPixelWand()
	ok := C.MagickGetImagePixelColor(mw.mw, C.ssize_t(x), C.ssize_t(y), pw.pw)
	return pw, mw.getLastErrorIfFailed(ok)
}

// Returns the chromaticy red primary point.
//
// x, y, z: the chromaticity red primary x/y/z-point.
//
func (mw *MagickWand) GetImageRedPrimary() (x, y, z float64, err error) {
	var cdx, cdy, cdz C.double
	ok := C.MagickGetImageRedPrimary(mw.mw, &cdx, &cdy, &cdz)
	return float64(cdx), float64(cdy), float64(cdz), mw.getLastErrorIfFailed(ok)
}

// Extracts a region of the image and returns it as a a new wand.
func (mw *MagickWand) GetImageRegion(width uint, height uint, x int, y int) *MagickWand {
	ret := newMagickWand(C.MagickGetImageRegion(mw.mw, C.size_t(width), C.size_t(height), C.ssize_t(x), C.ssize_t(y)))
	runtime.KeepAlive(mw)
	return ret
}

// Gets the image rendering intent.
func (mw *MagickWand) GetImageRenderingIntent() RenderingIntent {
	ret := RenderingIntent(C.MagickGetImageRenderingIntent(mw.mw))
	runtime.KeepAlive(mw)
	return ret
}

// Gets the image X and Y resolution.
func (mw *MagickWand) GetImageResolution() (x, y float64, err error) {
	var dx, dy C.double
	ok := C.MagickGetImageResolution(mw.mw, &dx, &dy)
	return float64(dx), float64(dy), mw.getLastErrorIfFailed(ok)
}

// Gets the image scene.
func (mw *MagickWand) GetImageScene() uint {
	ret := uint(C.MagickGetImageScene(mw.mw))
	runtime.KeepAlive(mw)
	return ret
}

// Generates an SHA-256 message digest for the image pixel stream.
func (mw *MagickWand) GetImageSignature() string {
	p := C.MagickGetImageSignature(mw.mw)
	defer relinquishMemory(unsafe.Pointer(p))
	runtime.KeepAlive(mw)
	return C.GoString(p)
}

// Gets the image ticks-per-second.
func (mw *MagickWand) GetImageTicksPerSecond() uint {
	ret := uint(C.MagickGetImageTicksPerSecond(mw.mw))
	runtime.KeepAlive(mw)
	return ret
}

// Gets the potential image type
// To ensure the image type matches its potential, use SetImageType():
// wand.SetImageType(wand.GetImageType())
func (mw *MagickWand) GetImageType() ImageType {
	ret := ImageType(C.MagickGetImageType(mw.mw))
	runtime.KeepAlive(mw)
	return ret
}

// Gets the image units of resolution.
func (mw *MagickWand) GetImageUnits() ResolutionType {
	ret := ResolutionType(C.MagickGetImageUnits(mw.mw))
	runtime.KeepAlive(mw)
	return ret
}

// Returns the virtual pixel method for the specified image.
func (mw *MagickWand) GetImageVirtualPixelMethod() VirtualPixelMethod {
	ret := VirtualPixelMethod(C.MagickGetImageVirtualPixelMethod(mw.mw))
	runtime.KeepAlive(mw)
	return ret
}

// Returns the chromaticy white point.
//
// x, y, z: the chromaticity white x/y-point.
//
func (mw *MagickWand) GetImageWhitePoint() (x, y, z float64, err error) {
	ok := C.MagickGetImageWhitePoint(mw.mw,
		(*C.double)(&x), (*C.double)(&y), (*C.double)(&z))
	err = mw.getLastErrorIfFailed(ok)
	return
}

// Returns the image width.
func (mw *MagickWand) GetImageWidth() uint {
	ret := uint(C.MagickGetImageWidth(mw.mw))
	runtime.KeepAlive(mw)
	return ret
}

// Returns the number of images associated with a magick wand.
func (mw *MagickWand) GetNumberImages() uint {
	ret := uint(C.MagickGetNumberImages(mw.mw))
	runtime.KeepAlive(mw)
	return ret
}

// Gets the image total ink density.
func (mw *MagickWand) GetImageTotalInkDensity() float64 {
	ret := float64(C.MagickGetImageTotalInkDensity(mw.mw))
	runtime.KeepAlive(mw)
	return ret
}

// Replaces colors in the image from a Hald color lookup table. A Hald color
// lookup table is a 3-dimensional color cube mapped to 2 dimensions. Create
// it with the HALD coder. You can apply any color transformation to the Hald
// image and then use this method to apply the transform to the image.
func (mw *MagickWand) HaldClutImage(hald *MagickWand) error {
	ok := C.MagickHaldClutImage(mw.mw, hald.mw)
	return mw.getLastErrorIfFailed(ok)
}

// Returns true if the wand has more images when traversing the list in the
// forward direction
func (mw *MagickWand) HasNextImage() bool {
	ret := 1 == C.MagickHasNextImage(mw.mw)
	runtime.KeepAlive(mw)
	return ret
}

// Returns true if the wand has more images when traversing the list in the
// reverse direction
func (mw *MagickWand) HasPreviousImage() bool {
	ret := 1 == C.MagickHasPreviousImage(mw.mw)
	runtime.KeepAlive(mw)
	return ret
}

// Identifies an image by printing its attributes to the file. Attributes
// include the image width, height, size, and others.
func (mw *MagickWand) IdentifyImage() string {
	p := C.MagickIdentifyImage(mw.mw)
	defer relinquishMemory(unsafe.Pointer(p))
	runtime.KeepAlive(mw)
	return C.GoString(p)
}

// Creates a new image that is a copy of an existing one with the image pixels
// "implode" by the specified percentage.
//
// radius: Define the extent of the implosion.
//
// method: the pixel interpolation method.
//
func (mw *MagickWand) ImplodeImage(radius float64, method PixelInterpolateMethod) error {
	ok := C.MagickImplodeImage(mw.mw, C.double(radius), C.PixelInterpolateMethod(method))
	return mw.getLastErrorIfFailed(ok)
}

// Identifies the type of pixels and returns the storage type and an
// unsafe pointer to the data.
func pixelInterfaceToPtr(pixels interface{}) (unsafe.Pointer, StorageType, error) {
	var ptr unsafe.Pointer
	var stype StorageType

	switch t := pixels.(type) {

	case []byte:
		v := &t[0]
		ptr = unsafe.Pointer(v)
		stype = PIXEL_CHAR

	case []float64:
		v := &t[0]
		ptr = unsafe.Pointer(v)
		stype = PIXEL_DOUBLE

	case []float32:
		v := &t[0]
		ptr = unsafe.Pointer(v)
		stype = PIXEL_FLOAT

	case []int16:
		v := &t[0]
		ptr = unsafe.Pointer(v)
		stype = PIXEL_SHORT

	case []int32:
		v := &t[0]
		ptr = unsafe.Pointer(v)
		stype = PIXEL_LONG

	case []int64:
		v := &t[0]
		ptr = unsafe.Pointer(v)
		stype = PIXEL_LONGLONG

	default:
		return ptr, PIXEL_UNDEFINED, fmt.Errorf("Type %T is not valid for this operation", t)
	}

	return ptr, stype, nil
}

// Accepts pixel data and stores it in the image at the location you specify.
// The pixel data can be either byte, int16, int32, int64, float32, or float64
// in the order specified by map. Suppose your want to upload the first
// scanline of a 640x480 image from character data in red-green-blue order:
//
//     wand.ImportImagePixels(0, 0, 640, 1, "RGB", PIXEL_CHAR, pixels)
//
// x, y, cols, rows: These values define the perimeter of a region of pixels
// you want to define.
//
// pmap: This string reflects the expected ordering of the pixel array. It can
// be any combination or order of R = red, G = green, B = blue, A = alpha (0
// is transparent), O = opacity (0 is opaque), C = cyan, Y = yellow,
// M = magenta, K = black, I = intensity (for grayscale), P = pad.
//
// stype: Define the data type of the pixels. Float and double types are
// expected to be normalized [0..1] otherwise [0..QuantumRange]. Choose from
// these types:
//
//     PIXEL_CHAR
//     PIXEL_SHORT
//     PIXEL_INTEGER
//     PIXEL_LONG
//     PIXEL_FLOAT
//     PIXEL_DOUBLE
//
// pixels: This slice of values contain the pixel components as defined by map
// and type. You must pre-allocate this slice where the expected length varies
// depending on the values of width, height, map, and type.
//
func (mw *MagickWand) ImportImagePixels(x, y int, cols, rows uint, pmap string,
	stype StorageType, pixels interface{}) error {

	cspmap := C.CString(pmap)
	defer C.free(unsafe.Pointer(cspmap))

	var ptr unsafe.Pointer

	ptr, calculatedStype, err := pixelInterfaceToPtr(pixels)
	if err != nil {
		return err
	}
	if stype == PIXEL_UNDEFINED {
		stype = calculatedStype
	}

	ok := C.MagickImportImagePixels(mw.mw, C.ssize_t(x), C.ssize_t(y), C.size_t(cols),
		C.size_t(rows), cspmap, C.StorageType(stype), ptr)

	return mw.getLastErrorIfFailed(ok)
}

// Implements the inverse discrete Fourier transform (DFT) of the image either
// as a magnitude/phase or real/imaginary image pair.
//
// magnitudeWand: the magnitude or real wand.
//
// phaseWand: the phase or imaginary wand.
//
// magnitude: if true, return as magnitude/phase pair otherwise a
// real/imaginary image pair.
//
func (mw *MagickWand) InverseFourierTransformImage(phaseWand *MagickWand, magnitude bool) error {
	ok := C.MagickInverseFourierTransformImage(mw.mw, phaseWand.mw, b2i(magnitude))
	return mw.getLastErrorIfFailed(ok)
}

// Adds a label to your image.
func (mw *MagickWand) LabelImage(label string) error {
	cslabel := C.CString(label)
	defer C.free(unsafe.Pointer(cslabel))
	ok := C.MagickLabelImage(mw.mw, cslabel)
	return mw.getLastErrorIfFailed(ok)
}

// Adjusts the levels of an image by scaling the colors falling between
// specified white and black points to the full available quantum range. The
// parameters provided represent the black, mid, and white points. The black
// point specifies the darkest color in the image. Colors darker than the
// black point are set to zero. Mid point specifies a gamma correction to
// apply to the image. White point specifies the lightest color in the image.
// Colors brighter than the white point are set to the maximum quantum value.
func (mw *MagickWand) LevelImage(blackPoint, gamma, whitePoint float64) error {
	ok := C.MagickLevelImage(mw.mw, C.double(blackPoint), C.double(gamma), C.double(whitePoint))
	return mw.getLastErrorIfFailed(ok)
}

// Stretches with saturation the image intensity. You can also reduce the
// influence of a particular channel with a gamma value of 0.
func (mw *MagickWand) LinearStretchImage(blackPoint, whitePoint float64) error {
	ok := C.MagickLinearStretchImage(mw.mw, C.double(blackPoint), C.double(whitePoint))
	return mw.getLastErrorIfFailed(ok)
}

// Rescales image with seam carving.
//
// cols, rows: the number of cols and rows in the scaled image.
//
// deltaX: maximum seam transversal step (0 means straight seams).
//
// rigidity: introduce a bias for non-straight seams (typically 0).
//
func (mw *MagickWand) LiquidRescaleImage(cols, rows uint, deltaX, rigidity float64) error {
	ok := C.MagickLiquidRescaleImage(mw.mw, C.size_t(cols), C.size_t(rows), C.double(deltaX), C.double(rigidity))
	return mw.getLastErrorIfFailed(ok)
}

// This is a convenience method that scales an image proportionally to twice
// its original size.
func (mw *MagickWand) MagnifyImage() error {
	ok := C.MagickMagnifyImage(mw.mw)
	return mw.getLastErrorIfFailed(ok)
}

// Composes all the image layers from the current given image onward to
// produce a single image of the merged layers. The inital canvas's size
// depends on the given ImageLayerMethod, and is initialized using the first
// images background color. The images are then compositied onto that image in
// sequence using the given composition that has been assigned to each
// individual image.
//
// method: the method of selecting the size of the initial canvas. MergeLayer:
// Merge all layers onto a canvas just large enough to hold all the actual
// images. The virtual canvas of the first image is preserved but otherwise
// ignored. FlattenLayer: Use the virtual canvas size of first image. Images
// which fall outside this canvas is clipped. This can be used to 'fill out'
// a given virtual canvas. MosaicLayer: Start with the virtual canvas of the
// first image, enlarging left and right edges to contain all images. Images
// with negative offsets will be clipped.
func (mw *MagickWand) MergeImageLayers(method LayerMethod) *MagickWand {
	ret := newMagickWand(C.MagickMergeImageLayers(mw.mw, C.LayerMethod(method)))
	runtime.KeepAlive(mw)
	return ret
}

// This is a convenience method that scales an image proportionally to
// one-half its original size
func (mw *MagickWand) MinifyImage() error {
	ok := C.MagickMinifyImage(mw.mw)
	return mw.getLastErrorIfFailed(ok)
}

// Lets you control the brightness, saturation, and hue of an image. Hue is
// the percentage of absolute rotation from the current position. For example
// 50 results in a counter-clockwise rotation of 90 degrees, 150 results in a
// clockwise rotation of 90 degrees, with 0 and 200 both resulting in a
// rotation of 180 degrees. To increase the color brightness by 20 and
// decrease the color saturation by 10 and leave the hue unchanged, use: 120,
// 90, 100.
//
// brightness: the percent change in brighness.
//
// saturation: the percent change in saturation.
//
// hue: the percent change in hue.
//
func (mw *MagickWand) ModulateImage(brightness, saturation, hue float64) error {
	ok := C.MagickModulateImage(mw.mw, C.double(brightness), C.double(saturation), C.double(hue))
	return mw.getLastErrorIfFailed(ok)
}

// Creates a composite image by combining several separate images. The images
// are tiled on the composite image with the name of the image optionally
// appearing just below the individual tile.
//
// dw: the drawing wand. The font name, size, and color are obtained from this
// wand.
//
// tileGeo: the number of tiles per row and page (e.g. 6x4+0+0).
//
// thumbGeo: Preferred image size and border size of each thumbnail (e.g.
// 120x120+4+3>).
//
// mode: Thumbnail framing mode: Frame, Unframe, or Concatenate.
//
// frame: Surround the image with an ornamental border (e.g. 15x15+3+3). The
// frame color is that of the thumbnail's matte color.
//
func (mw *MagickWand) MontageImage(dw *DrawingWand, tileGeo string, thumbGeo string, mode MontageMode, frame string) *MagickWand {
	cstile := C.CString(tileGeo)
	defer C.free(unsafe.Pointer(cstile))
	csthumb := C.CString(thumbGeo)
	defer C.free(unsafe.Pointer(csthumb))
	csframe := C.CString(frame)
	defer C.free(unsafe.Pointer(csframe))

	ret := newMagickWand(C.MagickMontageImage(mw.mw, dw.dw, cstile, csthumb, C.MontageMode(mode), csframe))
	runtime.KeepAlive(mw)
	runtime.KeepAlive(dw)
	return ret
}

// Method morphs a set of images. Both the image pixels and size are linearly
// interpolated to give the appearance of a meta-morphosis from one image to
// the next.
//
// numFrames: the number of in-between images to generate.
func (mw *MagickWand) MorphImages(numFrames uint) *MagickWand {
	ret := newMagickWand(C.MagickMorphImages(mw.mw, C.size_t(numFrames)))
	runtime.KeepAlive(mw)
	return ret
}

// Applies a user supplied kernel to the image according to the given mophology
// method.
//
// channel: the image channel(s).
//
// method: the morphology method to be applied.
//
// iterations: apply the operation this many times (or no change). A value of
// -1 means loop until no change found. How this is applied may depend on the
// morphology method. Typically this is a value of 1.
//
// kernel: An array of doubles representing the morphology kernel.
func (mw *MagickWand) MorphologyImage(method MorphologyMethod, iterations int, kernel *KernelInfo) error {
	ok := C.MagickMorphologyImage(mw.mw, C.MorphologyMethod(method), C.ssize_t(iterations), kernel.info)
	return mw.getLastErrorIfFailed(ok)
}

// Simulates motion blur. We convolve the image with a Gaussian operator of
// the given radius and standard deviation (sigma). For reasonable results,
// radius should be larger than sigma. Use a radius of 0 and MotionBlurImage()
// selects a suitable radius for you. Angle gives the angle of the blurring
// motion.
//
// radius: the radius of the Gaussian, in pixels, not counting the center pixel.
//
// sigma: the standard deviation of the Gaussian, in pixels.
//
// angle: apply the effect along this angle.
//
func (mw *MagickWand) MotionBlurImage(radius, sigma, angle float64) error {
	ok := C.MagickMotionBlurImage(mw.mw, C.double(radius), C.double(sigma), C.double(angle))
	return mw.getLastErrorIfFailed(ok)
}

// Negates the colors in the reference image. The Grayscale option means that
// only grayscale values within the image are negated. You can also reduce the
// influence of a particular channel with a gamma value of 0.
//
// gray: If true, only negate grayscale pixels within the image.
//
func (mw *MagickWand) NegateImage(gray bool) error {
	ok := C.MagickNegateImage(mw.mw, b2i(gray))
	return mw.getLastErrorIfFailed(ok)
}

// Adds a blank image canvas of the specified size and background color to the
// wand.
func (mw *MagickWand) NewImage(cols uint, rows uint, background *PixelWand) error {
	ok := C.MagickNewImage(mw.mw, C.size_t(cols), C.size_t(rows), background.pw)
	return mw.getLastErrorIfFailed(ok)
}

// Sets the next image in the wand as the current image. It is typically used
// after ResetIterator(), after which its first use will set the first image
// as the current image (unless the wand is empty). It will return false when
// no more images are left to be returned which happens when the wand is empty,
// or the current image is the last image. When the above condition (end of
// image list) is reached, the iterator is automaticall set so that you can
// start using PreviousImage() to again/ iterate over the images in the
// reverse direction, starting with the last image (again). You can jump to
// this condition immeditally using SetLastIterator().
func (mw *MagickWand) NextImage() bool {
	ret := 1 == C.MagickNextImage(mw.mw)
	runtime.KeepAlive(mw)
	return ret
}

// Enhances the contrast of a color image by adjusting the pixels color to
// span the entire range of colors available. You can also reduce the
// influence of a particular channel with a gamma value of 0.
func (mw *MagickWand) NormalizeImage() error {
	ok := C.MagickNormalizeImage(mw.mw)
	return mw.getLastErrorIfFailed(ok)
}

// Applies a special effect filter that simulates an oil painting. Each pixel
// is replaced by the most frequent color occurring in a circular region
// defined by radius.
//
// radius: the radius of the circular neighborhood.
//
// sigma: the standard deviation of the Gaussian, in pixels.
//
func (mw *MagickWand) OilPaintImage(radius, sigma float64) error {
	ok := C.MagickOilPaintImage(mw.mw, C.double(radius), C.double(sigma))
	return mw.getLastErrorIfFailed(ok)
}

// Changes any pixel that matches color with the color defined by fill.
//
// target: Change this target color to the fill color within the image.
//
// fill: the fill pixel wand.
//
// fuzz: By default target must match a particular pixel color exactly.
// However, in many cases two colors may differ by a small amount. The fuzz
// member of image defines how much tolerance is acceptable to consider two
// colors as the same. For example, set fuzz to 10 and the color red at
// intensities of 100 and 102 respectively are now interpreted as the same
// color for the purposes of the floodfill.
//
// invert: paint any pixel that does not match the target color.
//
func (mw *MagickWand) OpaquePaintImage(target, fill *PixelWand, fuzz float64, invert bool) error {
	ok := C.MagickOpaquePaintImage(mw.mw, target.pw, fill.pw, C.double(fuzz), b2i(invert))
	return mw.getLastErrorIfFailed(ok)
}

// Compares each image the GIF disposed forms of the previous image in the
// sequence. From this it attempts to select the smallest cropped image to
// replace each frame, while preserving the results of the animation.
func (mw *MagickWand) OptimizeImageLayers() *MagickWand {
	ret := newMagickWand(C.MagickOptimizeImageLayers(mw.mw))
	runtime.KeepAlive(mw)
	return ret
}

// Unsupported in ImageMagick 6.7.7
// Takes a frame optimized GIF animation, and compares the overlayed pixels
// against the disposal image resulting from all the previous frames in the
// animation. Any pixel that does not change the disposal image (and thus does
// not effect the outcome of an overlay) is made transparent.
// WARNING: This modifies the current images directly, rather than generate a
// new image sequence.
func (mw *MagickWand) OptimizeImageTransparency() error {
	ok := C.MagickOptimizeImageTransparency(mw.mw)
	return mw.getLastErrorIfFailed(ok)
}

// This is like ReadImage() except the only valid information returned is the
// image width, height, size, and format. It is designed to efficiently obtain
// this information from a file without reading the entire image sequence into
// memory.
func (mw *MagickWand) PingImage(filename string) error {
	csfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(csfilename))
	ok := C.MagickPingImage(mw.mw, csfilename)
	return mw.getLastErrorIfFailed(ok)
}

// Pings an image or image sequence from a blob.
func (mw *MagickWand) PingImageBlob(blob []byte) error {
	if len(blob) == 0 {
		return errors.New("zero-length blob not permitted")
	}
	ok := C.MagickPingImageBlob(mw.mw, unsafe.Pointer(&blob[0]), C.size_t(len(blob)))
	return mw.getLastErrorIfFailed(ok)
}

// Pings an image or image sequence from an open file descriptor.
func (mw *MagickWand) PingImageFile(img *os.File) error {
	file, err := cfdopen(img, "rb")
	if err != nil {
		return err
	}
	defer C.fclose(file)
	ok := C.MagickPingImageFile(mw.mw, file)
	return mw.getLastErrorIfFailed(ok)
}

// Simulates a Polaroid picture.
//
// caption: the Polaroid caption.
//
// angle: Apply the effect along this angle.
//
// method: the pixel interpolation method.
//
func (mw *MagickWand) PolaroidImage(dw *DrawingWand, caption string,
	angle float64, method PixelInterpolateMethod) error {

	ccaption := C.CString(caption)
	defer C.free(unsafe.Pointer(ccaption))

	ok := C.MagickPolaroidImage(mw.mw, dw.dw,
		ccaption,
		C.double(angle),
		C.PixelInterpolateMethod(method))

	runtime.KeepAlive(dw)
	return mw.getLastErrorIfFailed(ok)
}

// Reduces the image to a limited number of color level.
//
// levels: Number of color levels allowed in each channel. Very low values
// (2, 3, or 4) have the most visible effect.
//
// dither: choose the dither method: UndefinedDitherMethod, NoDitherMethod,
// RiemersmaDitherMethod, or FloydSteinbergDitherMethod.
//
func (mw *MagickWand) PosterizeImage(levels uint, dither DitherMethod) error {
	ok := C.MagickPosterizeImage(mw.mw, C.size_t(levels), C.DitherMethod(dither))
	return mw.getLastErrorIfFailed(ok)
}

// Tiles 9 thumbnails of the specified image with an image processing
// operation applied at varying strengths. This helpful to quickly pin-point
// an appropriate parameter for an image processing operation.
func (mw *MagickWand) PreviewImages(preview PreviewType) *MagickWand {
	ret := newMagickWand(C.MagickPreviewImages(mw.mw, C.PreviewType(preview)))
	runtime.KeepAlive(mw)
	return ret
}

// Sets the previous image in the wand as the current image. It is typically
// used after SetLastIterator(), after which its first use will set the last
// image as the current image (unless the wand is empty). It will return false
// when no more images are left to be returned which happens when the wand is
// empty, or the current image is the first image. At that point the iterator
// is than reset to again process images in the forward direction, again
// starting with the first image in list. Images added at this point are
// prepended. Also at that point any images added to the wand using AddImages()
// or ReadImages() will be prepended before the first image. In this sense the
// condition is not quite exactly the same as ResetIterator().
func (mw *MagickWand) PreviousImage() bool {
	ret := 1 == C.MagickPreviousImage(mw.mw)
	runtime.KeepAlive(mw)
	return ret
}

// Analyzes the colors within a reference image and chooses a fixed number of
// colors to represent the image. The goal of the algorithm is to minimize the
// color difference between the input and output image while minimizing the
// processing time.
//
// numColors: the number of colors.
//
// colorspace: Perform color reduction in this colorspace, typically
// RGBColorspace.
//
// treedepth: Normally, this integer value is zero or one. A zero or one tells
// Quantize to choose a optimal tree depth of Log4(number_colors). A tree of
// this depth generally allows the best representation of the reference image
// with the least amount of memory and the fastest computational speed. In
// some cases, such as an image with low color dispersion (a few number of
// colors), a value other than Log4(number_colors) is required. To expand the
// color tree completely, use a value of 8.
//
// dither: choose from UndefinedDitherMethod, NoDitherMethod,
// RiemersmaDitherMethod, FloydSteinbergDitherMethod.
//
// measureError: A value other than zero measures the difference between the
// original and quantized images. This difference is the total quantization
// error. The error is computed by summing over all pixels in an image the
// distance squared in RGB space between each reference pixel value and its
// quantized value.
//
func (mw *MagickWand) QuantizeImage(numColors uint, colorspace ColorspaceType,
	treedepth uint, dither DitherMethod, measureError bool) error {

	ok := C.MagickQuantizeImage(mw.mw, C.size_t(numColors),
		C.ColorspaceType(colorspace), C.size_t(treedepth),
		C.DitherMethod(dither), b2i(measureError))

	return mw.getLastErrorIfFailed(ok)
}

// Analyzes the colors within a sequence of images and chooses a fixed number
// of colors to represent the image. The goal of the algorithm is to minimize
// the color difference between the input and output image while minimizing the
// processing time.
//
// numColors: the number of colors.
//
// colorspace: Perform color reduction in this colorspace, typically
// RGBColorspace.
//
// treedepth: Normally, this integer value is zero or one. A zero or one tells
// Quantize to choose a optimal tree depth of Log4(number_colors). A tree of
// this depth generally allows the best representation of the reference image
// with the least amount of memory and the fastest computational speed. In
// some cases, such as an image with low color dispersion (a few number of
// colors), a value other than Log4(number_colors) is required. To expand the
// color tree completely, use a value of 8.
//
// dither: choose from these dither methods: NoDitherMethod,
// RiemersmaDitherMethod, or FloydSteinbergDitherMethod.
//
// measureError: A value other than zero measures the difference between the
// original and quantized images. This difference is the total quantization
// error. The error is computed by summing over all pixels in an image the
// distance squared in RGB space between each reference pixel value and its
// quantized value.
//
func (mw *MagickWand) QuantizeImages(numColors uint, colorspace ColorspaceType,
	treedepth uint, dither DitherMethod, measureError bool) error {

	ok := C.MagickQuantizeImages(mw.mw, C.size_t(numColors),
		C.ColorspaceType(colorspace), C.size_t(treedepth),
		C.DitherMethod(dither), b2i(measureError))

	return mw.getLastErrorIfFailed(ok)
}

// Radial blurs an image.
//
// angle: the angle of the blur in degrees.
//
func (mw *MagickWand) RadialBlurImage(angle float64) error {
	ok := C.MagickRotationalBlurImage(mw.mw, C.double(angle))
	return mw.getLastErrorIfFailed(ok)
}

// Creates a simulated three-dimensional button-like effect by lightening and
// darkening the edges of the image. Members width and height of raise_info
// define the width of the vertical and horizontal edge of the effect. width,
//
// height, x, y: Define the dimensions of the area to raise.
//
// raise: A value other than zero creates a 3-D raise effect, otherwise it has
// a lowered effect.
//
func (mw *MagickWand) RaiseImage(width uint, height uint, x int, y int, raise bool) error {
	ok := C.MagickRaiseImage(mw.mw, C.size_t(width), C.size_t(height), C.ssize_t(x), C.ssize_t(y), b2i(raise))
	return mw.getLastErrorIfFailed(ok)
}

// Changes the value of individual pixels based on the intensity of each pixel
// compared to threshold. The result is a high-contrast, two color image.
//
// low, high: Specify the high and low thresholds. These values range from 0
// to QuantumRange.
//
func (mw *MagickWand) RandomThresholdImage(low, high float64) error {
	ok := C.MagickRandomThresholdImage(mw.mw, C.double(low), C.double(high))
	return mw.getLastErrorIfFailed(ok)
}

// Reads an image or image sequence. The images are inserted at the current
// image pointer position. Use SetFirstIterator(), SetLastIterator, or
// SetImageIndex() to specify the current image pointer position at the
// beginning of the image list, the end, or anywhere in-between respectively.
func (mw *MagickWand) ReadImage(filename string) error {
	csfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(csfilename))
	ok := C.MagickReadImage(mw.mw, csfilename)
	return mw.getLastErrorIfFailed(ok)
}

// Reads an image or image sequence from a blob.
func (mw *MagickWand) ReadImageBlob(blob []byte) error {
	if len(blob) == 0 {
		return errors.New("zero-length blob not permitted")
	}
	ok := C.MagickReadImageBlob(mw.mw, unsafe.Pointer(&blob[0]), C.size_t(len(blob)))
	return mw.getLastErrorIfFailed(ok)
}

// Reads an image or image sequence from an open file descriptor.
func (mw *MagickWand) ReadImageFile(img *os.File) error {
	file, err := cfdopen(img, "rb")
	if err != nil {
		return err
	}
	defer C.fclose(file)
	ok := C.MagickReadImageFile(mw.mw, file)
	return mw.getLastErrorIfFailed(ok)
}

// Replaces the colors of an image with the closest color from a reference image.
//
// method: choose from these dither methods: NoDitherMethod, RiemersmaDitherMethod, or FloydSteinbergDitherMethod.
//
func (mw *MagickWand) RemapImage(remap *MagickWand, method DitherMethod) error {
	ok := C.MagickRemapImage(mw.mw, remap.mw, C.DitherMethod(method))
	runtime.KeepAlive(remap)
	return mw.getLastErrorIfFailed(ok)
}

// Removes an image from the image list.
func (mw *MagickWand) RemoveImage() error {
	ok := C.MagickRemoveImage(mw.mw)
	return mw.getLastErrorIfFailed(ok)
}

// Resample image to desired resolution.
//
// xRes/yRes: the new image x/y resolution.
//
// filter: Image filter to use.
//
func (mw *MagickWand) ResampleImage(xRes, yRes float64, filter FilterType) error {
	ok := C.MagickResampleImage(mw.mw, C.double(xRes), C.double(yRes), C.FilterType(filter))
	return mw.getLastErrorIfFailed(ok)
}

// Resets the Wand page canvas and position.
// page: the relative page specification.
func (mw *MagickWand) ResetImagePage(page string) error {
	cspage := C.CString(page)
	defer C.free(unsafe.Pointer(cspage))
	ok := C.MagickResetImagePage(mw.mw, cspage)
	return mw.getLastErrorIfFailed(ok)
}

// Scales an image to the desired dimensions
//
// cols: the number of cols in the scaled image.
//
// rows: the number of rows in the scaled image.
//
// filter: Image filter to use.
//
func (mw *MagickWand) ResizeImage(cols, rows uint, filter FilterType) error {
	ok := C.MagickResizeImage(mw.mw, C.size_t(cols), C.size_t(rows), C.FilterType(filter))
	return mw.getLastErrorIfFailed(ok)
}

// Offsets an image as defined by x and y.
//
// x: the x offset.
//
// y: the y offset.
//
func (mw *MagickWand) RollImage(x, y int) error {
	ok := C.MagickRollImage(mw.mw, C.ssize_t(x), C.ssize_t(y))
	return mw.getLastErrorIfFailed(ok)
}

// Rotates an image the specified number of degrees. Empty triangles left over
// from rotating the image are filled with the background color.
//
// background: the background pixel wand.
//
// degrees: the number of degrees to rotate the image.
//
func (mw *MagickWand) RotateImage(background *PixelWand, degrees float64) error {
	ok := C.MagickRotateImage(mw.mw, background.pw, C.double(degrees))
	runtime.KeepAlive(background)
	return mw.getLastErrorIfFailed(ok)
}

// Scales an image to the desired dimensions with pixel sampling. Unlike other
// scaling methods, this method does not introduce any additional color into
// the scaled image.
func (mw *MagickWand) SampleImage(cols, rows uint) error {
	ok := C.MagickSampleImage(mw.mw, C.size_t(cols), C.size_t(rows))
	return mw.getLastErrorIfFailed(ok)
}

// Scales the size of an image to the given dimensions.
func (mw *MagickWand) ScaleImage(cols, rows uint) error {
	ok := C.MagickScaleImage(mw.mw, C.size_t(cols), C.size_t(rows))
	return mw.getLastErrorIfFailed(ok)
}

// Segments an image by analyzing the histograms of the color components and
// identifying units that are homogeneous with the fuzzy C-means technique.
//
// verbose: set to true to print detailed information about the identified
// classes.
//
// clusterThreshold: this represents the minimum number of pixels contained in
// a hexahedra before it can be considered valid (expressed as a percentage).
//
// smoothThreshold: the smoothing threshold eliminates noise in the second
// derivative of the histogram. As the value is increased, you can expect a
// smoother second derivative.
//
func (mw *MagickWand) SegmentImage(colorspace ColorspaceType, verbose bool, clusterThreshold, smoothThreshold float64) error {
	ok := C.MagickSegmentImage(mw.mw, C.ColorspaceType(colorspace), b2i(verbose), C.double(clusterThreshold), C.double(smoothThreshold))
	return mw.getLastErrorIfFailed(ok)
}

// Selectively blur an image within a contrast threshold. It is similar to the
// unsharpen mask that sharpens everything with contrast above a certain
// threshold.
//
// radius: the radius of the gaussian, in pixels, not counting the center pixel.
//
// sigma: the standard deviation of the gaussian, in pixels.
//
// threshold: only pixels within this contrast threshold are included in the
// blur operation.
//
func (mw *MagickWand) SelectiveBlurImage(radius, sigma, threshold float64) error {
	ok := C.MagickSelectiveBlurImage(mw.mw, C.double(radius), C.double(sigma), C.double(threshold))
	return mw.getLastErrorIfFailed(ok)
}

// Applies a special effect to the image, similar to the effect achieved in a
// photo darkroom by sepia toning. Threshold ranges from 0 to QuantumRange and
// is a measure of the extent of the sepia toning. A threshold of 80 is a good
// starting point for a reasonable tone.
//
// threshold: Define the extent of the sepia toning.
//
func (mw *MagickWand) SepiaToneImage(threshold float64) error {
	ok := C.MagickSepiaToneImage(mw.mw, C.double(threshold))
	return mw.getLastErrorIfFailed(ok)
}

// Replaces the last image returned by SetImageIndex(), NextImage(),
// PreviousImage() with the images from the specified wand.
func (mw *MagickWand) SetImage(source *MagickWand) error {
	ok := C.MagickSetImage(mw.mw, source.mw)
	runtime.KeepAlive(source)
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image to the specified alpha level.
//
// alpha: the level of transparency: 1.0 is fully opaque and 0.0 is fully
// transparent.
//
func (mw *MagickWand) SetImageAlpha(alpha float64) error {
	ok := C.MagickSetImageAlpha(mw.mw, C.double(alpha))
	runtime.KeepAlive(mw)
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image background color.
func (mw *MagickWand) SetImageBackgroundColor(background *PixelWand) error {
	ok := C.MagickSetImageBackgroundColor(mw.mw, background.pw)
	runtime.KeepAlive(background)
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image chromaticity blue primary point.
func (mw *MagickWand) SetImageBluePrimary(x, y, z float64) error {
	ok := C.MagickSetImageBluePrimary(mw.mw,
		C.double(x), C.double(y), C.double(z))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image border color.
func (mw *MagickWand) SetImageBorderColor(border *PixelWand) error {
	ok := C.MagickSetImageBorderColor(mw.mw, border.pw)
	runtime.KeepAlive(border)
	return mw.getLastErrorIfFailed(ok)
}

// SetImageChannelMask sets image channel mask, and returns
// the previous channel type setting.
// Setting the channel mask impacts the effect of image operations,
// to limit the operation to the given channel.
func (mw *MagickWand) SetImageChannelMask(channel ChannelType) ChannelType {
	prevChannel := C.MagickSetImageChannelMask(mw.mw, C.ChannelType(channel))
	return ChannelType(prevChannel)
}

// Set the entire wand canvas to the specified color.
func (mw *MagickWand) SetImageColor(color *PixelWand) error {
	ok := C.MagickSetImageColor(mw.mw, color.pw)
	runtime.KeepAlive(color)
	return mw.getLastErrorIfFailed(ok)
}

// SetImageAlphaChannel activates, deactivates, resets, or sets the alpha channel.
func (mw *MagickWand) SetImageAlphaChannel(operation AlphaChannelType) error {
	ok := C.MagickSetImageAlphaChannel(mw.mw, C.AlphaChannelOption(operation))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the color of the specified colormap index.
//
// index: the offset into the image colormap.
//
// color: return the colormap color in this wand.
//
func (mw *MagickWand) SetImageColormapColor(index uint, color *PixelWand) error {
	ok := C.MagickSetImageColormapColor(mw.mw, C.size_t(index), color.pw)
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image colorspace.
func (mw *MagickWand) SetImageColorspace(colorspace ColorspaceType) error {
	ok := C.MagickSetImageColorspace(mw.mw, C.ColorspaceType(colorspace))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image composite operator, useful for specifying how to composite
/// the image thumbnail when using the MontageImage() method.
func (mw *MagickWand) SetImageCompose(compose CompositeOperator) error {
	ok := C.MagickSetImageCompose(mw.mw, C.CompositeOperator(compose))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image compression.
func (mw *MagickWand) SetImageCompression(compression CompressionType) error {
	ok := C.MagickSetImageCompression(mw.mw, C.CompressionType(compression))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image compression quality.
func (mw *MagickWand) SetImageCompressionQuality(quality uint) error {
	ok := C.MagickSetImageCompressionQuality(mw.mw, C.size_t(quality))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image delay.
//
// delay: the image delay in ticks-per-second units.
//
func (mw *MagickWand) SetImageDelay(delay uint) error {
	ok := C.MagickSetImageDelay(mw.mw, C.size_t(delay))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image depth.
//
// depth: the image depth in bits: 8, 16, or 32.
//
func (mw *MagickWand) SetImageDepth(depth uint) error {
	ok := C.MagickSetImageDepth(mw.mw, C.size_t(depth))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image disposal method.
func (mw *MagickWand) SetImageDispose(dispose DisposeType) error {
	ok := C.MagickSetImageDispose(mw.mw, C.DisposeType(dispose))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image endian method.
func (mw *MagickWand) SetImageEndian(endian EndianType) error {
	ok := C.MagickSetImageEndian(mw.mw, C.EndianType(endian))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image size (i.e. cols & rows).
//
// cols: The image width in pixels.
//
// rows: The image height in pixels.
//
func (mw *MagickWand) SetImageExtent(cols, rows uint) error {
	ok := C.MagickSetImageExtent(mw.mw, C.size_t(cols), C.size_t(rows))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the filename of a particular image in a sequence.
func (mw *MagickWand) SetImageFilename(filename string) error {
	csfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(csfilename))
	ok := C.MagickSetImageFilename(mw.mw, csfilename)
	return mw.getLastErrorIfFailed(ok)
}

// Sets the format of a particular image in a sequence.
//
// format: the image format.
//
func (mw *MagickWand) SetImageFormat(format string) error {
	csformat := C.CString(format)
	defer C.free(unsafe.Pointer(csformat))
	ok := C.MagickSetImageFormat(mw.mw, csformat)
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image fuzz.
func (mw *MagickWand) SetImageFuzz(fuzz float64) error {
	ok := C.MagickSetImageFuzz(mw.mw, C.double(fuzz))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image gamma.
func (mw *MagickWand) SetImageGamma(gamma float64) error {
	ok := C.MagickSetImageGamma(mw.mw, C.double(gamma))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image gravity type.
func (mw *MagickWand) SetImageGravity(gravity GravityType) error {
	ok := C.MagickSetImageGravity(mw.mw, C.GravityType(gravity))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image chromaticity green primary point.
func (mw *MagickWand) SetImageGreenPrimary(x, y, z float64) error {
	ok := C.MagickSetImageGreenPrimary(mw.mw, C.double(x), C.double(y), C.double(z))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image interlace scheme.
func (mw *MagickWand) SetImageInterlaceScheme(interlace InterlaceType) error {
	ok := C.MagickSetImageInterlaceScheme(mw.mw, C.InterlaceType(interlace))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image interpolate pixel method.
func (mw *MagickWand) SetImageInterpolateMethod(method PixelInterpolateMethod) error {
	ok := C.MagickSetImageInterpolateMethod(mw.mw, C.PixelInterpolateMethod(method))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image iterations.
func (mw *MagickWand) SetImageIterations(iterations uint) error {
	ok := C.MagickSetImageIterations(mw.mw, C.size_t(iterations))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image matte channel.
func (mw *MagickWand) SetImageMatte(matte bool) error {
	ok := C.MagickSetImageMatte(mw.mw, b2i(matte))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image orientation.
func (mw *MagickWand) SetImageOrientation(orientation OrientationType) error {
	ok := C.MagickSetImageOrientation(mw.mw, C.OrientationType(orientation))
	return mw.getLastErrorIfFailed(ok)
}

// Auto orient the image
func (mw *MagickWand) AutoOrientImage() error {
	ok := C.MagickAutoOrientImage(mw.mw)
	return mw.getLastErrorIfFailed(ok)
}

// Sets the page geometry of the image.
func (mw *MagickWand) SetImagePage(width, height uint, x, y int) error {
	ok := C.MagickSetImagePage(mw.mw, C.size_t(width), C.size_t(height), C.ssize_t(x), C.ssize_t(y))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image chromaticity red primary point.
func (mw *MagickWand) SetImageRedPrimary(x, y, z float64) error {
	ok := C.MagickSetImageRedPrimary(mw.mw,
		C.double(x), C.double(y), C.double(z))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image rendering intent.
func (mw *MagickWand) SetImageRenderingIntent(ri RenderingIntent) error {
	ok := C.MagickSetImageRenderingIntent(mw.mw, C.RenderingIntent(ri))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image resolution.
func (mw *MagickWand) SetImageResolution(xRes, yRes float64) error {
	ok := C.MagickSetImageResolution(mw.mw, C.double(xRes), C.double(yRes))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image scene.
func (mw *MagickWand) SetImageScene(scene uint) error {
	ok := C.MagickSetImageScene(mw.mw, C.size_t(scene))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image ticks-per-second.
func (mw *MagickWand) SetImageTicksPerSecond(tps int) error {
	ok := C.MagickSetImageTicksPerSecond(mw.mw, C.ssize_t(tps))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image type.
func (mw *MagickWand) SetImageType(imgtype ImageType) error {
	ok := C.MagickSetImageType(mw.mw, C.ImageType(imgtype))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image units of resolution.
func (mw *MagickWand) SetImageUnits(units ResolutionType) error {
	ok := C.MagickSetImageUnits(mw.mw, C.ResolutionType(units))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image virtual pixel method.
func (mw *MagickWand) SetImageVirtualPixelMethod(method VirtualPixelMethod) VirtualPixelMethod {
	return VirtualPixelMethod(C.MagickSetImageVirtualPixelMethod(mw.mw, C.VirtualPixelMethod(method)))
}

// Sets the image chromaticity white point.
func (mw *MagickWand) SetImageWhitePoint(x, y, z float64) error {
	ok := C.MagickSetImageWhitePoint(mw.mw, C.double(x), C.double(y), C.double(z))
	return mw.getLastErrorIfFailed(ok)
}

// Shines a distant light on an image to create a three-dimensional effect.
// You control the positioning of the light with azimuth and elevation;
// azimuth is measured in degrees off the x axis and elevation is measured in
// pixels above the Z axis.
//
// gray: if true, shades the intensity of each pixel.
//
// azimuth, elevation: define the light source direction.
//
func (mw *MagickWand) ShadeImage(gray bool, azimuth, elevation float64) error {
	ok := C.MagickShadeImage(mw.mw, b2i(gray), C.double(azimuth), C.double(elevation))
	return mw.getLastErrorIfFailed(ok)
}

// Simulates an image shadow.
//
// opacity: percentage transparency.
//
// sigma: the standard deviation of the Gaussian, in pixels.
//
// x: the shadow x-offset.
//
// y: the shadow y-offset.
//
func (mw *MagickWand) ShadowImage(opacity, sigma float64, x, y int) error {
	ok := C.MagickShadowImage(mw.mw, C.double(opacity), C.double(sigma), C.ssize_t(x), C.ssize_t(y))
	return mw.getLastErrorIfFailed(ok)
}

// Sharpens an image. We convolve the image with a Gaussian operator of the
// given radius and standard deviation (sigma). For reasonable results, the
// radius should be larger than sigma. Use a radius of 0 and SharpenImage()
// selects a suitable radius for you.
//
// radius: the radius of the Gaussian, in pixels, not counting the center pixel.
//
// sigma: the standard deviation of the Gaussian, in pixels.
//
func (mw *MagickWand) SharpenImage(radius, sigma float64) error {
	ok := C.MagickSharpenImage(mw.mw, C.double(radius), C.double(sigma))
	return mw.getLastErrorIfFailed(ok)
}

// Shaves pixels from the image edges. It allocates the memory necessary for
// the new Image structure and returns a pointer to the new image.
func (mw *MagickWand) ShaveImage(cols, rows uint) error {
	ok := C.MagickShaveImage(mw.mw, C.size_t(cols), C.size_t(rows))
	return mw.getLastErrorIfFailed(ok)
}

// Slides one edge of an image along the X or Y axis, creating a parallelogram.
// An X direction shear slides an edge along the X axis, while a Y direction
// shear slides an edge along the Y axis. The amount of the shear is controlled
// by a shear angle. For X direction shears, xShear is measured relative to the
// Y axis, and similarly, for Y direction shears yShear is measured relative to
// the X axis. Empty triangles left over from shearing the image are filled
// with the background color.
func (mw *MagickWand) ShearImage(background *PixelWand, xShear, yShear float64) error {
	ok := C.MagickShearImage(mw.mw, background.pw, C.double(xShear), C.double(yShear))
	runtime.KeepAlive(background)
	return mw.getLastErrorIfFailed(ok)
}

// Adjusts the contrast of an image with a non-linear sigmoidal contrast
// algorithm. Increase the contrast of the image using a sigmoidal transfer
// function without saturating highlights or shadows. Contrast indicates how
// much to increase the contrast (0 is none; 3 is typical; 20 is pushing it);
// mid-point indicates where midtones fall in the resultant image (0 is white;
// 50 is middle-gray; 100 is black). Set sharpen to true to increase the image
// contrast otherwise the contrast is reduced.
//
// sharpen: Increase or decrease image contrast.
//
// alpha: strength of the contrast, the larger the number the more
// 'threshold-like' it becomes.
//
// beta: midpoint of the function as a color value 0 to QuantumRange.
//
func (mw *MagickWand) SigmoidalContrastImage(sharpen bool, alpha, beta float64) error {
	ok := C.MagickSigmoidalContrastImage(mw.mw, b2i(sharpen), C.double(alpha), C.double(beta))
	return mw.getLastErrorIfFailed(ok)
}

// Compares the reference image of the image and returns the best match offset.
// In addition, it returns a similarity image such that an exact match location
// is completely white and if none of the pixels match, black, otherwise some
// gray level in-between.
//
// reference: the reference wand.
//
// offset: the best match offset of the reference image within the image.
//
// metric: the metric.
//
// threshold: similarity; minimum distortion for (sub)image match.
//
func (mw *MagickWand) SimilarityImage(reference *MagickWand, metric MetricType,
	threshold float64) (offset *RectangleInfo, similarity float64, area *MagickWand) {

	var rectInfo C.RectangleInfo

	mwarea := C.MagickSimilarityImage(mw.mw, reference.mw,
		C.MetricType(metric), C.double(similarity),
		&rectInfo, (*C.double)(&similarity))

	runtime.KeepAlive(reference)
	return &RectangleInfo{&rectInfo}, similarity, newMagickWand(mwarea)
}

// Simulates a pencil sketch. We convolve the image with a Gaussian operator
// of the given radius and standard deviation (sigma). For reasonable results,
// radius should be larger than sigma. Use a radius of 0 and SketchImage()
// selects a suitable radius for you. Angle gives the angle of the blurring
// motion.
//
// radius: the radius of the Gaussian, in pixels, not counting the center pixel.
//
// sigma: the standard deviation of the Gaussian, in pixels.
//
// angle: Apply the effect along this angle.
//
func (mw *MagickWand) SketchImage(radius, sigma, angle float64) error {
	ok := C.MagickSketchImage(mw.mw, C.double(radius), C.double(sigma), C.double(angle))
	return mw.getLastErrorIfFailed(ok)
}

// Takes all images from the current image pointer to the end of the image
// list and smushs them to each other top-to-bottom if the stack parameter is
// true, otherwise left-to-right.
//
// stack: by default, images are stacked left-to-right. Set stack to true to
// stack them top-to-bottom.
//
// offset: minimum distance in pixels between images.
//
func (mw *MagickWand) SmushImages(stack bool, offset int) *MagickWand {
	ret := newMagickWand(C.MagickSmushImages(mw.mw, b2i(stack), C.ssize_t(offset)))
	runtime.KeepAlive(mw)
	return ret
}

// Applies a special effect to the image, similar to the effect achieved in a
// photo darkroom by selectively exposing areas of photo sensitive paper to
// light. Threshold ranges from 0 to QuantumRange and is a measure of the
// extent of the solarization.
//
// threshold: define the extent of the solarization.
//
func (mw *MagickWand) SolarizeImage(threshold float64) error {
	ok := C.MagickSolarizeImage(mw.mw, C.double(threshold))
	return mw.getLastErrorIfFailed(ok)
}

// Unsupported in ImageMagick 6.7.7
// Applies a special effect to the image's channel, similar to the effect
// achieved in a photo darkroom by selectively exposing areas of photo
// sensitive paper to light. Threshold ranges from 0 to QuantumRange and is a
// measure of the extent of the solarization.
//
// threshold: define the extent of the solarization.
//
//func (mw *MagickWand) SolarizeImageChannel(channel ChannelType, threshold float64) error {
//	ok := C.MagickSolarizeImageChannel(mw.mw, C.ChannelType(channel), C.double(threshold))
//	return mw.getLastErrorIfFailed(ok)
//}

// Given a set of coordinates, interpolates the colors found at those
// coordinates, across the whole image, using various methods.
//
// method: the method of image sparseion. ArcSparseColorion will always ignore
// source image offset, and always 'bestfit' the destination image with the top
// left corner offset relative to the polar mapping center. Bilinear has no
// simple inverse mapping so will not allow 'bestfit' style of image sparseion.
// Affine, Perspective, and Bilinear, will do least squares fitting of the
// distortion when more than the minimum number of control point pairs are
// provided. Perspective, and Bilinear, will fall back to a Affine sparseion
// when less than 4 control point pairs are provided. While Affine sparseions
// will let you use any number of control point pairs, that is Zero pairs is a
// No-Op (viewport only) distortion, one pair is a translation and two pairs
// of control points will do a scale-rotate-translate, without any shearing.
//
// arguments: the arguments for this sparseion method.
//
func (mw *MagickWand) SparseColorImage(channel ChannelType, method SparseColorMethod,
	arguments []float64) error {

	ok := C.MagickSparseColorImage(mw.mw, C.SparseColorMethod(method),
		C.size_t(len(arguments)), (*C.double)(&arguments[0]))

	return mw.getLastErrorIfFailed(ok)
}

// Splices a solid color into the image.
func (mw *MagickWand) SpliceImage(width, height uint, x, y int) error {
	ok := C.MagickSpliceImage(mw.mw, C.size_t(width), C.size_t(height), C.ssize_t(x), C.ssize_t(y))
	return mw.getLastErrorIfFailed(ok)
}

// Is a special effects method that randomly displaces each pixel in a block
// defined by the radius parameter.
//
// method: intepolation method.
//
// radius: Choose a random pixel in a neighborhood of this extent.
//
func (mw *MagickWand) SpreadImage(method PixelInterpolateMethod, radius float64) error {
	ok := C.MagickSpreadImage(mw.mw, C.PixelInterpolateMethod(method), C.double(radius))
	return mw.getLastErrorIfFailed(ok)
}

// Replace each pixel with corresponding statistic from the neighborhood of
// the specified width and height.
//
// type: the statistic type (e.g. median, mode, etc.).
//
// width: the width of the pixel neighborhood.
//
// height: the height of the pixel neighborhood.
//
func (mw *MagickWand) StatisticImage(stype StatisticType, width, height uint) error {
	ok := C.MagickStatisticImage(mw.mw, C.StatisticType(stype), C.size_t(width), C.size_t(height))
	return mw.getLastErrorIfFailed(ok)
}

// Hides a digital watermark within the image. Recover the hidden watermark
// later to prove that the authenticity of an image. Offset defines the start
// position within the image to hide the watermark.
//
// offset: start hiding at this offset into the image.
//
func (mw *MagickWand) SteganoImage(watermark *MagickWand, offset int) *MagickWand {
	ret := newMagickWand(C.MagickSteganoImage(mw.mw, watermark.mw, C.ssize_t(offset)))
	runtime.KeepAlive(mw)
	runtime.KeepAlive(watermark)
	return ret
}

// Composites two images and produces a single image that is the composite of
// a left and right image of a stereo pair.
func (mw *MagickWand) StereoImage(offset *MagickWand) *MagickWand {
	ret := newMagickWand(C.MagickStereoImage(mw.mw, offset.mw))
	runtime.KeepAlive(mw)
	runtime.KeepAlive(offset)
	return ret
}

// Strips an image of all profiles and comments.
func (mw *MagickWand) StripImage() error {
	ok := C.MagickStripImage(mw.mw)
	return mw.getLastErrorIfFailed(ok)
}

// Swirls the pixels about the center of the image, where degrees indicates the
// sweep of the arc through which each pixel is moved. You get a more dramatic
// effect as the degrees move from 1 to 360.
//
// degrees: define the tightness of the swirling effect.
//
// method: the pixel interpolation method.
//
func (mw *MagickWand) SwirlImage(degrees float64, method PixelInterpolateMethod) error {
	ok := C.MagickSwirlImage(mw.mw, C.double(degrees), C.PixelInterpolateMethod(method))
	return mw.getLastErrorIfFailed(ok)
}

// Repeatedly tiles the texture image across and down the image canvas.
func (mw *MagickWand) TextureImage(texture *MagickWand) *MagickWand {
	ret := newMagickWand(C.MagickTextureImage(mw.mw, texture.mw))
	runtime.KeepAlive(mw)
	runtime.KeepAlive(texture)
	return ret
}

// Changes the value of individual pixels based on the intensity of each pixel
// compared to threshold. The result is a high-contrast, two color image.
//
// threshold: define the threshold value.
//
func (mw *MagickWand) ThresholdImage(threshold float64) error {
	ok := C.MagickThresholdImage(mw.mw, C.double(threshold))
	return mw.getLastErrorIfFailed(ok)
}

// Changes the value of individual pixels based on the intensity of each pixel
// compared to threshold. The result is a high-contrast, two color image.
//
// threshold: define the threshold value.
//
func (mw *MagickWand) ThresholdImageChannel(channel ChannelType, threshold float64) error {
	ok := C.MagickThresholdImageChannel(mw.mw, C.ChannelType(channel), C.double(threshold))
	return mw.getLastErrorIfFailed(ok)
}

// Changes the size of an image to the given dimensions and removes any
// associated profiles. The goal is to produce small low cost thumbnail images
// suited for display on the Web.
func (mw *MagickWand) ThumbnailImage(cols, rows uint) error {
	ok := C.MagickThumbnailImage(mw.mw, C.size_t(cols), C.size_t(rows))
	return mw.getLastErrorIfFailed(ok)
}

// Applies a color vector to each pixel in the image. The length of the vector
// is 0 for black and white and at its maximum for the midtones. The vector
// weighting function is f(x)=(1-(4.0*((x-0.5)*(x-0.5)))).
//
// tint: the tint pixel wand.
//
// opacity: the opacity pixel wand.
//
func (mw *MagickWand) TintImage(tint, opacity *PixelWand) error {
	ok := C.MagickTintImage(mw.mw, tint.pw, opacity.pw)
	runtime.KeepAlive(tint)
	runtime.KeepAlive(opacity)
	return mw.getLastErrorIfFailed(ok)
}

// Transform the image colorspace, setting the images colorspace while
// transforming the images data to that colorspace.
func (mw *MagickWand) TransformImageColorspace(colorspace ColorspaceType) error {
	ok := C.MagickTransformImageColorspace(mw.mw, C.ColorspaceType(colorspace))
	return mw.getLastErrorIfFailed(ok)
}

// Changes any pixel that matches color with the color defined by fill.
//
// alpha: the level of transparency: 1.0 is fully opaque and 0.0 is fully
// transparent.
//
// fuzz: by default target must match a particular pixel color exactly.
// However, in many cases two colors may differ by a small amount. The fuzz
// member of image defines how much tolerance is acceptable to consider two
// colors as the same. For example, set fuzz to 10 and the color red at
// intensities of 100 and 102 respectively are now interpreted as the same
// color for the purposes of the floodfill.
//
// invert: paint any pixel that does not match the target color.
//
func (mw *MagickWand) TransparentPaintImage(target *PixelWand, alpha, fuzz float64, invert bool) error {
	ok := C.MagickTransparentPaintImage(mw.mw, target.pw, C.double(alpha), C.double(fuzz), b2i(invert))
	runtime.KeepAlive(target)
	return mw.getLastErrorIfFailed(ok)
}

// Creates a vertical mirror image by reflecting the pixels around the central
// x-axis while rotating them 90-degrees.
func (mw *MagickWand) TransposeImage() error {
	ok := C.MagickTransposeImage(mw.mw)
	return mw.getLastErrorIfFailed(ok)
}

// Creates a horizontal mirror image by reflecting the pixels around the
// central y-axis while rotating them 270-degrees.
func (mw *MagickWand) TransverseImage() error {
	ok := C.MagickTransverseImage(mw.mw)
	return mw.getLastErrorIfFailed(ok)
}

// Remove edges that are the background color from the image.
//
// fuzz: by default target must match a particular pixel color exactly.
// However, in many cases two colors may differ by a small amount. The fuzz
// member of image defines how much tolerance is acceptable to consider two
// colors as the same. For example, set fuzz to 10 and the color red at
// intensities of 100 and 102 respectively are now interpreted as the same
// color for the purposes of the floodfill.
func (mw *MagickWand) TrimImage(fuzz float64) error {
	ok := C.MagickTrimImage(mw.mw, C.double(fuzz))
	return mw.getLastErrorIfFailed(ok)
}

// Discards all but one of any pixel color.
func (mw *MagickWand) UniqueImageColors() error {
	ok := C.MagickUniqueImageColors(mw.mw)
	return mw.getLastErrorIfFailed(ok)
}

// Unsharpens an image. We convolve the image with a Gaussian operator of the
// given radius and standard deviation (sigma). For reasonable results, radius
// should be larger than sigma. Use a radius of 0 and UnsharpMaskImage()
// selects a suitable radius for you.
//
// radius: the radius of the Gaussian, in pixels, not counting the center pixel.
//
// sigma: the standard deviation of the Gaussian, in pixels.
//
// amount: the percentage of the difference between the original and the blur
// image that is added back into the original.
//
// threshold: the threshold in pixels needed to apply the diffence amount.
//
func (mw *MagickWand) UnsharpMaskImage(radius, sigma, amount, threshold float64) error {
	ok := C.MagickUnsharpMaskImage(mw.mw, C.double(radius), C.double(sigma), C.double(amount), C.double(threshold))
	return mw.getLastErrorIfFailed(ok)
}

// Softens the edges of the image in vignette style.
//
// x, y: define the x and y ellipse offset.
//
func (mw *MagickWand) VignetteImage(blackPoint, whitePoint float64, x, y int) error {
	ok := C.MagickVignetteImage(mw.mw, C.double(blackPoint), C.double(whitePoint), C.ssize_t(x), C.ssize_t(y))
	return mw.getLastErrorIfFailed(ok)
}

// Creates a "ripple" effect in the image by shifting the pixels vertically
// along a sine wave whose amplitude and wavelength is specified by the given
// parameters.
//
// amplitude, wavelength: Define the amplitude and wave length of the sine wave.
//
// method: the pixel interpolation method
//
func (mw *MagickWand) WaveImage(amplitude, wavelength float64,
	method PixelInterpolateMethod) error {

	ok := C.MagickWaveImage(mw.mw, C.double(amplitude),
		C.double(wavelength), C.PixelInterpolateMethod(method))

	return mw.getLastErrorIfFailed(ok)
}

// Is like ThresholdImage() but force all pixels above the threshold into white
// while leaving all pixels below the threshold unchanged.
func (mw *MagickWand) WhiteThresholdImage(threshold *PixelWand) error {
	ok := C.MagickWhiteThresholdImage(mw.mw, threshold.pw)
	runtime.KeepAlive(threshold)
	return mw.getLastErrorIfFailed(ok)
}

// Writes an image to the specified filename.
func (mw *MagickWand) WriteImage(filename string) error {
	csfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(csfilename))
	ok := C.MagickWriteImage(mw.mw, csfilename)
	return mw.getLastErrorIfFailed(ok)
}

// Writes an image to an open file descriptor.
func (mw *MagickWand) WriteImageFile(out *os.File) error {
	file, err := cfdopen(out, "w")
	if err != nil {
		return err
	}
	defer C.fclose(file)
	ok := C.MagickWriteImageFile(mw.mw, file)
	return mw.getLastErrorIfFailed(ok)
}

// Writes an image or image sequence.
func (mw *MagickWand) WriteImages(filename string, adjoin bool) error {
	csfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(csfilename))
	ok := C.MagickWriteImages(mw.mw, csfilename, b2i(adjoin))
	return mw.getLastErrorIfFailed(ok)
}

// Writes an image sequence to an open file descriptor.
func (mw *MagickWand) WriteImagesFile(out *os.File) error {
	file, err := cfdopen(out, "w")
	if err != nil {
		return err
	}
	defer C.fclose(file)
	ok := C.MagickWriteImagesFile(mw.mw, file)
	return mw.getLastErrorIfFailed(ok)
}

// cfdopen returns a C-level FILE*. mode should be as described in fdopen(3).
// Caller is responsible for closing the file when successfully returned,
// via C.fclose()
func cfdopen(file *os.File, mode string) (*C.FILE, error) {
	cmode := C.CString(mode)
	defer C.free(unsafe.Pointer(cmode))

	cfile, err := C.fdopen(C.dup(C.int(file.Fd())), cmode)
	if err != nil {
		return nil, err
	}
	if cfile == nil {
		return nil, syscall.EINVAL
	}

	return cfile, nil
}
