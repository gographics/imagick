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

func ConvertRGBToHSB(qr, qg, qb Quantum) (h, s, b float64) {
	var dh, ds, db C.double
	C.ConvertRGBToHSB(C.Quantum(qr), C.Quantum(qg), C.Quantum(qb), &dh, &ds, &db)
	return float64(dh), float64(ds), float64(db)
}

func ConvertHSBToRGB(fh, fs, fb float64) (r, g, b Quantum) {
	var qr, qg, qb C.Quantum
	C.ConvertHSBToRGB(C.double(fh), C.double(fs), C.double(fb), &qr, &qg, &qb)
	return Quantum(qr), Quantum(qg), Quantum(qb)
}

func ConvertRGBToHSL(qr, qg, qb Quantum) (h, s, l float64) {
	var dh, ds, dl C.double
	C.ConvertRGBToHSL(C.Quantum(qr), C.Quantum(qg), C.Quantum(qb), &dh, &ds, &dl)
	return float64(dh), float64(ds), float64(dl)
}

func ConvertHSLToRGB(fh, fs, fl float64) (r, g, b Quantum) {
	var qr, qg, qb C.Quantum
	C.ConvertHSLToRGB(C.double(fh), C.double(fs), C.double(fl), &qr, &qg, &qb)
	return Quantum(qr), Quantum(qg), Quantum(qb)
}
