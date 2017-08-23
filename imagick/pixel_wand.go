// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

import (
	"runtime"
	"sync"
	"sync/atomic"
	"unsafe"
)

type PixelWand struct {
	pw   *C.PixelWand
	init sync.Once
}

// Returns a new pixel wand
func newPixelWand(cpw *C.PixelWand) *PixelWand {
	pw := &PixelWand{pw: cpw}
	runtime.SetFinalizer(pw, Destroy)
	pw.IncreaseCount()

	return pw
}

// Returns a new pixel wand
func NewPixelWand() *PixelWand {
	return newPixelWand(C.NewPixelWand())
}

// Clears resources associated with the wand
func (pw *PixelWand) Clear() {
	C.ClearPixelWand(pw.pw)
	runtime.KeepAlive(pw)
}

// Makes an exact copy of the wand
func (pw *PixelWand) Clone() *PixelWand {
	ret := newPixelWand(C.ClonePixelWand(pw.pw))
	runtime.KeepAlive(pw)
	return ret
}

// Deallocates resources associated with a pixel wand
func (pw *PixelWand) Destroy() {
	if pw.pw == nil {
		return
	}

	pw.init.Do(func() {
		pw.pw = C.DestroyPixelWand(pw.pw)
		relinquishMemory(unsafe.Pointer(pw.pw))
		pw.pw = nil

		pw.DecreaseCount()
	})
}

// Returns true if the distance between two colors is less than the specified distance
func (pw *PixelWand) IsSimilar(pixelWand *PixelWand, fuzz float64) bool {
	ret := 1 == C.int(C.IsPixelWandSimilar(pw.pw, pixelWand.pw, C.double(fuzz)))
	runtime.KeepAlive(pw)
	runtime.KeepAlive(pixelWand)
	return ret
}

// Returns true if the wand is verified as a pixel wand
func (pw *PixelWand) IsVerified() bool {
	if pw.pw != nil {
		return 1 == C.int(C.IsPixelWand(pw.pw))
	}
	runtime.KeepAlive(pw)
	return false
}

// Increase PixelWand ref counter and set according "can`t be terminated status"
func (pw *PixelWand) IncreaseCount() {
	atomic.AddInt64(&pixelWandCounter, int64(1))
	unsetCanTerminate()
}

// Decrease PixelWand ref counter and set according "can be terminated status"
func (pw *PixelWand) DecreaseCount() {
	atomic.AddInt64(&pixelWandCounter, int64(-1))
	setCanTerminate()
}

// Returns the normalized alpha color of the pixel wand
func (pw *PixelWand) GetAlpha() float64 {
	ret := float64(C.PixelGetAlpha(pw.pw))
	runtime.KeepAlive(pw)
	return ret
}

// Returns the alpha value of the pixel wand
func (pw *PixelWand) GetAlphaQuantum() Quantum {
	ret := Quantum(C.PixelGetAlphaQuantum(pw.pw))
	runtime.KeepAlive(pw)
	return ret
}

// Returns the normalized black color of the pixel wand
func (pw *PixelWand) GetBlack() float64 {
	ret := float64(C.PixelGetBlack(pw.pw))
	runtime.KeepAlive(pw)
	return ret
}

// Returns the black color of the pixel wand
func (pw *PixelWand) GetBlackQuantum() Quantum {
	ret := Quantum(C.PixelGetBlackQuantum(pw.pw))
	runtime.KeepAlive(pw)
	return ret
}

// Returns the normalized blue color of the pixel wand
func (pw *PixelWand) GetBlue() float64 {
	ret := float64(C.PixelGetBlue(pw.pw))
	runtime.KeepAlive(pw)
	return ret
}

// Returns the blue color of the pixel wand
func (pw *PixelWand) GetBlueQuantum() Quantum {
	ret := Quantum(C.PixelGetBlueQuantum(pw.pw))
	runtime.KeepAlive(pw)
	return ret
}

// Returns the color of the pixel wand as a string
func (pw *PixelWand) GetColorAsString() string {
	p := C.PixelGetColorAsString(pw.pw)
	runtime.KeepAlive(pw)
	defer relinquishMemory(unsafe.Pointer(p))
	return C.GoString(p)
}

// Returns the normalized color of the pixel wand as string
func (pw *PixelWand) GetColorAsNormalizedString() string {
	p := C.PixelGetColorAsNormalizedString(pw.pw)
	runtime.KeepAlive(pw)
	defer relinquishMemory(unsafe.Pointer(p))
	return C.GoString(p)
}

// Returns the color count associated with this color
func (pw *PixelWand) GetColorCount() uint {
	ret := uint(C.PixelGetColorCount(pw.pw))
	runtime.KeepAlive(pw)
	return ret
}

// Returns the normalized cyan color of the pixel wand
func (pw *PixelWand) GetCyan() float64 {
	ret := float64(C.PixelGetCyan(pw.pw))
	runtime.KeepAlive(pw)
	return ret
}

