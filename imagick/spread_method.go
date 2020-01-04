package imagick

/*
#include <MagickCore/MagickCore.h>
*/
import "C"

type SpreadMethod int

const (
	SPREAD_METHOD_UNDEFINED SpreadMethod = C.UndefinedGradient
	SPREAD_METHOD_PAD       SpreadMethod = C.PadSpread
	SPREAD_METHOD_REFLECT   SpreadMethod = C.ReflectSpread
	SPREAD_METHOD_REPEAT    SpreadMethod = C.RepeatSpread
)
