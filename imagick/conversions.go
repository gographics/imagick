package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"
import "unsafe"

// Convert a boolean to a integer
func b2i(boolean bool) C.MagickBooleanType {
	if boolean {
		return C.MagickBooleanType(1)
	}
	return C.MagickBooleanType(0)
}

func cStringArrayToStringSlice(p **C.char) []string {
	var strings []string
	q := uintptr(unsafe.Pointer(p))
	for {
		p = (**C.char)(unsafe.Pointer(q))
		if *p == nil {
			break
		}
		defer C.free(unsafe.Pointer(*p))
		strings = append(strings, C.GoString(*p))
		q += unsafe.Sizeof(q)
	}
	return strings
}

func sizedCStringArrayToStringSlice(p **C.char, num C.size_t) []string {
	var strings []string
	q := uintptr(unsafe.Pointer(p))
	for i := 0; i < int(num); i++ {
		p = (**C.char)(unsafe.Pointer(q))
		if *p == nil {
			break
		}
		defer C.free(unsafe.Pointer(*p))
		strings = append(strings, C.GoString(*p))
		q += unsafe.Sizeof(q)
	}
	return strings
}

func sizedDoubleArrayToFloat64Slice(p *C.double, num C.size_t) []float64 {
	var nums []float64
	q := uintptr(unsafe.Pointer(p))
	for i := 0; i < int(num); i++ {
		p = (*C.double)(unsafe.Pointer(q))
		nums = append(nums, float64(*p))
		q += unsafe.Sizeof(q)
	}
	return nums
}
