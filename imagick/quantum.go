package imagick

/*
#cgo pkg-config: MagickCore
#include <magick/MagickCore.h>
*/
import "C"

type Quantum C.Quantum

const QUANTUM_RANGE = C.QuantumRange
