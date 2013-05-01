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

func ptrCStringArrayToStringSlice(p **C.char) []string {
	var strings []string
	q := uintptr(unsafe.Pointer(p))
	for {
		p = (**C.char)(unsafe.Pointer(q))
		if *p == nil {
			break
		}
		strings = append(strings, C.GoString(*p))
		q += unsafe.Sizeof(q)
	}
	return strings
}
