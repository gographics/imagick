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
	UnrecognizedDispose DisposeType = iota
)

const (
	UndefinedDispose DisposeType = iota
	NoneDispose
	BackgroundDispose
	PreviousDispose
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
