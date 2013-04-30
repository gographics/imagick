package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
)

type ChannelType int

const (
	UndefinedChannel  ChannelType = iota
	RedChannel                    = 0x0001
	GrayChannel                   = 0x0001
	CyanChannel                   = 0x0001
	GreenChannel                  = 0x0002
	MagentaChannel                = 0x0002
	BlueChannel                   = 0x0004
	YellowChannel                 = 0x0004
	AlphaChannel                  = 0x0008
	OpacityChannel                = 0x0008
	BlackChannel                  = 0x0020
	IndexChannel                  = 0x0020
	CompositeChannels             = 0x002F
	AllChannels                   = 0x7ffffff
	TrueAlphaChannel              = 0x0040
	RGBChannels                   = 0x0080
	GrayChannels                  = 0x0080
	SyncChannels                  = 0x0100
	DefaultChannels               = ((AllChannels | SyncChannels) &^ OpacityChannel)
)

var channelTypeStrings = map[ChannelType]string{
	UndefinedChannel:  "UndefinedChannel",
	RedChannel:        "RedOrGrayOrCyanChannel",
	GreenChannel:      "GreenOrMagentaChannel",
	BlueChannel:       "BlueOrYellowChannel",
	AlphaChannel:      "AlphaOrOpacityChannel",
	BlackChannel:      "BlackOrIndexChannel",
	CompositeChannels: "CompositeChannels",
	AllChannels:       "AllChannels",
	TrueAlphaChannel:  "TrueAlphaChannel",
	RGBChannels:       "RGBOrGrayChannels",
	SyncChannels:      "SyncChannels",
	DefaultChannels:   "DefaultChannels",
}

func (cst *ChannelType) String() string {
	if v, ok := channelTypeStrings[ChannelType(*cst)]; ok {
		return v
	}
	return fmt.Sprintf("UnknownChannel[%d]", *cst)
}
