package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

type PixelIterator struct {
	pi *C.PixelIterator
}

// Returns a new pixel iterator
//
// mw: the magick wand to iterate on
//
func NewPixelIterator(mw *MagickWand) *PixelIterator {
	return &PixelIterator{C.NewPixelIterator(mw.wand)}
}

// Returns a new pixel iterator
//
// mw: the magick wand to iterate on
// x, y, cols, rows: there values define the perimeter of a region of pixels
//
func NewPixelRegionIterator(mw *MagickWand, x, y int, width, height uint) *PixelIterator {
	return &PixelIterator{C.NewPixelRegionIterator(mw.wand, C.ssize_t(x), C.ssize_t(y), C.size_t(width), C.size_t(height))}
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
	pi.pi = C.DestroyPixelIterator(pi.pi)
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
func (pi *PixelIterator) GetCurrentIteratorRow() []*PixelWand {
	numWands := C.size_t
	pis := C.PixelGetCurrentIteratorRow(pi.pi, &numWands)
	var pixWands []*PixelWand
	q := uintptr(unsafe.Pointer(p))
	for i := 0; i < numWands; i++ {
		p = (**C.PixelWand)(unsafe.Pointer(q))
		if *p == nil {
			break
		}
		pixWands = append(pixWands, &PixelWand(*p))
		q += unsafe.Sizeof(q)
	}
	return pixWands
}

// Returns the current pixel iterator row.
func (pi *PixelIterator) GetIteratorRow() int {
	return int(C.PixelGetIteratorRow(pi.pi))
}

// Returns the next row as an array of pixel wands from the pixel iterator.
func (pi *PixelIterator) GetNextIteratorRow() []*PixelWand {
	numWands := C.size_t
	pis := C.PixelGetNextIteratorRow(pi.pi, &numWands)
	var pixWands []*PixelWand
	q := uintptr(unsafe.Pointer(p))
	for i := 0; i < numWands; i++ {
		p = (**C.PixelWand)(unsafe.Pointer(q))
		if *p == nil {
			break
		}
		pixWands = append(pixWands, &PixelWand(*p))
		q += unsafe.Sizeof(q)
	}
	return pixWands
}

// Returns the previous row as an array of pixel wands from the pixel iterator.
func (pi *PixelIterator) GetPreviousIteratorRow() []*PixelWand {
	numWands := C.size_t
	pis := C.PixelGetPreviousIteratorRow(pi.pi, &numWands)
	var pixWands []*PixelWand
	q := uintptr(unsafe.Pointer(p))
	for i := 0; i < numWands; i++ {
		p = (**C.PixelWand)(unsafe.Pointer(q))
		if *p == nil {
			break
		}
		pixWands = append(pixWands, &PixelWand(*p))
		q += unsafe.Sizeof(q)
	}
	return pixWands
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
	PixelSyncIterator(pi.pi)
	return pi.GetLastError()
}
