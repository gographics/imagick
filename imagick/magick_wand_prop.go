package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

import (
	//"fmt"
	"unsafe"
)

// This method deletes a wand artifact
func (mw *MagickWand) DeleteImageArtifact(artifact string) error {
	csartifact := C.CString(artifact)
	defer C.free(unsafe.Pointer(csartifact))
	C.MagickDeleteImageArtifact(mw.mw, csartifact)
	return mw.GetLastError()
}

// This method deletes a image property
func (mw *MagickWand) DeleteImageProperty(property string) error {
	csproperty := C.CString(property)
	defer C.free(unsafe.Pointer(csproperty))
	C.MagickDeleteImageProperty(mw.mw, csproperty)
	return mw.GetLastError()
}

// This method deletes a wand option
func (mw *MagickWand) DeleteOption(option string) error {
	csoption := C.CString(option)
	defer C.free(unsafe.Pointer(csoption))
	C.MagickDeleteOption(mw.mw, csoption)
	return mw.GetLastError()
}

// Returns the antialias property associated with the wand
func (mw *MagickWand) GetAntialias() bool {
	return 1 == C.int(C.MagickGetAntialias(mw.mw))
}

// Returns the wand background color
func (mw *MagickWand) GetBackgroundColor() *PixelWand {
	return NewPixelWandFromCAPI(C.MagickGetBackgroundColor(mw.mw))
}

// Returns the wand colorspace type
func (mw *MagickWand) GetColorspace() ColorspaceType {
	ccst := C.MagickGetColorspace(mw.mw)
	return ColorspaceType(ccst)
}
