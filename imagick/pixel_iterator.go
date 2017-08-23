// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>

static PixelWand* get_pw_at(PixelWand** pws, size_t pos) {
	return pws[pos];
}
*/
import "C"
import (
	"runtime"
	"sync"
	"sync/atomic"
	"unsafe"
)

type PixelIterator struct {
	pi   *C.PixelIterator
	init sync.Once
}

func newPixelIterator(cpi *C.PixelIterator) *PixelIterator {
	pi := &PixelIterator{pi: cpi}
	runtime.SetFinalizer(pi, Destroy)
	pi.IncreaseCount()

	return pi
}

// Returns a new pixel iterator
//
// mw: the magick wand to iterate on
//
func (mw *MagickWand) NewPixelIterator() *PixelIterator {
	ret := newPixelIterator(C.NewPixelIterator(mw.mw))
	runtime.KeepAlive(mw)
	return ret
}

// Returns a new pixel iterator
//
// mw: the magick wand to iterate on
// x, y, cols, rows: there values define the perimeter of a region of pixels
//
func (mw *MagickWand) NewPixelRegionIterator(x, y int, width, height uint) *PixelIterator {
	ret := newPixelIterator(C.NewPixelRegionIterator(mw.mw, C.ssize_t(x), C.ssize_t(y), C.size_t(width), C.size_t(height)))
	runtime.KeepAlive(mw)
	return ret
}

// Clear resources associated with a PixelIterator.
func (pi *PixelIterator) Clear() {
	C.ClearPixelIterator(pi.pi)
	runtime.KeepAlive(pi)
}

// Makes an exact copy of the specified iterator.
func (pi *PixelIterator) Clone() *PixelIterator {
	ret := newPixelIterator(C.ClonePixelIterator(pi.pi))
	runtime.KeepAlive(pi)
	return ret
}

// Deallocates resources associated with a PixelIterator.
func (pi *PixelIterator) Destroy() {
	if pi.pi == nil {
		return
	}

	pi.init.Do(func() {
		pi.pi = C.DestroyPixelIterator(pi.pi)
		relinquishMemory(unsafe.Pointer(pi.pi))
		pi.pi = nil

		pi.DecreaseCount()
	})
}

// Returns true if the iterator is verified as a pixel iterator.
func (pi *PixelIterator) IsVerified() bool {
	if pi.pi == nil {
		return false
	}
	ret := 1 == C.IsPixelIterator(pi.pi)
	runtime.KeepAlive(pi)
	return ret
}

// Increase PixelIterator ref counter and set according "can`t be terminated status"
func (pi *PixelIterator) IncreaseCount() {
	atomic.AddInt64(&pixelIteratorCounter, int64(1))
	unsetCanTerminate()
}

// Decrease DrawingWand ref counter and set according "can be terminated status"
func (pi *PixelIterator) DecreaseCount() {
	atomic.AddInt64(&pixelIteratorCounter, int64(-1))
	setCanTerminate()
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
		// PixelWand is a private pointer, borrowed from the pixel
		// iterator. We don't want to reference count it. It will
		// get destroyed by C API when PixelIterator is destroyed.
		pws = append(pws, &PixelWand{pw: cpw})
	}
	runtime.KeepAlive(pi)
	return
}

// Returns the current pixel iterator row.
func (pi *PixelIterator) GetIteratorRow() int {
	ret := int(C.PixelGetIteratorRow(pi.pi))
	runtime.KeepAlive(pi)
	return ret
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
		// PixelWand is a private pointer, borrowed from the pixel
		// iterator. We don't want to reference count it. It will
		// get destroyed by C API when PixelIterator is destroyed.
		pws = append(pws, &PixelWand{pw: cpw})
	}
	runtime.KeepAlive(pi)
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
		// PixelWand is a private pointer, borrowed from the pixel
		// iterator. We don't want to reference count it. It will
		// get destroyed by C API when PixelIterator is destroyed.
		pws = append(pws, &PixelWand{pw: cpw})
	}
	runtime.KeepAlive(pi)
	return
}

// Resets the pixel iterator. Use it in conjunction with GetNextIteratorRow()
// to iterate over all the pixels in a pixel container.
func (pi *PixelIterator) Reset() {
	C.PixelResetIterator(pi.pi)
	runtime.KeepAlive(pi)
}

// Sets the pixel iterator to the first pixel row.
func (pi *PixelIterator) SetFirstIteratorRow() {
	C.PixelSetFirstIteratorRow(pi.pi)
	runtime.KeepAlive(pi)
}

// Set the pixel iterator row.
func (pi *PixelIterator) SetIteratorRow(row int) error {
	ok := C.PixelSetIteratorRow(pi.pi, C.ssize_t(row))
	return pi.getLastErrorIfFailed(ok)
}

// Sets the pixel iterator to the last pixel row.
func (pi *PixelIterator) SetLastIteratorRow() {
	C.PixelSetLastIteratorRow(pi.pi)
	runtime.KeepAlive(pi)
}

// Syncs the pixel iterator.
func (pi *PixelIterator) SyncIterator() error {
	ok := C.PixelSyncIterator(pi.pi)
	return pi.getLastErrorIfFailed(ok)
}
