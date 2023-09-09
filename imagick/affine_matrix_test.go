// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

import "testing"

func TestAffineMatrix(t *testing.T) {
	m := NewAffineMatrix()
	assertAffineMatrixIdentity(m, t)

	mw := NewMagickWand()
	dw := NewDrawingWand()

	mw.SetSize(300, 100)
	mw.ReadImage("xc:white")

	dw.SetFontSize(72)
	dw.SetFont("Times-New-Roman")

	m.SetTranslateX(2)
	m.SetTranslateY(4)
	m.SetScaleX(-1)
	m.SetScaleY(-1)

	if m.TranslateX() != 2 {
		t.Errorf("Expected TX=2, got %v", m.TranslateX())
	}
	if m.TranslateY() != 4 {
		t.Errorf("Expected TY=4, got %v", m.TranslateY())
	}
	if m.ScaleX() != -1 {
		t.Errorf("Expected SX=-1, got %v", m.ScaleX())
	}
	if m.ScaleY() != -1 {
		t.Errorf("Expected SY=-1, got %v", m.ScaleY())
	}

	dw.Affine(m)
	mw.DrawImage(dw)

	m.ResetToIdentity()
	assertAffineMatrixIdentity(m, t)
}

func assertAffineMatrixIdentity(m *AffineMatrix, t *testing.T) {
	if m.TranslateX() != 0 {
		t.Errorf("Expected TX=0, got %v", m.TranslateX())
	}
	if m.TranslateY() != 0 {
		t.Errorf("Expected TY=0, got %v", m.TranslateY())
	}
	if m.RotateX() != 0 {
		t.Errorf("Expected RX=0, got %v", m.RotateX())
	}
	if m.RotateY() != 0 {
		t.Errorf("Expected RY=0, got %v", m.RotateY())
	}
	if m.ScaleX() != 1 {
		t.Errorf("Expected SX=1, got %v", m.ScaleX())
	}
	if m.ScaleY() != 1 {
		t.Errorf("Expected SY=1, got %v", m.ScaleY())
	}
}
