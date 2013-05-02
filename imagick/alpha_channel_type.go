package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

type AlphaChannelType int

const (
	ALPHA_CHANNEL_UNDEFINED   AlphaChannelType = C.UndefinedAlphaChannel
	ALPHA_CHANNEL_ACTIVATE    AlphaChannelType = C.ActivateAlphaChannel
	ALPHA_CHANNEL_BACKGROUND  AlphaChannelType = C.BackgroundAlphaChannel
	ALPHA_CHANNEL_COPY        AlphaChannelType = C.CopyAlphaChannel
	ALPHA_CHANNEL_DEACTIVATE  AlphaChannelType = C.DeactivateAlphaChannel
	ALPHA_CHANNEL_EXTRACT     AlphaChannelType = C.ExtractAlphaChannel
	ALPHA_CHANNEL_OPAQUE      AlphaChannelType = C.OpaqueAlphaChannel
	ALPHA_CHANNEL_RESET       AlphaChannelType = C.ResetAlphaChannel
	ALPHA_CHANNEL_SET         AlphaChannelType = C.SetAlphaChannel
	ALPHA_CHANNEL_SHAPE       AlphaChannelType = C.ShapeAlphaChannel
	ALPHA_CHANNEL_TRANSPARENT AlphaChannelType = C.TransparentAlphaChannel
	ALPHA_CHANNEL_FLATTEN     AlphaChannelType = C.FlattenAlphaChannel
	ALPHA_CHANNEL_REMOVE      AlphaChannelType = C.RemoveAlphaChannel
)
