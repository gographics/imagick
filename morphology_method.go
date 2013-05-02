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
	UndefinedMorphology         MorphologyMethod = C.UndefinedMorphology
	ConvolveMorphology          MorphologyMethod = C.ConvolveMorphology          /* Weighted Sum with reflected kernel */
	CorrelateMorphology         MorphologyMethod = C.CorrelateMorphology         /* Weighted Sum using a sliding window */
	ErodeMorphology             MorphologyMethod = C.ErodeMorphology             /* Minimum Value in Neighbourhood */
	DilateMorphology            MorphologyMethod = C.DilateMorphology            /* Maximum Value in Neighbourhood */
	ErodeIntensityMorphology    MorphologyMethod = C.ErodeIntensityMorphology    /* Pixel Pick using GreyScale Erode */
	DilateIntensityMorphology   MorphologyMethod = C.DilateIntensityMorphology   /* Pixel Pick using GreyScale Dialate */
	DistanceMorphology          MorphologyMethod = C.DistanceMorphology          /* Add Kernel Value, take Minimum */
	OpenMorphology              MorphologyMethod = C.OpenMorphology              /* Dilate then Erode */
	CloseMorphology             MorphologyMethod = C.CloseMorphology             /* Erode then Dilate */
	OpenIntensityMorphology     MorphologyMethod = C.OpenIntensityMorphology     /* Pixel Pick using GreyScale Open */
	CloseIntensityMorphology    MorphologyMethod = C.CloseIntensityMorphology    /* Pixel Pick using GreyScale Close */
	SmoothMorphology            MorphologyMethod = C.SmoothMorphology            /* Open then Close */
	EdgeInMorphology            MorphologyMethod = C.EdgeInMorphology            /* Dilate difference from Original */
	EdgeOutMorphology           MorphologyMethod = C.EdgeOutMorphology           /* Erode difference from Original */
	EdgeMorphology              MorphologyMethod = C.EdgeMorphology              /* Dilate difference with Erode */
	TopHatMorphology            MorphologyMethod = C.TopHatMorphology            /* Close difference from Original */
	BottomHatMorphology         MorphologyMethod = C.BottomHatMorphology         /* Open difference from Original */
	HitAndMissMorphology        MorphologyMethod = C.HitAndMissMorphology        /* Foreground/Background pattern matching */
	ThinningMorphology          MorphologyMethod = C.ThinningMorphology          /* Remove matching pixels from image */
	ThickenMorphology           MorphologyMethod = C.ThickenMorphology           /* Add matching pixels from image */
	VoronoiMorphology           MorphologyMethod = C.VoronoiMorphology           /* distance matte channel copy nearest color */
	IterativeDistanceMorphology MorphologyMethod = C.IterativeDistanceMorphology /* Add Kernel Value, take Minimum */
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