// Returns the cyan color of the pixel wand
func (pw *PixelWand) GetCyanQuantum() Quantum {
	ret := Quantum(C.PixelGetCyanQuantum(pw.pw))
	runtime.KeepAlive(pw)
	return ret
}

// Returns the normalized fuzz value of the pixel wand
func (pw *PixelWand) GetFuzz() float64 {
	ret := float64(C.PixelGetFuzz(pw.pw))
	runtime.KeepAlive(pw)
	return ret
}

// Returns the normalized green color of the pixel wand
func (pw *PixelWand) GetGreen() float64 {
	ret := float64(C.PixelGetGreen(pw.pw))
	runtime.KeepAlive(pw)
	return ret
}

// Returns the green color of the pixel wand
func (pw *PixelWand) GetGreenQuantum() Quantum {
	ret := Quantum(C.PixelGetGreenQuantum(pw.pw))
	runtime.KeepAlive(pw)
	return ret
}

// Returns the normalized HSL color of the pixel wand
func (pw *PixelWand) GetHSL() (hue, saturation, brightness float64) {
	var cdhue, cdsaturation, cdbrightness C.double
	C.PixelGetHSL(pw.pw, &cdhue, &cdsaturation, &cdbrightness)
	runtime.KeepAlive(pw)
	hue, saturation, brightness = float64(cdhue), float64(cdsaturation), float64(cdbrightness)
	return
}

// Returns the colormap index from the pixel wand
func (pw *PixelWand) GetIndex() IndexPacket {
	cip := C.PixelGetIndex(pw.pw)
	runtime.KeepAlive(pw)
	return IndexPacket(cip)
}

// Returns the normalized magenta color of the pixel wand
func (pw *PixelWand) GetMagenta() float64 {
	ret := float64(C.PixelGetMagenta(pw.pw))
	runtime.KeepAlive(pw)
	return ret
}

// Returns the magenta color of the pixel wand
func (pw *PixelWand) GetMagentaQuantum() Quantum {
	ret := Quantum(C.PixelGetMagentaQuantum(pw.pw))
	runtime.KeepAlive(pw)
	return ret
}

// Gets the magick color of the pixel wand
func (pw *PixelWand) GetMagickColor() *PixelInfo {
	var pi C.PixelInfo
	C.PixelGetMagickColor(pw.pw, &pi)
	runtime.KeepAlive(pw)
	return newPixelInfoFromCAPI(&pi)
}

// Deprecated: Use GetAlpha()
func (pw *PixelWand) GetOpacity() float64 {
	return pw.GetAlpha()
}

// Deprecated: Use GetAlphaQuantum()
func (pw *PixelWand) GetOpacityQuantum() Quantum {
	return pw.GetAlphaQuantum()
}

// Gets the color of the pixel wand as a PixelPacket
func (pw *PixelWand) GetQuantumColor() *PixelInfo {
	var pi C.PixelInfo
	C.PixelGetQuantumPacket(pw.pw, &pi)
	runtime.KeepAlive(pw)
	return newPixelInfoFromCAPI(&pi)
}

// Returns the normalized red color of the pixel wand
func (pw *PixelWand) GetRed() float64 {
	ret := float64(C.PixelGetRed(pw.pw))
	runtime.KeepAlive(pw)
	return ret
}

// Returns the red color of the pixel wand
func (pw *PixelWand) GetRedQuantum() Quantum {
	ret := Quantum(C.PixelGetRedQuantum(pw.pw))
	runtime.KeepAlive(pw)
	return ret
}

// Returns the normalized yellow color of the pixel wand
func (pw *PixelWand) GetYellow() float64 {
	ret := float64(C.PixelGetYellow(pw.pw))
	runtime.KeepAlive(pw)
	return ret
}

// Returns the yellow color of the pixel wand
func (pw *PixelWand) GetYellowQuantum() Quantum {
	ret := Quantum(C.PixelGetYellowQuantum(pw.pw))
	runtime.KeepAlive(pw)
	return ret
}

// Sets the normalized alpha color of the pixel wand.
// 1.0 is fully opaque and 0.0 is fully transparent.
func (pw *PixelWand) SetAlpha(alpha float64) {
	C.PixelSetAlpha(pw.pw, C.double(alpha))
	runtime.KeepAlive(pw)
}

// Sets the alpha color of the pixel wand
func (pw *PixelWand) SetAlphaQuantum(opacity Quantum) {
	C.PixelSetAlphaQuantum(pw.pw, C.Quantum(opacity))
	runtime.KeepAlive(pw)
}

// Sets the normalized black color of the pixel wand
func (pw *PixelWand) SetBlack(black float64) {
	C.PixelSetBlack(pw.pw, C.double(black))
	runtime.KeepAlive(pw)
}

// Sets the black color of the pixel wand
func (pw *PixelWand) SetBlackQuantum(black Quantum) {
	C.PixelSetBlackQuantum(pw.pw, C.Quantum(black))
	runtime.KeepAlive(pw)
}

// Sets the normalized blue color of the pixel wand
func (pw *PixelWand) SetBlue(blue float64) {
	C.PixelSetBlue(pw.pw, C.double(blue))
	runtime.KeepAlive(pw)
}

