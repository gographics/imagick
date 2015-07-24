// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <wand/MagickWand.h>

static PixelWand* get_pw_at(PixelWand** pws, size_t pos) {
	return pws[pos];
}
*/
import "C"
import "unsafe"

type PixelIterator struct {
	pi *C.PixelIterator
}

// Returns a new pixel iterator
//
// mw: the magick wand to iterate on
//
func (mw *MagickWand) NewPixelIterator() *PixelIterator {
	return &PixelIterator{C.NewPixelIterator(mw.mw)}
}

// Returns a new pixel iterator
//
// mw: the magick wand to iterate on
// x, y, cols, rows: there values define the perimeter of a region of pixels
//
func (mw *MagickWand) NewPixelRegionIterator(x, y int, width, height uint) *PixelIterator {
	return &PixelIterator{C.NewPixelRegionIterator(mw.mw, C.ssize_t(x), C.ssize_t(y), C.size_t(width), C.size_t(height))}
}

// Clear resources associated with a PixelIterator.
func (pi *PixelIterator) Clear() {
	C.ClearPixelIterator(pi.pi)
}

// Makes an exact copy of the specified iterator.
func (pi *PixelIterator) Clone() *PixelIterator {
	return &PixelIterator{C.ClonePixelIterator(pi.pi)}
}

// Deallocates resources associated with a PixelIterator.
func (pi *PixelIterator) Destroy() {
	if pi.pi == nil {
		return
	}
	pi.pi = C.DestroyPixelIterator(pi.pi)
	relinquishMemory(unsafe.Pointer(pi.pi))
	pi.pi = nil
}

// Returns true if the iterator is verified as a pixel iterator.
func (pi *PixelIterator) IsVerified() bool {
	if pi.pi == nil {
		return false
	}
	return 1 == C.IsPixelIterator(pi.pi)
}

// Returns the current row as an array of pixel wands from the pixel iterator.
func (pi *PixelIterator) GetCurrentIteratorRow() (pws []*PixelWand) {
	num := C.size_t(0)
	pp := C.PixelGetCurrentIteratorRow(pi.pi, &num)
	if pp == nil {
		return
	}
	for i := 0; i < int(num); i++ {
		cpw := C.get_pw_at(pp, C.size_t(i))
		pws = append(pws, &PixelWand{cpw})
	}
	return
}

// Returns the current pixel iterator row.
func (pi *PixelIterator) GetIteratorRow() int {
	return int(C.PixelGetIteratorRow(pi.pi))
}

// Returns the next row as an array of pixel wands from the pixel iterator.
func (pi *PixelIterator) GetNextIteratorRow() (pws []*PixelWand) {
	num := C.size_t(0)
	pp := C.PixelGetNextIteratorRow(pi.pi, &num)
	if pp == nil {
		return
	}
	for i := 0; i < int(num); i++ {
		cpw := C.get_pw_at(pp, C.size_t(i))
		pws = append(pws, &PixelWand{cpw})
	}
	return
}

// Returns the previous row as an array of pixel wands from the pixel iterator.
func (pi *PixelIterator) GetPreviousIteratorRow() (pws []*PixelWand) {
	num := C.size_t(0)
	pp := C.PixelGetPreviousIteratorRow(pi.pi, &num)
	if pp == nil {
		return
	}
	for i := 0; i < int(num); i++ {
		cpw := C.get_pw_at(pp, C.size_t(i))
		pws = append(pws, &PixelWand{cpw})
	}
	return
}

// Resets the pixel iterator. Use it in conjunction with GetNextIteratorRow()
// to iterate over all the pixels in a pixel container.
func (pi *PixelIterator) Reset() {
	C.PixelResetIterator(pi.pi)
}

// Sets the pixel iterator to the first pixel row.
func (pi *PixelIterator) SetFirstIteratorRow() {
	C.PixelSetFirstIteratorRow(pi.pi)
}

// Set the pixel iterator row.
func (pi *PixelIterator) SetIteratorRow(row int) error {
	C.PixelSetIteratorRow(pi.pi, C.ssize_t(row))
	return pi.GetLastError()
}

// Sets the pixel iterator to the last pixel row.
func (pi *PixelIterator) SetLastIteratorRow() {
	C.PixelSetLastIteratorRow(pi.pi)
}

// Syncs the pixel iterator.
func (pi *PixelIterator) SyncIterator() error {
	C.PixelSyncIterator(pi.pi)
	return pi.GetLastError()
}
