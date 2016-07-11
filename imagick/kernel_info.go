// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <magick/MagickCore.h>
#include <magick/geometry.h>
#include <magick/morphology.h>
*/
import "C"
//import "runtime"
import "unsafe"


// This struct represents the KernelInfo C API of ImageMagick
type KernelInfo struct {
	info *C.KernelInfo
}

func newKernelInfo(cki *C.KernelInfo) *KernelInfo {
	ki := &KernelInfo{info: cki}
	//TODO - understand and implement this.
	// runtime.SetFinalizer(ki, Destroy)
	//ki.IncreaseCount()
	return ki
}


// Convert the current KernelInfo to an 2d-array of values. The values are either
// a float64 if the element is used, or NaN if the element is not used by the kernel
func (kernel_info *KernelInfo) ToArray() [][]float64 {
	count := 0
	values := [][]float64{}

	for y := C.size_t(0); y<kernel_info.info.height ; y++ {
		row_values := make([]float64, kernel_info.info.width)
		for x:= C.size_t(0); x<kernel_info.info.width ; x++ {
			p2 := (*[1<<10]C.double)(unsafe.Pointer(kernel_info.info.values))
			row_values[x] = float64(p2[count])
			count++;
		}
		values = append(values, row_values)
	}

	return values
}

// Create a kernel from a builtin in kernel. See http://www.imagemagick.org/Usage/morphology/#kernel
// for examples. Currently the 'rotation' symbols are not supported. Example:
// kernel_info := imagick.AcquireKernelBuiltIn(imagick.KERNEL_RING, "2,1")
func AcquireKernelBuiltIn(kernel_type KernelInfoType, kernelString string) *KernelInfo  {
	geometry_info := (*GeometryInfo)(C.malloc(C.size_t(unsafe.Sizeof(C.GeometryInfo{}))))
	result := C.ParseGeometry(C.CString(kernelString), &geometry_info.gi);
	var geometry_flags int = int(result)
	FiddleWithGeometryInfo(kernel_type, geometry_flags, geometry_info)
	kernel_info := C.AcquireKernelBuiltIn(C.KernelInfoType(kernel_type), &geometry_info.gi);

	return newKernelInfo(kernel_info);
}

// ScaleKernelInfo() scales the given kernel list by the given amount, with or without
// normalization of the sum of the kernel values (as per given flags). The exact behaviour
// of this function depends on the normalization type being used please see
// http://www.imagemagick.org/api/morphology.php#ScaleKernelInfo for details.
// Flag should be one of:
// imagick.KERNEL_NORMALIZE_NONE
// imagick.KERNEL_NORMALIZE_VALUE
// imagick.KERNEL_NORMALIZE_CORRELATE
// imagick.KERNEL_NORMALIZE_PERCENT
func (kernel_info *KernelInfo) Scale(scale float64, normalize_type KernelNormalizeType) {
	C.ScaleKernelInfo(kernel_info.info, C.double(scale), C.GeometryFlags(normalize_type));
}

// This does .....stuff. Basically some tidy up of the kernel is required apparently.
func FiddleWithGeometryInfo(kernel_type KernelInfoType, geometry_flags int, geometry_info *GeometryInfo) {

	/* special handling of missing values in input string */
	switch( kernel_type ) {
		/* Shape Kernel Defaults */
		case KERNEL_UNITY: {
			if ((geometry_flags & WIDTHVALUE) == 0) {
				geometry_info.gi.rho = 1.0; /* Default scale = 1.0, zero is valid */
			}
			break;
		}
		case KERNEL_SQUARE:
		case KERNEL_DIAMOND:
		case KERNEL_OCTAGON:
		case KERNEL_DISK:
		case KERNEL_PLUS:
		case KERNEL_CROSS: {
			if ( (geometry_flags & HEIGHTVALUE) == 0 ) {
				geometry_info.gi.sigma = 1.0;    /* Default scale = 1.0, zero is valid */
			}
			break;
		}
		case KERNEL_RING: {
			if ((geometry_flags & XVALUE) == 0) {
				geometry_info.gi.xi = 1.0;       /* Default scale = 1.0, zero is valid */
			}
			break;
		}
		case KERNEL_RECTANGLE: {    /* Rectangle - set size defaults */
			if ((geometry_flags & WIDTHVALUE) == 0) { /* if no width then */
				geometry_info.gi.rho = geometry_info.gi.sigma;         /* then  width = height */
			}
			if (geometry_info.gi.rho < 1.0) {            /* if width too small */
				geometry_info.gi.rho = 3;                 /* then  width = 3 */
			}
			if (geometry_info.gi.sigma < 1.0) {          /* if height too small */
				geometry_info.gi.sigma = geometry_info.gi.rho;         /* then  height = width */
			}
			//TODO - casting shenanigans
			//if ((geometry_flags & XVALUE) == 0) {    /* center offset if not defined */
			//	geometry_info.gi.xi = (double)(((ssize_t)geometry_info.gi.rho-1)/2);
			//}
			//if ((geometry_flags & YVALUE) == 0) {
			//	geometry_info.gi.psi = (double)(((ssize_t)geometry_info.gi.sigma-1)/2);
			//}
			break;
		}
		/* Distance Kernel Defaults */
		case KERNEL_CHEBYSHEV:
		case KERNEL_MANHATTAN:
		case KERNEL_OCTAGONAL:
		case KERNEL_EUCLIDEAN: {
			if ((geometry_flags & HEIGHTVALUE) == 0) {           /* no distance scale */
				geometry_info.gi.sigma = 100.0;                       /* default distance scaling */
			}
			//TODO casting shenanigans
			//else if ((flags & AspectValue ) != 0) {     /* '!' flag */
			//	geometry_info.gi.sigma = QuantumRange/(geometry_info.gi.sigma+1); /* maximum pixel distance */
			//}
			//else if ((flags & PercentValue ) != 0) {    /* '%' flag */
			//	geometry_info.gi.sigma *= QuantumRange/100.0;         /* percentage of color range */
			//}
			break;
		}
		default: {
			break;
		}
	}
}


