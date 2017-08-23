// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickCore/MagickCore.h>
*/
import "C"

import (
	"runtime"
	"unsafe"
)

// This struct represents the KernelInfo C API of ImageMagick
type KernelInfo struct {
	info *C.KernelInfo
}

func newKernelInfo(cki *C.KernelInfo) *KernelInfo {
	ki := &KernelInfo{info: cki}
	runtime.SetFinalizer(ki, Destroy)
	return ki
}

// Destroy the KernelInfo immediately.
// This will also be called automatically during garbage collection.
func (ki *KernelInfo) Destroy() {
	if ki.info != nil {
		C.DestroyKernelInfo(ki.info)
		ki.info = nil
	}
}

// Convert the current KernelInfo to an 2d-array of values. The values are either
// a float64 if the element is used, or NaN if the element is not used by the kernel
func (ki *KernelInfo) ToArray() [][]float64 {
	var values [][]float64

	count := 0
	cValues := (*[1 << 10]C.double)(unsafe.Pointer(ki.info.values))

	for y := C.size_t(0); y < ki.info.height; y++ {
		rowValues := make([]float64, ki.info.width)
		for x := C.size_t(0); x < ki.info.width; x++ {
			rowValues[x] = float64(cValues[count])
			count++
		}
		values = append(values, rowValues)
	}

	runtime.KeepAlive(ki)
	return values
}

/*
NewKernelInfo takes the given string (generally supplied by the user) and converts it
into a Morphology/Convolution Kernel. This allows users to specify a kernel from a number
of pre-defined kernels, or to fully specify their own kernel for a specific Convolution
or Morphology Operation.

Returns a *ExceptionInfo on error

Ref: http://www.imagemagick.org/api/morphology.php#AcquireKernelBuiltIn

Details:

The kernel so generated can be any rectangular array of floating point values (doubles)
with the 'control point' or 'pixel being affected' anywhere within that array of values.

Previously IM was restricted to a square of odd size using the exact center as origin,
this is no longer the case, and any rectangular kernel with any value being declared the origin.
This in turn allows the use of highly asymmetrical kernels.

The floating point values in the kernel can also include a special value known as 'nan'
or 'not a number' to indicate that this value is not part of the kernel array. This allows
you to shaped the kernel within its rectangular area. That is 'nan' values provide a 'mask'
for the kernel shape. However at least one non-nan value must be provided for correct working
of a kernel.

Input kernel defintion strings can consist of any of three types.

"name:args[[@><]" Select from one of the built in kernels, using the name and geometry
arguments supplied. See NewKernelInfoBuiltIn()

"WxH[+X+Y][@><]:num, num, num ..." a kernel of size W by H, with W*H floating point numbers
following. the 'center' can be optionally be defined at +X+Y (such that +0+0 is top left corner).
If not defined the pixel in the center, for odd sizes, or to the immediate top or left of center
for even sizes is automatically selected.

"num, num, num, num, ..." list of floating point numbers defining an 'old style' odd sized
square kernel. At least 9 values should be provided for a 3x3 square kernel, 25 for a 5x5 square
kernel, 49 for 7x7, etc. Values can be space or comma separated. This is not recommended.

You can define a 'list of kernels' which can be used by some morphology operators A list is
defined as a semi-colon separated list kernels.

" kernel ; kernel ; kernel ; "

Any extra ';' characters, at start, end or between kernel defintions are simply ignored.

The special flags will expand a single kernel, into a list of rotated kernels. A '@' flag will
expand a 3x3 kernel into a list of 45-degree cyclic rotations, while a '>' will generate a list
of 90-degree rotations. The '<' also exands using 90-degree rotates, but giving a 180-degree
reflected kernel before the +/- 90-degree rotations, which can be important for Thinning operations.

Note that 'name' kernels will start with an alphabetic character while the new kernel specification
has a ':' character in its specification string. If neither is the case, it is assumed an old style
of a simple list of numbers generating a odd-sized square kernel has been given.
*/
func NewKernelInfo(kernel string) (*KernelInfo, error) {
	ckernel := C.CString(kernel)
	defer C.free(unsafe.Pointer(ckernel))

	var exc C.ExceptionInfo
	kernel_info := C.AcquireKernelInfo(ckernel, &exc)
	if err := checkExceptionInfo(&exc); err != nil {
		return nil, err
	}

	return newKernelInfo(kernel_info), nil
}