// Sets the blue color of the pixel wand
func (pw *PixelWand) SetBlueQuantum(blue Quantum) {
	C.PixelSetBlueQuantum(pw.pw, C.Quantum(blue))
	runtime.KeepAlive(pw)
}

// Sets the color of the pixel wand with a string (e.g. "blue", "#0000ff", "rgb(0,0,255)", "cmyk(100,100,100,10)", etc.)
func (pw *PixelWand) SetColor(color string) bool {
	cscolor := C.CString(color)
	defer C.free(unsafe.Pointer(cscolor))
	ret := 1 == int(C.PixelSetColor(pw.pw, cscolor))
	runtime.KeepAlive(pw)
	return ret
}

// Sets the color count of the pixel wand
func (pw *PixelWand) SetColorCount(count uint) {
	C.PixelSetColorCount(pw.pw, C.size_t(count))
	runtime.KeepAlive(pw)
}

// Sets the color of the pixel wand from another one
func (pw *PixelWand) SetColorFromWand(pixelWand *PixelWand) {
	C.PixelSetColorFromWand(pw.pw, pixelWand.pw)
	runtime.KeepAlive(pw)
}

// Sets the normalized cyan color of the pixel wand
func (pw *PixelWand) SetCyan(cyan float64) {
	C.PixelSetCyan(pw.pw, C.double(cyan))
	runtime.KeepAlive(pw)
}

// Sets the cyan color of the pixel wand
func (pw *PixelWand) SetCyanQuantum(cyan Quantum) {
	C.PixelSetCyanQuantum(pw.pw, C.Quantum(cyan))
	runtime.KeepAlive(pw)
}

// Sets the fuzz value of the pixel wand
func (pw *PixelWand) SetFuzz(fuzz float64) {
	C.PixelSetFuzz(pw.pw, C.double(fuzz))
	runtime.KeepAlive(pw)
}

// Sets the normalized green color of the pixel wand
func (pw *PixelWand) SetGreen(green float64) {
	C.PixelSetGreen(pw.pw, C.double(green))
	runtime.KeepAlive(pw)
}

// Sets the green color of the pixel wand
func (pw *PixelWand) SetGreenQuantum(green Quantum) {
	C.PixelSetGreenQuantum(pw.pw, C.Quantum(green))
	runtime.KeepAlive(pw)
}

// Sets the normalized HSL color of the pixel wand
func (pw *PixelWand) SetHSL(hue, saturation, brightness float64) {
	C.PixelSetHSL(pw.pw, C.double(hue), C.double(saturation), C.double(brightness))
	runtime.KeepAlive(pw)
}

// Sets the colormap index of the pixel wand
func (pw *PixelWand) SetIndex(quantum Quantum) {
	C.PixelSetIndex(pw.pw, C.Quantum(quantum))
	runtime.KeepAlive(pw)
}

// Sets the normalized magenta color of the pixel wand
func (pw *PixelWand) SetMagenta(magenta float64) {
	C.PixelSetMagenta(pw.pw, C.double(magenta))
	runtime.KeepAlive(pw)
}

// Sets the magenta color of the pixel wand
func (pw *PixelWand) SetMagentaQuantum(magenta Quantum) {
	C.PixelSetMagentaQuantum(pw.pw, C.Quantum(magenta))
	runtime.KeepAlive(pw)
}

// Deprecated: Use SetAlpha()
func (pw *PixelWand) SetOpacity(opacity float64) {
	pw.SetAlpha(opacity)
}

// Deprecared: USe SetAlphaQuantum()
func (pw *PixelWand) SetOpacityQuantum(opacity Quantum) {
	pw.SetAlphaQuantum(opacity)
}

// Deprecared: Use SetPixelColor
func (pw *PixelWand) SetMagickColor(color *PixelInfo) {
	C.PixelSetPixelColor(pw.pw, color.pi)
	runtime.KeepAlive(pw)
}

// Sets the color of the pixel wand
func (pw *PixelWand) SetPixelColor(color *PixelInfo) {
	C.PixelSetPixelColor(pw.pw, color.pi)
	runtime.KeepAlive(pw)
}

// Sets the normalized red color of the pixel wand
func (pw *PixelWand) SetRed(red float64) {
	C.PixelSetRed(pw.pw, C.double(red))
	runtime.KeepAlive(pw)
}

// Sets the red color of the pixel wand
func (pw *PixelWand) SetRedQuantum(red Quantum) {
	C.PixelSetRedQuantum(pw.pw, C.Quantum(red))
	runtime.KeepAlive(pw)
}

// Sets the normalized yellow color of the pixel wand
func (pw *PixelWand) SetYellow(yellow float64) {
	C.PixelSetYellow(pw.pw, C.double(yellow))
	runtime.KeepAlive(pw)
}

// Sets the yellow color of the pixel wand
func (pw *PixelWand) SetYellowQuantum(yellow Quantum) {
	C.PixelSetYellowQuantum(pw.pw, C.Quantum(yellow))
	runtime.KeepAlive(pw)
}
