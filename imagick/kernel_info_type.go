package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type KernelInfoType int

const (
	KERNEL_UNDEFINED     KernelInfoType = C.UndefinedKernel /* equivalent to UnityKernel */
	KERNEL_UNITY         KernelInfoType = C.UnityKernel     /* The no-op or 'original image' kernel */
	KERNEL_GAUSSIAN      KernelInfoType = C.GaussianKernel  /* Convolution Kernels Gaussian Based */
	KERNEL_DOG           KernelInfoType = C.DoGKernel
	KERNEL_LOG           KernelInfoType = C.LoGKernel
	KERNEL_BLUR          KernelInfoType = C.BlurKernel
	KERNEL_COMET         KernelInfoType = C.CometKernel
	KERNEL_LAPLACIAN     KernelInfoType = C.LaplacianKernel /* Convolution Kernels by Name */
	KERNEL_SOBEL         KernelInfoType = C.SobelKernel
	KERNEL_FREICHEN      KernelInfoType = C.FreiChenKernel
	KERNEL_ROBERTS       KernelInfoType = C.RobertsKernel
	KERNEL_PREWITT       KernelInfoType = C.PrewittKernel
	KERNEL_COMPASS       KernelInfoType = C.CompassKernel
	KERNEL_KIRSCH        KernelInfoType = C.KirschKernel
	KERNEL_DIAMOND       KernelInfoType = C.DiamondKernel /* Shape Kernels */
	KERNEL_SQUARE        KernelInfoType = C.SquareKernel
	KERNEL_RECTANGLE     KernelInfoType = C.RectangleKernel
	KERNEL_OCTAGON       KernelInfoType = C.OctagonKernel
	KERNEL_DISK          KernelInfoType = C.DiskKernel
	KERNEL_PLUS          KernelInfoType = C.PlusKernel
	KERNEL_CROSS         KernelInfoType = C.CrossKernel
	KERNEL_RING          KernelInfoType = C.RingKernel
	KERNEL_PEAKS         KernelInfoType = C.PeaksKernel /* Hit And Miss Kernels */
	KERNEL_EDGES         KernelInfoType = C.EdgesKernel
	KERNEL_CORNERS       KernelInfoType = C.CornersKernel
	KERNEL_DIAGONALS     KernelInfoType = C.DiagonalsKernel
	KERNEL_LINEENDS      KernelInfoType = C.LineEndsKernel
	KERNEL_LINEJUNCTIONS KernelInfoType = C.LineJunctionsKernel
	KERNEL_RIDGES        KernelInfoType = C.RidgesKernel
	KERNEL_CONVEXHULL    KernelInfoType = C.ConvexHullKernel
	KERNEL_THINSE        KernelInfoType = C.ThinSEKernel
	KERNEL_SKELETON      KernelInfoType = C.SkeletonKernel
	KERNEL_CHEBYSHEV     KernelInfoType = C.ChebyshevKernel /* Distance Measuring Kernels */
	KERNEL_MANHATTAN     KernelInfoType = C.ManhattanKernel
	KERNEL_OCTAGONAL     KernelInfoType = C.OctagonalKernel
	KERNEL_EUCLIDEAN     KernelInfoType = C.EuclideanKernel
	KERNEL_USERDEFINED   KernelInfoType = C.UserDefinedKernel /* User Specified Kernel Array */
	KERNEL_BINOMIAL      KernelInfoType = C.BinomialKernel
)
