package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

type MagickFunction int

const (
	FUNCTION_UNDEFINED  MagickFunction = C.UndefinedFunction
	FUNCTION_POLYNOMIAL MagickFunction = C.PolynomialFunction
	FUNCTION_SINUSOID   MagickFunction = C.SinusoidFunction
	FUNCTION_ARCSIN     MagickFunction = C.ArcsinFunction
	FUNCTION_ARCTAN     MagickFunction = C.ArctanFunction
)
