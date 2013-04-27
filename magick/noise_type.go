package magick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
)

type NoiseType int

const (
	UndefinedNoise NoiseType = iota
	UniformNoise
	GaussianNoise
	MultiplicativeGaussianNoise
	ImpulseNoise
	LaplacianNoise
	PoissonNoise
	RandomNoise
)

var noiseTypeStrings = map[NoiseType]string{
	UndefinedNoise:              "UndefinedNoise",
	UniformNoise:                "UniformNoise",
	GaussianNoise:               "GaussianNoise",
	MultiplicativeGaussianNoise: "MultiplicativeGaussianNoise",
	ImpulseNoise:                "ImpulseNoise",
	LaplacianNoise:              "LaplacianNoise",
	PoissonNoise:                "PoissonNoise",
	RandomNoise:                 "RandomNoise",
}

func (cst *NoiseType) String() string {
	if v, ok := noiseTypeStrings[NoiseType(*cst)]; ok {
		return v
	}
	return fmt.Sprintf("UnknownNoise[%d]", *cst)
}
