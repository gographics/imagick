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
	UndefinedVirtualPixelMethod          VirtualPixelMethod = C.UndefinedVirtualPixelMethod
	BackgroundVirtualPixelMethod         VirtualPixelMethod = C.BackgroundVirtualPixelMethod
	ConstantVirtualPixelMethod           VirtualPixelMethod = C.ConstantVirtualPixelMethod
	DitherVirtualPixelMethod             VirtualPixelMethod = C.DitherVirtualPixelMethod
	EdgeVirtualPixelMethod               VirtualPixelMethod = C.EdgeVirtualPixelMethod
	MirrorVirtualPixelMethod             VirtualPixelMethod = C.MirrorVirtualPixelMethod
	RandomVirtualPixelMethod             VirtualPixelMethod = C.RandomVirtualPixelMethod
	TileVirtualPixelMethod               VirtualPixelMethod = C.TileVirtualPixelMethod
	TransparentVirtualPixelMethod        VirtualPixelMethod = C.TransparentVirtualPixelMethod
	MaskVirtualPixelMethod               VirtualPixelMethod = C.MaskVirtualPixelMethod
	BlackVirtualPixelMethod              VirtualPixelMethod = C.BlackVirtualPixelMethod
	GrayVirtualPixelMethod               VirtualPixelMethod = C.GrayVirtualPixelMethod
	WhiteVirtualPixelMethod              VirtualPixelMethod = C.WhiteVirtualPixelMethod
	HorizontalTileVirtualPixelMethod     VirtualPixelMethod = C.HorizontalTileVirtualPixelMethod
	VerticalTileVirtualPixelMethod       VirtualPixelMethod = C.VerticalTileVirtualPixelMethod
	HorizontalTileEdgeVirtualPixelMethod VirtualPixelMethod = C.HorizontalTileEdgeVirtualPixelMethod
	VerticalTileEdgeVirtualPixelMethod   VirtualPixelMethod = C.VerticalTileEdgeVirtualPixelMethod
	CheckerTileVirtualPixelMethod        VirtualPixelMethod = C.CheckerTileVirtualPixelMethod
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
