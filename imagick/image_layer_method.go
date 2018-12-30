// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type LayerMethod int

const (
	IMAGE_LAYER_UNDEFINED       LayerMethod = C.UndefinedLayer
	IMAGE_LAYER_COALESCE        LayerMethod = C.CoalesceLayer
	IMAGE_LAYER_COMPARE_ANY     LayerMethod = C.CompareAnyLayer
	IMAGE_LAYER_COMPARE_CLEAR   LayerMethod = C.CompareClearLayer
	IMAGE_LAYER_COMPARE_OVERLAY LayerMethod = C.CompareOverlayLayer
	IMAGE_LAYER_DISPOSE         LayerMethod = C.DisposeLayer
	IMAGE_LAYER_OPTIMIZE        LayerMethod = C.OptimizeLayer
	IMAGE_LAYER_OPTIMIZE_IMAGE  LayerMethod = C.OptimizeImageLayer
	IMAGE_LAYER_OPTIMIZE_PLUS   LayerMethod = C.OptimizePlusLayer
	IMAGE_LAYER_OPTIMIZE_TRANS  LayerMethod = C.OptimizeTransLayer
	IMAGE_LAYER_REMOVE_DUPS     LayerMethod = C.RemoveDupsLayer
	IMAGE_LAYER_REMOVE_ZERO     LayerMethod = C.RemoveZeroLayer
	IMAGE_LAYER_COMPOSITE       LayerMethod = C.CompositeLayer
	IMAGE_LAYER_MERGE           LayerMethod = C.MergeLayer
	IMAGE_LAYER_FLATTEN         LayerMethod = C.FlattenLayer
	IMAGE_LAYER_MOSAIC          LayerMethod = C.MosaicLayer
	IMAGE_LAYER_TRIM_BOUNDS     LayerMethod = C.TrimBoundsLayer
)
