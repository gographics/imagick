// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <magick/MagickCore.h>
*/
import "C"
import "unsafe"

// MagickPixelPacket is constructed using NewMagickPixelPacket
type MagickPixelPacket struct {
	mpp *C.MagickPixelPacket
}

func (p *MagickPixelPacket) Red() float64 {
	return float64(p.mpp.red)
}

func (p *MagickPixelPacket) SetRed(value float64) {
	p.mpp.red = C.MagickRealType(value)
}

func (p *MagickPixelPacket) Green() float64 {
	return float64(p.mpp.green)
}

func (p *MagickPixelPacket) SetGreen(value float64) {
	p.mpp.green = C.MagickRealType(value)
}

func (p *MagickPixelPacket) Blue() float64 {
	return float64(p.mpp.blue)
}

func (p *MagickPixelPacket) SetBlue(value float64) {
	p.mpp.blue = C.MagickRealType(value)
}

func (p *MagickPixelPacket) Opacity() float64 {
	return float64(p.mpp.opacity)
}

func (p *MagickPixelPacket) SetOpacity(value float64) {
	p.mpp.opacity = C.MagickRealType(value)
}

func (p *MagickPixelPacket) Index() float64 {
	return float64(p.mpp.index)
}

func (p *MagickPixelPacket) SetIndex(value float64) {
	p.mpp.index = C.MagickRealType(value)
}

func NewMagickPixelPacket() *MagickPixelPacket {
	var mpp C.MagickPixelPacket
	return &MagickPixelPacket{mpp: &mpp}
}

func newMagickPixelPacketFromCAPI(mpp *C.MagickPixelPacket) *MagickPixelPacket {
	return &MagickPixelPacket{mpp}
}

// QueryMagickColor returns the red, green, blue, and opacity
// intensities for a given color name.
func QueryMagickColor(colorName string) (*MagickPixelPacket, error) {
	cstr := C.CString(colorName)
	defer C.free(unsafe.Pointer(cstr))

	var mpp C.MagickPixelPacket

	var exc *C.ExceptionInfo = C.AcquireExceptionInfo()
	defer C.DestroyExceptionInfo(exc)

	ok := C.QueryMagickColor(cstr, &mpp, exc)
	if C.int(ok) == 0 {
		return nil, newExceptionInfo(exc)
	}

	return newMagickPixelPacketFromCAPI(&mpp), nil
}
