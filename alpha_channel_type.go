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
	UndefinedAlphaChannel   AlphaChannelType = C.UndefinedAlphaChannel
	ActivateAlphaChannel    AlphaChannelType = C.ActivateAlphaChannel
	BackgroundAlphaChannel  AlphaChannelType = C.BackgroundAlphaChannel
	CopyAlphaChannel        AlphaChannelType = C.CopyAlphaChannel
	DeactivateAlphaChannel  AlphaChannelType = C.DeactivateAlphaChannel
	ExtractAlphaChannel     AlphaChannelType = C.ExtractAlphaChannel
	OpaqueAlphaChannel      AlphaChannelType = C.OpaqueAlphaChannel
	ResetAlphaChannel       AlphaChannelType = C.ResetAlphaChannel
	SetAlphaChannel         AlphaChannelType = C.SetAlphaChannel
	ShapeAlphaChannel       AlphaChannelType = C.ShapeAlphaChannel
	TransparentAlphaChannel AlphaChannelType = C.TransparentAlphaChannel
	FlattenAlphaChannel     AlphaChannelType = C.FlattenAlphaChannel
	RemoveAlphaChannel      AlphaChannelType = C.RemoveAlphaChannel
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
