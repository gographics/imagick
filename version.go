package magick

/*
#cgo pkg-config: MagickCore
#include <magick/MagickCore.h>
*/
import "C"

func GetCopyright() string {
	return C.GoString(C.GetMagickCopyright())
}

func GetFeatures() string {
	return C.GoString(C.GetMagickFeatures())
}

func GetHomeURL() string {
	return C.GoString(C.GetMagickHomeURL())
}

func GetPackageName() string {
	return C.GoString(C.GetMagickPackageName())
}

func GetQuantumDepth() (string, int) {
	var depth C.size_t
	cstr := C.GetMagickQuantumDepth(&depth)
	return C.GoString(cstr), int(depth)
}

func GetQuantumRange() (string, int) {
	var qrange C.size_t
	cstr := C.GetMagickQuantumRange(&qrange)
	return C.GoString(cstr), int(qrange)
}

func GetReleaseDate() string {
	return C.GoString(C.GetMagickReleaseDate())
}

func GetSignature() (uint, string) {
	panic("unsigned int GetMagickSignature(const StringInfo *nonce)")
}

func GetVersion() (string, int) {
	var version C.size_t
	cstr := C.GetMagickVersion(&version)
	return C.GoString(cstr), int(version)
}
