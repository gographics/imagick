package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
)

type VirtualPixelMethod int

const (
	UndefinedVirtualPixelMethod VirtualPixelMethod = iota
	BackgroundVirtualPixelMethod
	ConstantVirtualPixelMethod // deprecated
	DitherVirtualPixelMethod
	EdgeVirtualPixelMethod
	MirrorVirtualPixelMethod
	RandomVirtualPixelMethod
	TileVirtualPixelMethod
	TransparentVirtualPixelMethod
	MaskVirtualPixelMethod
	BlackVirtualPixelMethod
	GrayVirtualPixelMethod
	WhiteVirtualPixelMethod
	HorizontalTileVirtualPixelMethod
	VerticalTileVirtualPixelMethod
	HorizontalTileEdgeVirtualPixelMethod
	VerticalTileEdgeVirtualPixelMethod
	CheckerTileVirtualPixelMethod
)

var virtualPixelMethodStrings = map[VirtualPixelMethod]string{
	UndefinedVirtualPixelMethod:          "UndefinedVirtualPixelMethod",
	BackgroundVirtualPixelMethod:         "BackgroundVirtualPixelMethod",
	ConstantVirtualPixelMethod:           "ConstantVirtualPixelMethod",
	DitherVirtualPixelMethod:             "DitherVirtualPixelMethod",
	EdgeVirtualPixelMethod:               "EdgeVirtualPixelMethod",
	MirrorVirtualPixelMethod:             "MirrorVirtualPixelMethod",
	RandomVirtualPixelMethod:             "RandomVirtualPixelMethod",
	TileVirtualPixelMethod:               "TileVirtualPixelMethod",
	TransparentVirtualPixelMethod:        "TransparentVirtualPixelMethod",
	MaskVirtualPixelMethod:               "MaskVirtualPixelMethod",
	BlackVirtualPixelMethod:              "BlackVirtualPixelMethod",
	GrayVirtualPixelMethod:               "GrayVirtualPixelMethod",
	WhiteVirtualPixelMethod:              "WhiteVirtualPixelMethod",
	HorizontalTileVirtualPixelMethod:     "HorizontalTileVirtualPixelMethod",
	VerticalTileVirtualPixelMethod:       "VerticalTileVirtualPixelMethod",
	HorizontalTileEdgeVirtualPixelMethod: "HorizontalTileEdgeVirtualPixelMethod",
	VerticalTileEdgeVirtualPixelMethod:   "VerticalTileEdgeVirtualPixelMethod",
	CheckerTileVirtualPixelMethod:        "CheckerTileVirtualPixelMethod",
}

func (ct *VirtualPixelMethod) String() string {
	if v, ok := virtualPixelMethodStrings[VirtualPixelMethod(*ct)]; ok {
		return v
	}
	return fmt.Sprintf("UnknownCompression[%d]", *ct)
}