// Create a kernel from a builtin in kernel.
//
// Returns an *ExceptionInfo on error
//
// See http://www.imagemagick.org/Usage/morphology/#kernel
// for examples. Currently the 'rotation' symbols are not supported. Example:
// kernel_info := NewKernelInfoBuiltIn(KERNEL_RING, "2,1")
func NewKernelInfoBuiltIn(typ KernelInfoType, kernel string) (*KernelInfo, error) {
	var ginfo C.GeometryInfo

	ckernel := C.CString(kernel)
	defer C.free(unsafe.Pointer(ckernel))

	result := C.ParseGeometry(ckernel, &ginfo)
	gflags := int(result)

	cleanGeometryInfo(typ, gflags, &ginfo)

	var exc C.ExceptionInfo
	kernel_info := C.AcquireKernelBuiltIn(C.KernelInfoType(typ), &ginfo, &exc)
	if err := checkExceptionInfo(&exc); err != nil {
		return nil, err
	}

	return newKernelInfo(kernel_info), nil
}

// ScaleKernelInfo() scales the given kernel list by the given amount, with or without
// normalization of the sum of the kernel values (as per given flags). The exact behaviour
// of this function depends on the normalization type being used please see
// http://www.imagemagick.org/api/morphology.php#ScaleKernelInfo for details.
//
// Flag should be one of:
//     KERNEL_NORMALIZE_NONE
//     KERNEL_NORMALIZE_VALUE
//     KERNEL_NORMALIZE_CORRELATE
//     KERNEL_NORMALIZE_PERCENT
func (ki *KernelInfo) Scale(scale float64, normalizeType KernelNormalizeType) {
	C.ScaleKernelInfo(ki.info, C.double(scale), C.GeometryFlags(normalizeType))
	runtime.KeepAlive(ki)
}

// cleanGeometryInfo peforms some tidy up of the geometry info for the kernel.
func cleanGeometryInfo(typ KernelInfoType, geometryFlags int,
	geometryInfo *C.GeometryInfo) {

	// special handling of missing values in input string
	switch typ {

	// Shape Kernel Defaults
	case KERNEL_UNITY:
		if (geometryFlags & WIDTHVALUE) == 0 {
			geometryInfo.rho = 1.0 /* Default scale = 1.0, zero is valid */
		}

	case KERNEL_SQUARE, KERNEL_DIAMOND, KERNEL_OCTAGON,
		KERNEL_DISK, KERNEL_PLUS, KERNEL_CROSS:
		if (geometryFlags & HEIGHTVALUE) == 0 {
			geometryInfo.sigma = 1.0 /* Default scale = 1.0, zero is valid */
		}

	case KERNEL_RING:
		if (geometryFlags & XVALUE) == 0 {
			geometryInfo.xi = 1.0 /* Default scale = 1.0, zero is valid */
		}

	case KERNEL_RECTANGLE:
		// Rectangle - set size defaults
		if (geometryFlags & WIDTHVALUE) == 0 { /* if no width then */
			geometryInfo.rho = geometryInfo.sigma /* then  width = height */
		}
		if geometryInfo.rho < 1.0 { /* if width too small */
			geometryInfo.rho = 3 /* then  width = 3 */
		}
		if geometryInfo.sigma < 1.0 { /* if height too small */
			geometryInfo.sigma = geometryInfo.rho /* then  height = width */
		}
		//TODO - casting shenanigans
		//if ((geometryFlags & XVALUE) == 0) {    /* center offset if not defined */
		//	geometryInfo.xi = (double)(((ssize_t)geometryInfo.rho-1)/2);
		//}
		//if ((geometryFlags & YVALUE) == 0) {
		//	geometryInfo.psi = (double)(((ssize_t)geometryInfo.sigma-1)/2);
		//}

	// Distance Kernel Defaults
	case KERNEL_CHEBYSHEV, KERNEL_MANHATTAN, KERNEL_OCTAGONAL, KERNEL_EUCLIDEAN:
		if (geometryFlags & HEIGHTVALUE) == 0 { /* no distance scale */
			geometryInfo.sigma = 100.0 /* default distance scaling */
		}
		//TODO casting shenanigans
		//else if ((flags & AspectValue ) != 0) {     /* '!' flag */
		//	geometryInfo.sigma = QuantumRange/(geometryInfo.sigma+1); /* maximum pixel distance */
		//}
		//else if ((flags & PercentValue ) != 0) {    /* '%' flag */
		//	geometryInfo.sigma *= QuantumRange/100.0;         /* percentage of color range */
		//}

	}
}
