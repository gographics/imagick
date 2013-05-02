package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
)

type DisposeType int

const (
	UnrecognizedDispose DisposeType = C.UnrecognizedDispose
	UndefinedDispose    DisposeType = C.UndefinedDispose
	NoneDispose         DisposeType = C.NoneDispose
	BackgroundDispose   DisposeType = C.BackgroundDispose
	PreviousDispose     DisposeType = C.PreviousDispose
)

var disposeTypeStrings = map[DisposeType]string{
	UndefinedDispose:  "UndefinedDispose",
	NoneDispose:       "NoneDispose",
	BackgroundDispose: "BackgroundDispose",
	PreviousDispose:   "PreviousDispose",
}

func (cst *DisposeType) String() string {
	if v, ok := disposeTypeStrings[DisposeType(*cst)]; ok {
		return v
	}
	return fmt.Sprintf("UnknownDisposeType[%d]", *cst)
}
