package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type KernelNormalizeType int

const (
	// The kernel is scaled directly using given scaling factor without change.
	KERNEL_NORMALIZE_NONE KernelNormalizeType = 0

	// Kernel normalization ('normalize_flags' given) is designed to ensure
	// that any use of the kernel scaling factor with 'Convolve' or 'Correlate'
	// morphology methods will fall into -1.0 to +1.0 range.
	KERNEL_NORMALIZE_VALUE KernelNormalizeType = C.NormalizeValue

	// For special kernels designed for locating shapes using 'Correlate', (often
	// only containing +1 and -1 values, representing foreground/brackground
	// matching) a special normalization method is provided to scale the positive
	// values separately to those of the negative values, so the kernel will be
	// forced to become a zero-sum kernel better suited to such searches.
	KERNEL_NORMALIZE_CORRELATE KernelNormalizeType = C.CorrelateNormalizeValue

	// Scale the kernel by a percent.
	KERNEL_NORMALIZE_PERCENT KernelNormalizeType = C.PercentValue
)
