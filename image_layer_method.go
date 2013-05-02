package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
)

type ImageLayerMethod int

const (
	UndefinedLayer      ImageLayerMethod = C.UndefinedLayer
	CoalesceLayer       ImageLayerMethod = C.CoalesceLayer
	CompareAnyLayer     ImageLayerMethod = C.CompareAnyLayer
	CompareClearLayer   ImageLayerMethod = C.CompareClearLayer
	CompareOverlayLayer ImageLayerMethod = C.CompareOverlayLayer
	DisposeLayer        ImageLayerMethod = C.DisposeLayer
	OptimizeLayer       ImageLayerMethod = C.OptimizeLayer
	OptimizeImageLayer  ImageLayerMethod = C.OptimizeImageLayer
	OptimizePlusLayer   ImageLayerMethod = C.OptimizePlusLayer
	OptimizeTransLayer  ImageLayerMethod = C.OptimizeTransLayer
	RemoveDupsLayer     ImageLayerMethod = C.RemoveDupsLayer
	RemoveZeroLayer     ImageLayerMethod = C.RemoveZeroLayer
	CompositeLayer      ImageLayerMethod = C.CompositeLayer
	MergeLayer          ImageLayerMethod = C.MergeLayer
	FlattenLayer        ImageLayerMethod = C.FlattenLayer
	MosaicLayer         ImageLayerMethod = C.MosaicLayer
	TrimBoundsLayer     ImageLayerMethod = C.TrimBoundsLayer
)

var imageLayerMethodStrings = map[ImageLayerMethod]string{
	UndefinedLayer:      "UndefinedLayer",
	CoalesceLayer:       "CoalesceLayer",
	CompareAnyLayer:     "CompareAnyLayer",
	CompareClearLayer:   "CompareClearLayer",
	CompareOverlayLayer: "CompareOverlayLayer",
	DisposeLayer:        "DisposeLayer",
	OptimizeLayer:       "OptimizeLayer",
	OptimizeImageLayer:  "OptimizeImageLayer",
	OptimizePlusLayer:   "OptimizePlusLayer",
	OptimizeTransLayer:  "OptimizeTransLayer",
	RemoveDupsLayer:     "RemoveDupsLayer",
	RemoveZeroLayer:     "RemoveZeroLayer",
	CompositeLayer:      "CompositeLayer",
	MergeLayer:          "MergeLayer",
	FlattenLayer:        "FlattenLayer",
	MosaicLayer:         "MosaicLayer",
	TrimBoundsLayer:     "TrimBoundsLayer",
}

func (cst *ImageLayerMethod) String() string {
	if v, ok := imageLayerMethodStrings[ImageLayerMethod(*cst)]; ok {
		return v
	}
	return fmt.Sprintf("UnknownImageLayerMethod[%d]", *cst)
}
