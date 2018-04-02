package imagick

/*
#include <wand/MagickWand.h>
*/
import "C"

type GradientType int

const (
	GRADIENT_TYPE_UNDEFINED GradientType = C.UndefinedGradient
	GRADIENT_TYPE_LINEAR    GradientType = C.LinearGradient
	GRADIENT_TYPE_RADIAL    GradientType = C.RadialGradient
)
