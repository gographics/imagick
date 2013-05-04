package imagick

/*
#cgo pkg-config: MagickCore
#include <magick/MagickCore.h>
*/
import "C"

type Quantum uint16

const QUANTUM_RANGE = C.QuantumRange
