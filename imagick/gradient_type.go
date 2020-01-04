package imagick

/*
#include <MagickCore/MagickCore.h>
*/
import "C"

type GradientType int

const (
	GRADIENT_TYPE_UNDEFINED GradientType = C.UndefinedGradient
	GRADIENT_TYPE_LINEAR    GradientType = C.LinearGradient
	GRADIENT_TYPE_RADIAL    GradientType = C.RadialGradient
)

// StopInfo describes the color and offset of a
// color stop component in a gradient
type StopInfo struct {
	info C.StopInfo
}

// NewStopInfo creates a StopInfo from the color of a
// PixelInfo, and a stop offset value in the gradient.
// The offset value is from 0.0 to 1.0
func NewStopInfo(pi *PixelInfo, offset float64) StopInfo {
	var stop StopInfo
	stop.SetPixelInfo(pi)
	stop.SetOffset(offset)
	return stop
}

// SetPixelInfo sets the gradient stop color value
func (s *StopInfo) SetPixelInfo(pi *PixelInfo) {
	s.info.color = *(pi.pi)
}

// SetOffset sets the stop offset within the gradient,
// from 0.0 to 1.0
func (s *StopInfo) SetOffset(offset float64) {
	s.info.offset = C.double(offset)
}
