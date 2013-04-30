package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
)

type AlphaChannelType int

const (
	UndefinedAlphaChannel AlphaChannelType = iota
	ActivateAlphaChannel
	BackgroundAlphaChannel
	CopyAlphaChannel
	DeactivateAlphaChannel
	ExtractAlphaChannel
	OpaqueAlphaChannel
	ResetAlphaChannel // Deprecated
	SetAlphaChannel
	ShapeAlphaChannel
	TransparentAlphaChannel
	FlattenAlphaChannel
	RemoveAlphaChannel
)

var alphaChannelTypeStrings = map[AlphaChannelType]string{
	UndefinedAlphaChannel:   "UndefinedAlphaChannel",
	ActivateAlphaChannel:    "ActivateAlphaChannel",
	BackgroundAlphaChannel:  "BackgroundAlphaChannel",
	CopyAlphaChannel:        "CopyAlphaChannel",
	DeactivateAlphaChannel:  "DeactivateAlphaChannel",
	ExtractAlphaChannel:     "ExtractAlphaChannel",
	OpaqueAlphaChannel:      "OpaqueAlphaChannel",
	ResetAlphaChannel:       "ResetAlphaChannel",
	SetAlphaChannel:         "SetAlphaChannel",
	ShapeAlphaChannel:       "ShapeAlphaChannel",
	TransparentAlphaChannel: "TransparentAlphaChannel",
	FlattenAlphaChannel:     "FlattenAlphaChannel",
	RemoveAlphaChannel:      "RemoveAlphaChannel",
}

func (ct *AlphaChannelType) String() string {
	if v, ok := alphaChannelTypeStrings[AlphaChannelType(*ct)]; ok {
		return v
	}
	return fmt.Sprintf("UnknownCompression[%d]", *ct)
}
