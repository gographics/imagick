// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"
import (
	//"fmt"
	"errors"
	"runtime"
	"unsafe"
)

// This method deletes a wand artifact
func (mw *MagickWand) DeleteImageArtifact(artifact string) error {
	csartifact := C.CString(artifact)
	defer C.free(unsafe.Pointer(csartifact))
	ok := C.MagickDeleteImageArtifact(mw.mw, csartifact)
	return mw.getLastErrorIfFailed(ok)
}

// This method deletes a image property
func (mw *MagickWand) DeleteImageProperty(property string) error {
	csproperty := C.CString(property)
	defer C.free(unsafe.Pointer(csproperty))
	ok := C.MagickDeleteImageProperty(mw.mw, csproperty)
	return mw.getLastErrorIfFailed(ok)
}

// This method deletes a wand option
func (mw *MagickWand) DeleteOption(option string) error {
	csoption := C.CString(option)
	defer C.free(unsafe.Pointer(csoption))
	ok := C.MagickDeleteOption(mw.mw, csoption)
	return mw.getLastErrorIfFailed(ok)
}

// Returns the antialias property associated with the wand
func (mw *MagickWand) GetAntialias() bool {
	ret := 1 == C.int(C.MagickGetAntialias(mw.mw))
	runtime.KeepAlive(mw)
	return ret
}

// Returns the wand background color
func (mw *MagickWand) GetBackgroundColor() *PixelWand {
	ret := newPixelWand(C.MagickGetBackgroundColor(mw.mw))
	runtime.KeepAlive(mw)
	return ret
}

// Returns the wand colorspace type
func (mw *MagickWand) GetColorspace() ColorspaceType {
	ccst := C.MagickGetColorspace(mw.mw)
	runtime.KeepAlive(mw)
	return ColorspaceType(ccst)
}

// Gets the wand compression type.
func (mw *MagickWand) GetCompression() CompressionType {
	ret := CompressionType(C.MagickGetCompression(mw.mw))
	runtime.KeepAlive(mw)
	return ret
}

// Gets the wand compression quality.
func (mw *MagickWand) GetCompressionQuality() uint {
	ret := uint(C.MagickGetCompressionQuality(mw.mw))
	runtime.KeepAlive(mw)
	return ret
}

