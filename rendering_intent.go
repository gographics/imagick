package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
)

type RenderingIntent int

const (
	UndefinedIntent RenderingIntent = iota
	SaturationIntent
	PerceptualIntent
	AbsoluteIntent
	RelativeIntent
)

var renderingIntentStrings = map[RenderingIntent]string{
	UndefinedIntent:  "UndefinedIntent",
	SaturationIntent: "SaturationIntent",
	PerceptualIntent: "PerceptualIntent",
	AbsoluteIntent:   "AbsoluteIntent",
	RelativeIntent:   "RelativeIntent",
}

func (ct *RenderingIntent) String() string {
	if v, ok := renderingIntentStrings[RenderingIntent(*ct)]; ok {
		return v
	}
	return fmt.Sprintf("UnknownCompression[%d]", *ct)
}
