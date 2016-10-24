// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type MorphologyMethod int

const (
	MORPHOLOGY_UNDEFINED          MorphologyMethod = C.UndefinedMorphology
	MORPHOLOGY_CONVOLVE           MorphologyMethod = C.ConvolveMorphology          /* Weighted Sum with reflected kernel */
	MORPHOLOGY_CORRELATE          MorphologyMethod = C.CorrelateMorphology         /* Weighted Sum using a sliding window */
	MORPHOLOGY_ERODE              MorphologyMethod = C.ErodeMorphology             /* Minimum Value in Neighbourhood */
	MORPHOLOGY_DILATE             MorphologyMethod = C.DilateMorphology            /* Maximum Value in Neighbourhood */
	MORPHOLOGY_ERODE_INTENSITY    MorphologyMethod = C.ErodeIntensityMorphology    /* Pixel Pick using GreyScale Erode */
	MORPHOLOGY_DILATE_INTENSITY   MorphologyMethod = C.DilateIntensityMorphology   /* Pixel Pick using GreyScale Dialate */
	MORPHOLOGY_DISTANCE           MorphologyMethod = C.DistanceMorphology          /* Add Kernel Value, take Minimum */
	MORPHOLOGY_OPEN               MorphologyMethod = C.OpenMorphology              /* Dilate then Erode */
	MORPHOLOGY_CLOSE              MorphologyMethod = C.CloseMorphology             /* Erode then Dilate */
	MORPHOLOGY_OPEN_INTENSITY     MorphologyMethod = C.OpenIntensityMorphology     /* Pixel Pick using GreyScale Open */
	MORPHOLOGY_CLOSE_INTENSITY    MorphologyMethod = C.CloseIntensityMorphology    /* Pixel Pick using GreyScale Close */
	MORPHOLOGY_SMOOTH             MorphologyMethod = C.SmoothMorphology            /* Open then Close */
	MORPHOLOGY_EDGE_IN            MorphologyMethod = C.EdgeInMorphology            /* Dilate difference from Original */
	MORPHOLOGY_EDGE_OUT           MorphologyMethod = C.EdgeOutMorphology           /* Erode difference from Original */
	MORPHOLOGY_EDGE               MorphologyMethod = C.EdgeMorphology              /* Dilate difference with Erode */
	MORPHOLOGY_TOP_HAT            MorphologyMethod = C.TopHatMorphology            /* Close difference from Original */
	MORPHOLOGY_BOTTOM_HAT         MorphologyMethod = C.BottomHatMorphology         /* Open difference from Original */
	MORPHOLOGY_HIT_AND_MISS       MorphologyMethod = C.HitAndMissMorphology        /* Foreground/Background pattern matching */
	MORPHOLOGY_THINNING           MorphologyMethod = C.ThinningMorphology          /* Remove matching pixels from image */
	MORPHOLOGY_THICKEN            MorphologyMethod = C.ThickenMorphology           /* Add matching pixels from image */
	MORPHOLOGY_VORONOI            MorphologyMethod = C.VoronoiMorphology           /* distance matte channel copy nearest color */
	MORPHOLOGY_ITERATIVE_DISTANCE MorphologyMethod = C.IterativeDistanceMorphology /* Add Kernel Value, take Minimum */
)
