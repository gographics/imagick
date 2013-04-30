package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
)

type MorphologyMethod int

const (
	UndefinedMorphology MorphologyMethod = iota
	/* Convolve / Correlate weighted sums */
	ConvolveMorphology  /* Weighted Sum with reflected kernel */
	CorrelateMorphology /* Weighted Sum using a sliding window */
	/* Low-level Morphology methods */
	ErodeMorphology           /* Minimum Value in Neighbourhood */
	DilateMorphology          /* Maximum Value in Neighbourhood */
	ErodeIntensityMorphology  /* Pixel Pick using GreyScale Erode */
	DilateIntensityMorphology /* Pixel Pick using GreyScale Dialate */
	DistanceMorphology        /* Add Kernel Value, take Minimum */
	/* Second-level Morphology methods */
	OpenMorphology           /* Dilate then Erode */
	CloseMorphology          /* Erode then Dilate */
	OpenIntensityMorphology  /* Pixel Pick using GreyScale Open */
	CloseIntensityMorphology /* Pixel Pick using GreyScale Close */
	SmoothMorphology         /* Open then Close */
	/* Difference Morphology methods */
	EdgeInMorphology    /* Dilate difference from Original */
	EdgeOutMorphology   /* Erode difference from Original */
	EdgeMorphology      /* Dilate difference with Erode */
	TopHatMorphology    /* Close difference from Original */
	BottomHatMorphology /* Open difference from Original */
	/* Recursive Morphology methods */
	HitAndMissMorphology /* Foreground/Background pattern matching */
	ThinningMorphology   /* Remove matching pixels from image */
	ThickenMorphology    /* Add matching pixels from image */
	/* Experimental Morphology methods */
	VoronoiMorphology           /* distance matte channel copy nearest color */
	IterativeDistanceMorphology /* Add Kernel Value, take Minimum */
)

var morphologyMethodStrings = map[MorphologyMethod]string{
	UndefinedMorphology:         "UndefinedMorphology",
	ConvolveMorphology:          "ConvolveMorphology",
	CorrelateMorphology:         "CorrelateMorphology",
	ErodeMorphology:             "ErodeMorphology",
	DilateMorphology:            "DilateMorphology",
	ErodeIntensityMorphology:    "ErodeIntensityMorphology",
	DilateIntensityMorphology:   "DilateIntensityMorphology",
	DistanceMorphology:          "DistanceMorphology",
	OpenMorphology:              "OpenMorphology",
	CloseMorphology:             "CloseMorphology",
	OpenIntensityMorphology:     "OpenIntensityMorphology",
	CloseIntensityMorphology:    "CloseIntensityMorphology",
	SmoothMorphology:            "SmoothMorphology",
	EdgeInMorphology:            "EdgeInMorphology",
	EdgeOutMorphology:           "EdgeOutMorphology",
	EdgeMorphology:              "EdgeMorphology",
	TopHatMorphology:            "TopHatMorphology",
	BottomHatMorphology:         "BottomHatMorphology",
	HitAndMissMorphology:        "HitAndMissMorphology",
	ThinningMorphology:          "ThinningMorphology",
	ThickenMorphology:           "ThickenMorphology",
	VoronoiMorphology:           "VoronoiMorphology",
	IterativeDistanceMorphology: "IterativeDistanceMorphology",
}

func (cst *MorphologyMethod) String() string {
	if v, ok := morphologyMethodStrings[MorphologyMethod(*cst)]; ok {
		return v
	}
	return fmt.Sprintf("UnknownMorphologyMethod[%d]", *cst)
}
