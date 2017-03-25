// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

// AffineMatrix represents an ImageMagick AffineMatrix struct
type AffineMatrix struct {
	am C.AffineMatrix
}

// NewAffineMatrix constructs an AffineMatrix that is
// initialized to the indentify matrix
func NewAffineMatrix() *AffineMatrix {
	am := &AffineMatrix{}
	C.GetAffineMatrix(am.ptr())
	return am
}

// ptr returns the underlying C AffineMatrix pointer
func (a *AffineMatrix) ptr() *C.AffineMatrix {
	return &(a.am)
}

// ResetToIdentity resets the AffineMatrix to the
// identity matrix
func (a *AffineMatrix) ResetToIdentity() {
	C.GetAffineMatrix(a.ptr())
}

// TranslateX returns the TX value
func (a *AffineMatrix) TranslateX() float64 {
	return float64(a.am.tx)
}

// SetTranslateX sets the TX value
func (a *AffineMatrix) SetTranslateX(val float64) {
	a.am.tx = C.double(val)
}

// TranslateY returns the TY value
func (a *AffineMatrix) TranslateY() float64 {
	return float64(a.am.ty)
}

// SetTranslateY sets the TY value
func (a *AffineMatrix) SetTranslateY(val float64) {
	a.am.ty = C.double(val)
}

// RotateX returns the RX value
func (a *AffineMatrix) RotateX() float64 {
	return float64(a.am.rx)
}

// SetRotateX sets the RX value
func (a *AffineMatrix) SetRotateX(val float64) {
	a.am.rx = C.double(val)
}

// RotateY returns the RY value
func (a *AffineMatrix) RotateY() float64 {
	return float64(a.am.ry)
}

// SetRotateY sets the RY value
func (a *AffineMatrix) SetRotateY(val float64) {
	a.am.ry = C.double(val)
}

// ScaleX returns the SX value
func (a *AffineMatrix) ScaleX() float64 {
	return float64(a.am.sx)
}

// SetScaleX sets the SX value
func (a *AffineMatrix) SetScaleX(val float64) {
	a.am.sx = C.double(val)
}

// ScaleY returns the SY value
func (a *AffineMatrix) ScaleY() float64 {
	return float64(a.am.sy)
}

// SetScaleY sets the SY value
func (a *AffineMatrix) SetScaleY(val float64) {
	a.am.sy = C.double(val)
}
