package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

type DisposeType int

const (
	DISPOSE_UNRECOGNIZED DisposeType = C.UnrecognizedDispose
	DISPOSE_UNDEFINED    DisposeType = C.UndefinedDispose
	DISPOSE_NONE         DisposeType = C.NoneDispose
	DISPOSE_BACKGROUND   DisposeType = C.BackgroundDispose
	DISPOSE_PREVIOUS     DisposeType = C.PreviousDispose
)
