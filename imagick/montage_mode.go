package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
)

type MontageMode int

const (
	UndefinedMode   MontageMode = C.UndefinedMode
	FrameMode       MontageMode = C.FrameMode
	UnframeMode     MontageMode = C.UnframeMode
	ConcatenateMode MontageMode = C.ConcatenateMode
)

var montageModeStrings = map[MontageMode]string{
	UndefinedMode:   "UndefinedMode",
	FrameMode:       "FrameMode",
	UnframeMode:     "UnframeMode",
	ConcatenateMode: "ConcatenateMode",
}

func (ct *MontageMode) String() string {
	if v, ok := montageModeStrings[MontageMode(*ct)]; ok {
		return v
	}
	return fmt.Sprintf("UnknownCompression[%d]", *ct)
}
