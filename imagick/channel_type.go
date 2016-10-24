// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type ChannelType int

const (
	CHANNEL_UNDEFINED  ChannelType = C.UndefinedChannel
	CHANNEL_RED        ChannelType = C.RedChannel
	CHANNEL_GRAY       ChannelType = C.GrayChannel
	CHANNEL_CYAN       ChannelType = C.CyanChannel
	CHANNEL_GREEN      ChannelType = C.GreenChannel
	CHANNEL_MAGENTA    ChannelType = C.MagentaChannel
	CHANNEL_BLUE       ChannelType = C.BlueChannel
	CHANNEL_YELLOW     ChannelType = C.YellowChannel
	CHANNEL_ALPHA      ChannelType = C.AlphaChannel
	CHANNEL_OPACITY    ChannelType = C.OpacityChannel
	CHANNEL_BLACK      ChannelType = C.BlackChannel
	CHANNEL_INDEX      ChannelType = C.IndexChannel
	CHANNEL_TRUE_ALPHA ChannelType = C.TrueAlphaChannel
	CHANNELS_COMPOSITE ChannelType = C.CompositeChannels
	CHANNELS_ALL       ChannelType = C.AllChannels
	CHANNELS_RGB       ChannelType = C.RGBChannels
	CHANNELS_GRAY      ChannelType = C.GrayChannels
	CHANNELS_SYNC      ChannelType = C.SyncChannels
	CHANNELS_DEFAULT   ChannelType = C.DefaultChannels
)
