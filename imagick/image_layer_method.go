// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type ImageLayerMethod int

const (
	IMAGE_LAYER_UNDEFINED       ImageLayerMethod = C.UndefinedLayer
	IMAGE_LAYER_COALESCE        ImageLayerMethod = C.CoalesceLayer
	IMAGE_LAYER_COMPARE_ANY     ImageLayerMethod = C.CompareAnyLayer
	IMAGE_LAYER_COMPARE_CLEAR   ImageLayerMethod = C.CompareClearLayer
	IMAGE_LAYER_COMPARE_OVERLAY ImageLayerMethod = C.CompareOverlayLayer
	IMAGE_LAYER_DISPOSE         ImageLayerMethod = C.DisposeLayer
	IMAGE_LAYER_OPTIMIZE        ImageLayerMethod = C.OptimizeLayer
	IMAGE_LAYER_OPTIMIZE_IMAGE  ImageLayerMethod = C.OptimizeImageLayer
	IMAGE_LAYER_OPTIMIZE_PLUS   ImageLayerMethod = C.OptimizePlusLayer
	IMAGE_LAYER_OPTIMIZE_TRANS  ImageLayerMethod = C.OptimizeTransLayer
	IMAGE_LAYER_REMOVE_DUPS     ImageLayerMethod = C.RemoveDupsLayer
	IMAGE_LAYER_REMOVE_ZERO     ImageLayerMethod = C.RemoveZeroLayer
	IMAGE_LAYER_COMPOSITE       ImageLayerMethod = C.CompositeLayer
	IMAGE_LAYER_MERGE           ImageLayerMethod = C.MergeLayer
	IMAGE_LAYER_FLATTEN         ImageLayerMethod = C.FlattenLayer
	IMAGE_LAYER_MOSAIC          ImageLayerMethod = C.MosaicLayer
	IMAGE_LAYER_TRIM_BOUNDS     ImageLayerMethod = C.TrimBoundsLayer
)
