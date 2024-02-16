package imagick

/*
#include <MagickCore/MagickCore.h>
*/
import "C"

type GeometryInfo struct {
	Rho,
	Sigma,
	Xi,
	Psi,
	Chi float64
}

func newGeometryInfo(gi *C.GeometryInfo) *GeometryInfo {
	if gi == nil {
		return nil
	}
	return &GeometryInfo{
		Rho:   float64(gi.rho),
		Sigma: float64(gi.sigma),
		Xi:    float64(gi.xi),
		Psi:   float64(gi.psi),
		Chi:   float64(gi.chi),
	}
}
