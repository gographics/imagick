// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type AlphaChannelType int

const (
	ALPHA_CHANNEL_UNDEFINED    AlphaChannelType = C.UndefinedAlphaChannel
	ALPHA_CHANNEL_ACTIVATE     AlphaChannelType = C.ActivateAlphaChannel
	ALPHA_CHANNEL_ASSOCIATE    AlphaChannelType = C.AssociateAlphaChannel
	ALPHA_CHANNEL_BACKGROUND   AlphaChannelType = C.BackgroundAlphaChannel
	ALPHA_CHANNEL_COPY         AlphaChannelType = C.CopyAlphaChannel
	ALPHA_CHANNEL_DEACTIVATE   AlphaChannelType = C.DeactivateAlphaChannel
	ALPHA_CHANNEL_DISASSOCIATE AlphaChannelType = C.DisassociateAlphaChannel
	ALPHA_CHANNEL_DISCRETE     AlphaChannelType = C.DiscreteAlphaChannel
	ALPHA_CHANNEL_EXTRACT      AlphaChannelType = C.ExtractAlphaChannel
	ALPHA_CHANNEL_OFF          AlphaChannelType = C.OffAlphaChannel
	ALPHA_CHANNEL_ON           AlphaChannelType = C.OnAlphaChannel
	ALPHA_CHANNEL_OPAQUE       AlphaChannelType = C.OpaqueAlphaChannel
	ALPHA_CHANNEL_REMOVE       AlphaChannelType = C.RemoveAlphaChannel
	ALPHA_CHANNEL_SET          AlphaChannelType = C.SetAlphaChannel
	ALPHA_CHANNEL_SHAPE        AlphaChannelType = C.ShapeAlphaChannel
	ALPHA_CHANNEL_TRANSPARENT  AlphaChannelType = C.TransparentAlphaChannel
)
