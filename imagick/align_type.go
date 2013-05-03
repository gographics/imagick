package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

type AlignType int

const (
	ALIGN_UNDEFINED AlignType = C.UndefinedAlign
	ALIGN_LEFT      AlignType = C.LeftAlign
	ALIGN_CENTER    AlignType = C.CenterAlign
	ALIGN_RIGHT     AlignType = C.RightAlign
)