// Returns the filename associated with an image sequence.
func (mw *MagickWand) GetFilename() string {
	cstr := C.MagickGetFilename(mw.mw)
	runtime.KeepAlive(mw)
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

// Returns the font associated with the MagickWand.
func (mw *MagickWand) GetFont() string {
	cstr := C.MagickGetFont(mw.mw)
	runtime.KeepAlive(mw)
	defer relinquishMemory(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

// Returns the format of the magick wand.
func (mw *MagickWand) GetFormat() string {
	cstr := C.MagickGetFormat(mw.mw)
	runtime.KeepAlive(mw)
	defer relinquishMemory(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

// Gets the wand gravity.
func (mw *MagickWand) GetGravity() GravityType {
	ret := GravityType(C.MagickGetGravity(mw.mw))
	runtime.KeepAlive(mw)
	return ret
}

// Returns a value associated with the specified artifact.
func (mw *MagickWand) GetImageArtifact(artifact string) string {
	csartifact := C.CString(artifact)
	defer C.free(unsafe.Pointer(csartifact))
	cstr := C.MagickGetImageArtifact(mw.mw, csartifact)
	runtime.KeepAlive(mw)
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

// Returns all the artifact names that match the specified pattern associated
// with a wand. Use GetImageProperty() to return the value of a particular
// artifact.
func (mw *MagickWand) GetImageArtifacts(pattern string) (artifacts []string) {
	cspattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(cspattern))
	num := C.size_t(0)
	p := C.MagickGetImageArtifacts(mw.mw, cspattern, &num)
	runtime.KeepAlive(mw)
	defer relinquishMemoryCStringArray(p)
	artifacts = sizedCStringArrayToStringSlice(p, num)
	return
}

// Returns the named image profile.
//
// name: Name of profile to return: ICC, IPTC, or generic profile.
func (mw *MagickWand) GetImageProfile(name string) string {
	csname := C.CString(name)
	defer C.free(unsafe.Pointer(csname))
	szlen := C.size_t(0)
	csprofile := C.MagickGetImageProfile(mw.mw, csname, &szlen)
	runtime.KeepAlive(mw)
	defer relinquishMemory(unsafe.Pointer(csprofile))
	return C.GoStringN((*C.char)((unsafe.Pointer)(csprofile)), C.int(szlen))
}

// Returns all the profile names that match the specified pattern associated
// with a wand. Use GetImageProfile() to return the value of a particular
// property.
func (mw *MagickWand) GetImageProfiles(pattern string) (profiles []string) {
	cspattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(cspattern))
	np := C.size_t(0)
	ps := C.MagickGetImageProfiles(mw.mw, cspattern, &np)
	runtime.KeepAlive(mw)
	defer relinquishMemoryCStringArray(ps)
	profiles = sizedCStringArrayToStringSlice(ps, np)
	return
}

// Returns a value associated with the specified property.
func (mw *MagickWand) GetImageProperty(property string) string {
	csproperty := C.CString(property)
	defer C.free(unsafe.Pointer(csproperty))
	cspv := C.MagickGetImageProperty(mw.mw, csproperty)
	runtime.KeepAlive(mw)
	defer relinquishMemory(unsafe.Pointer(cspv))
	return C.GoString(cspv)
}

// Returns all the property names that match the specified pattern associated
// with a wand. Use GetImageProperty() to return the value of a particular
// property.
func (mw *MagickWand) GetImageProperties(pattern string) (properties []string) {
	cspattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(cspattern))
	np := C.size_t(0)
	ps := C.MagickGetImageProperties(mw.mw, cspattern, &np)
	runtime.KeepAlive(mw)
	defer relinquishMemoryCStringArray(ps)
	properties = sizedCStringArrayToStringSlice(ps, np)
	return
}

// Gets the wand interlace scheme.
func (mw *MagickWand) GetInterlaceScheme() InterlaceType {
	ret := InterlaceType(C.MagickGetInterlaceScheme(mw.mw))
	runtime.KeepAlive(mw)
	return ret
}

// Gets the wand compression.
func (mw *MagickWand) GetInterpolateMethod() PixelInterpolateMethod {
	ret := PixelInterpolateMethod(C.MagickGetInterpolateMethod(mw.mw))
	runtime.KeepAlive(mw)
	return ret
}

// Returns a value associated with a wand and the specified key.
func (mw *MagickWand) GetOption(key string) string {
	cskey := C.CString(key)
	defer C.free(unsafe.Pointer(cskey))
	csval := C.MagickGetOption(mw.mw, cskey)
	runtime.KeepAlive(mw)
	defer relinquishMemory(unsafe.Pointer(csval))
	return C.GoString(csval)
}

// Returns all the option names that match the specified pattern associated
// with a wand. Use GetOption() to return the value of a particular option.
func (mw *MagickWand) GetOptions(pattern string) (options []string) {
	cspattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(cspattern))
	np := C.size_t(0)
	ps := C.MagickGetOptions(mw.mw, cspattern, &np)
	runtime.KeepAlive(mw)
	defer relinquishMemoryCStringArray(ps)
	options = sizedCStringArrayToStringSlice(ps, np)
	return
}

// Gets the wand orientation type.
func (mw *MagickWand) GetOrientation() OrientationType {
	ret := OrientationType(C.MagickGetOrientation(mw.mw))
	runtime.KeepAlive(mw)
	return ret
}

// Returns the page geometry associated with the magick wand.
func (mw *MagickWand) GetPage() (width, height uint, x, y int, err error) {
	var cw, ch C.size_t
	var cx, cy C.ssize_t
	ok := C.MagickGetPage(mw.mw, &cw, &ch, &cx, &cy)
	width, height, x, y = uint(cw), uint(ch), int(cx), int(cy)
	err = mw.getLastErrorIfFailed(ok)
	return
}

// Returns the font pointsize associated with the MagickWand.
func (mw *MagickWand) GetPointsize() float64 {
	ret := float64(C.MagickGetPointsize(mw.mw))
	runtime.KeepAlive(mw)
	return ret
}

// Gets the image X and Y resolution.
func (mw *MagickWand) GetResolution() (x, y float64, err error) {
	ok := C.MagickGetResolution(mw.mw, (*C.double)(&x), (*C.double)(&y))
	err = mw.getLastErrorIfFailed(ok)
	return
}

// Gets the horizontal and vertical sampling factor.
func (mw *MagickWand) GetSamplingFactors() (factors []float64) {
	num := C.size_t(0)
	pd := C.MagickGetSamplingFactors(mw.mw, &num)
	runtime.KeepAlive(mw)
	factors = sizedDoubleArrayToFloat64Slice(pd, num)
	return
}

// Returns the size associated with the magick wand.
func (mw *MagickWand) GetSize() (cols, rows uint, err error) {
	var cc, cr C.size_t
	ok := C.MagickGetSize(mw.mw, &cc, &cr)
	cols, rows, err = uint(cc), uint(cr), mw.getLastErrorIfFailed(ok)
	return
}

// Returns the size offset associated with the magick wand.
func (mw *MagickWand) GetSizeOffset() (offset int, err error) {
	var co C.ssize_t
	ok := C.MagickGetSizeOffset(mw.mw, &co)
	offset, err = int(co), mw.getLastErrorIfFailed(ok)
	return
}

// Returns the wand type.
func (mw *MagickWand) GetType() ImageType {
	ret := ImageType(C.MagickGetType(mw.mw))
	runtime.KeepAlive(mw)
	return ret
}

// Adds or removes a ICC, IPTC, or generic profile from an image. If the
// profile is empty, it is removed from the image otherwise added. Use a name
// of '*' and an empty profile to remove all profiles from the image.
//
// name: Name of profile to add or remove: ICC, IPTC, or generic profile.
//
func (mw *MagickWand) ProfileImage(name string, profile []byte) error {
	if len(profile) == 0 {
		return errors.New("zero-length profile not permitted")
	}
	csname := C.CString(name)
	defer C.free(unsafe.Pointer(csname))
	ok := C.MagickProfileImage(mw.mw, csname, unsafe.Pointer(&profile[0]), C.size_t(len(profile)))
	return mw.getLastErrorIfFailed(ok)
}

// Removes the named image profile and returns it.
//
// name: name of profile to return: ICC, IPTC, or generic profile.
//
func (mw *MagickWand) RemoveImageProfile(name string) []byte {
	csname := C.CString(name)
	defer C.free(unsafe.Pointer(csname))
	clen := C.size_t(0)
	profile := C.MagickRemoveImageProfile(mw.mw, csname, &clen)
	runtime.KeepAlive(mw)
	defer relinquishMemory(unsafe.Pointer(profile))
	return C.GoBytes(unsafe.Pointer(profile), C.int(clen))
}

// Sets the antialias propery of the wand.
func (mw *MagickWand) SetAntialias(antialias bool) error {
	ok := C.MagickSetAntialias(mw.mw, b2i(antialias))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the wand background color.
func (mw *MagickWand) SetBackgroundColor(background *PixelWand) error {
	ok := C.MagickSetBackgroundColor(mw.mw, background.pw)
	runtime.KeepAlive(background)
	return mw.getLastErrorIfFailed(ok)
}

// Sets the wand colorspace type.
func (mw *MagickWand) SetColorspace(colorspace ColorspaceType) error {
	ok := C.MagickSetColorspace(mw.mw, C.ColorspaceType(colorspace))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the wand compression type.
func (mw *MagickWand) SetCompression(compression CompressionType) error {
	ok := C.MagickSetCompression(mw.mw, C.CompressionType(compression))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the wand compression quality.
func (mw *MagickWand) SetCompressionQuality(quality uint) error {
	ok := C.MagickSetCompressionQuality(mw.mw, C.size_t(quality))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the wand pixel depth.
func (mw *MagickWand) SetDepth(depth uint) error {
	ok := C.MagickSetDepth(mw.mw, C.size_t(depth))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the extract geometry before you read or write an image file. Use it for
// inline cropping (e.g. 200x200+0+0) or resizing (e.g.200x200).
func (mw *MagickWand) SetExtract(geometry string) error {
	csgeometry := C.CString(geometry)
	defer C.free(unsafe.Pointer(csgeometry))
	ok := C.MagickSetExtract(mw.mw, csgeometry)
	return mw.getLastErrorIfFailed(ok)
}

// Sets the filename before you read or write an image file.
func (mw *MagickWand) SetFilename(filename string) error {
	csfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(csfilename))
	ok := C.MagickSetFilename(mw.mw, csfilename)
	return mw.getLastErrorIfFailed(ok)
}

// Sets the font associated with the MagickWand.
func (mw *MagickWand) SetFont(font string) error {
	csfont := C.CString(font)
	defer C.free(unsafe.Pointer(csfont))
	ok := C.MagickSetFont(mw.mw, csfont)
	return mw.getLastErrorIfFailed(ok)
}

// Sets the format of the magick wand.
func (mw *MagickWand) SetFormat(format string) error {
	csformat := C.CString(format)
	defer C.free(unsafe.Pointer(csformat))
	ok := C.MagickSetFormat(mw.mw, csformat)
	return mw.getLastErrorIfFailed(ok)
}

// Sets the gravity type.
func (mw *MagickWand) SetGravity(gtype GravityType) error {
	ok := C.MagickSetGravity(mw.mw, C.GravityType(gtype))
	return mw.getLastErrorIfFailed(ok)
}

// Associates a artifact with an image.
func (mw *MagickWand) SetImageArtifact(artifact, value string) error {
	csartifact := C.CString(artifact)
	defer C.free(unsafe.Pointer(csartifact))
	csvalue := C.CString(value)
	defer C.free(unsafe.Pointer(csvalue))
	ok := C.MagickSetImageArtifact(mw.mw, csartifact, csvalue)
	return mw.getLastErrorIfFailed(ok)
}

// Adds a named profile to the magick wand. If a profile with the same name
// already exists, it is replaced. This method differs from the ProfileImage()
// method in that it does not apply any CMS color profiles.
//
// name: Name of profile to add or remove: ICC, IPTC, or generic profile.
func (mw *MagickWand) SetImageProfile(name string, profile []byte) error {
	if len(profile) == 0 {
		return errors.New("zero-length profile not permitted")
	}
	csname := C.CString(name)
	defer C.free(unsafe.Pointer(csname))
	ok := C.MagickSetImageProfile(mw.mw, csname, unsafe.Pointer(&profile[0]), C.size_t(len(profile)))
	return mw.getLastErrorIfFailed(ok)
}

// Associates a property with an image.
func (mw *MagickWand) SetImageProperty(property, value string) error {
	csproperty := C.CString(property)
	defer C.free(unsafe.Pointer(csproperty))
	csvalue := C.CString(value)
	defer C.free(unsafe.Pointer(csvalue))
	ok := C.MagickSetImageProperty(mw.mw, csproperty, csvalue)
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image interlacing scheme
func (mw *MagickWand) SetInterlaceScheme(scheme InterlaceType) error {
	ok := C.MagickSetInterlaceScheme(mw.mw, C.InterlaceType(scheme))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the interpolate pixel method.
func (mw *MagickWand) SetInterpolateMethod(method PixelInterpolateMethod) error {
	ok := C.MagickSetInterpolateMethod(mw.mw, C.PixelInterpolateMethod(method))
	return mw.getLastErrorIfFailed(ok)
}

// Associates one or options with the wand (.e.g
// SetOption(wand, "jpeg:perserve", "yes")).
func (mw *MagickWand) SetOption(key, value string) error {
	cskey := C.CString(key)
	defer C.free(unsafe.Pointer(cskey))
	csvalue := C.CString(value)
	defer C.free(unsafe.Pointer(csvalue))
	ok := C.MagickSetOption(mw.mw, cskey, csvalue)
	return mw.getLastErrorIfFailed(ok)
}

// Sets the wand orientation type.
func (mw *MagickWand) SetOrientation(orientation OrientationType) error {
	ok := C.MagickSetOrientation(mw.mw, C.OrientationType(orientation))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the page geometry of the magick wand.
func (mw *MagickWand) SetPage(width, height uint, x, y int) error {
	ok := C.MagickSetPage(mw.mw, C.size_t(width), C.size_t(height), C.ssize_t(x), C.ssize_t(y))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the passphrase.
func (mw *MagickWand) SetPassphrase(passphrase string) error {
	cspassphrase := C.CString(passphrase)
	defer C.free(unsafe.Pointer(cspassphrase))
	ok := C.MagickSetPassphrase(mw.mw, cspassphrase)
	return mw.getLastErrorIfFailed(ok)
}

// Sets the font pointsize associated with the MagickWand.
func (mw *MagickWand) SetPointsize(pointSize float64) error {
	ok := C.MagickSetPointsize(mw.mw, C.double(pointSize))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the limit for a particular resource in megabytes.
func (mw *MagickWand) SetResourceLimit(rtype ResourceType, limit int64) error {
	ok := C.MagickSetResourceLimit(C.ResourceType(rtype), C.MagickSizeType(limit))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image resolution.
func (mw *MagickWand) SetResolution(xRes, yRes float64) error {
	ok := C.MagickSetResolution(mw.mw, C.double(xRes), C.double(yRes))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image sampling factors.
//
// samplingFactors: An array of floats representing the sampling factor for
// each color component (in RGB order).
func (mw *MagickWand) SetSamplingFactors(samplingFactors []float64) error {
	ok := C.MagickSetSamplingFactors(mw.mw, C.size_t(len(samplingFactors)), (*C.double)(&samplingFactors[0]))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the size of the magick wand. Set it before you read a raw image format
// such as RGB, GRAY, or CMYK.
func (mw *MagickWand) SetSize(cols, rows uint) error {
	ok := C.MagickSetSize(mw.mw, C.size_t(cols), C.size_t(rows))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the size and offset of the magick wand. Set it before you read a raw
// image format such as RGB, GRAY, or CMYK.
func (mw *MagickWand) SetSizeOffset(cols, rows uint, offset int) error {
	ok := C.MagickSetSizeOffset(mw.mw, C.size_t(cols), C.size_t(rows), C.ssize_t(offset))
	return mw.getLastErrorIfFailed(ok)
}

// Sets the image type attribute.
func (mw *MagickWand) SetType(itype ImageType) error {
	ok := C.MagickSetType(mw.mw, C.ImageType(itype))
	return mw.getLastErrorIfFailed(ok)
}
