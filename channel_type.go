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
	UndefinedChannel  ChannelType = C.UndefinedChannel
	RedChannel        ChannelType = C.RedChannel
	GrayChannel       ChannelType = C.GrayChannel
	CyanChannel       ChannelType = C.CyanChannel
	GreenChannel      ChannelType = C.GreenChannel
	MagentaChannel    ChannelType = C.MagentaChannel
	BlueChannel       ChannelType = C.BlueChannel
	YellowChannel     ChannelType = C.YellowChannel
	AlphaChannel      ChannelType = C.AlphaChannel
	OpacityChannel    ChannelType = C.OpacityChannel
	BlackChannel      ChannelType = C.BlackChannel
	IndexChannel      ChannelType = C.IndexChannel
	CompositeChannels ChannelType = C.CompositeChannels
	AllChannels       ChannelType = C.AllChannels
	TrueAlphaChannel  ChannelType = C.TrueAlphaChannel
	RGBChannels       ChannelType = C.RGBChannels
	GrayChannels      ChannelType = C.GrayChannels
	SyncChannels      ChannelType = C.SyncChannels
	DefaultChannels   ChannelType = C.DefaultChannels
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
