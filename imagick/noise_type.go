// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type NoiseType int

const (
	NOISE_UNDEFINED               NoiseType = C.UndefinedNoise
	NOISE_UNIFORM                 NoiseType = C.UniformNoise
	NOISE_GAUSSIAN                NoiseType = C.GaussianNoise
	NOISE_MULTIPLICATIVE_GAUSSIAN NoiseType = C.MultiplicativeGaussianNoise
	NOISE_IMPULSE                 NoiseType = C.ImpulseNoise
	NOISE_LAPLACIAN               NoiseType = C.LaplacianNoise
	NOISE_POISSON                 NoiseType = C.PoissonNoise
	NOISE_RANDOM                  NoiseType = C.RandomNoise
)
