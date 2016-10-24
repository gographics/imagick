// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type VirtualPixelMethod int

const (
	VIRTUAL_PIXEL_UNDEFINED            VirtualPixelMethod = C.UndefinedVirtualPixelMethod
	VIRTUAL_PIXEL_BACKGROUND           VirtualPixelMethod = C.BackgroundVirtualPixelMethod
	VIRTUAL_PIXEL_BLACK                VirtualPixelMethod = C.BlackVirtualPixelMethod
	VIRTUAL_PIXEL_CHECKER_TILE         VirtualPixelMethod = C.CheckerTileVirtualPixelMethod
	VIRTUAL_PIXEL_DITHER               VirtualPixelMethod = C.DitherVirtualPixelMethod
	VIRTUAL_PIXEL_EDGE                 VirtualPixelMethod = C.EdgeVirtualPixelMethod
	VIRTUAL_PIXEL_GRAY                 VirtualPixelMethod = C.GrayVirtualPixelMethod
	VIRTUAL_PIXEL_HORIZONTAL_TILE      VirtualPixelMethod = C.HorizontalTileVirtualPixelMethod
	VIRTUAL_PIXEL_HORIZONTAL_TILE_EDGE VirtualPixelMethod = C.HorizontalTileEdgeVirtualPixelMethod
	VIRTUAL_PIXEL_MASK                 VirtualPixelMethod = C.MaskVirtualPixelMethod
	VIRTUAL_PIXEL_MIRROR               VirtualPixelMethod = C.MirrorVirtualPixelMethod
	VIRTUAL_PIXEL_RANDOM               VirtualPixelMethod = C.RandomVirtualPixelMethod
	VIRTUAL_PIXEL_TILE                 VirtualPixelMethod = C.TileVirtualPixelMethod
	VIRTUAL_PIXEL_TRANSPARENT          VirtualPixelMethod = C.TransparentVirtualPixelMethod
	VIRTUAL_PIXEL_VERTICAL_TILE        VirtualPixelMethod = C.VerticalTileVirtualPixelMethod
	VIRTUAL_PIXEL_VERTICAL_TILE_EDGE   VirtualPixelMethod = C.VerticalTileEdgeVirtualPixelMethod
	VIRTUAL_PIXEL_WHITE                VirtualPixelMethod = C.WhiteVirtualPixelMethod
)
