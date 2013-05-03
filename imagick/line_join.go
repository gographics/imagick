package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

type LineJoin int

const (
	LINE_JOIN_UNDEFINED LineJoin = C.UndefinedJoin
	LINE_JOIN_MITER     LineJoin = C.MiterJoin
	LINE_JOIN_ROUND     LineJoin = C.RoundJoin
	LINE_JOIN_BEVEL     LineJoin = C.BevelJoin
)
