package imagick

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
	UndefinedNoise              NoiseType = C.UndefinedNoise
	UniformNoise                NoiseType = C.UniformNoise
	GaussianNoise               NoiseType = C.GaussianNoise
	MultiplicativeGaussianNoise NoiseType = C.MultiplicativeGaussianNoise
	ImpulseNoise                NoiseType = C.ImpulseNoise
	LaplacianNoise              NoiseType = C.LaplacianNoise
	PoissonNoise                NoiseType = C.PoissonNoise
	RandomNoise                 NoiseType = C.RandomNoise
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
