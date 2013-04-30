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
	UndefinedLayer ImageLayerMethod = iota
	CoalesceLayer
	CompareAnyLayer
	CompareClearLayer
	CompareOverlayLayer
	DisposeLayer
	OptimizeLayer
	OptimizeImageLayer
	OptimizePlusLayer
	OptimizeTransLayer
	RemoveDupsLayer
	RemoveZeroLayer
	CompositeLayer
	MergeLayer
	FlattenLayer
	MosaicLayer
	TrimBoundsLayer
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
