package imagick

/*
#include <MagickCore/MagickCore.h>
*/
import "C"
import "runtime"

// GetImageEntropy returns the entropy of one or more image channels.
func (mw *MagickWand) GetImageEntropy() (entropy float64, err error) {
	img := mw.GetImageFromMagickWand()

	var exc *C.ExceptionInfo = C.AcquireExceptionInfo()
	defer C.DestroyExceptionInfo(exc)

	ok := C.GetImageEntropy(img.img, (*C.double)(&entropy), exc)
	if C.int(ok) == 0 {
		return entropy, newExceptionInfo(exc)
	}

	runtime.KeepAlive(mw)
	runtime.KeepAlive(img)

	return entropy, nil
}
