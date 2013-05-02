package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
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
