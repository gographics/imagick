package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
)

type FilterType int

const (
	UndefinedFilter FilterType = iota
	PointFilter
	BoxFilter
	TriangleFilter
	HermiteFilter
	HanningFilter
	HammingFilter
	BlackmanFilter
	GaussianFilter
	QuadraticFilter
	CubicFilter
	CatromFilter
	MitchellFilter
	JincFilter
	SincFilter
	SincFastFilter
	KaiserFilter
	WelshFilter
	ParzenFilter
	BohmanFilter
	BartlettFilter
	LagrangeFilter
	LanczosFilter
	LanczosSharpFilter
	Lanczos2Filter
	Lanczos2SharpFilter
	RobidouxFilter
	RobidouxSharpFilter
	CosineFilter
	SplineFilter
	LanczosRadiusFilter
	SentinelFilter // a count of all the filters, not a real filter
)

var filterTypeStrings = map[FilterType]string{
	UndefinedFilter:     "UndefinedFilter",
	PointFilter:         "PointFilter",
	BoxFilter:           "BoxFilter",
	TriangleFilter:      "TriangleFilter",
	HermiteFilter:       "HermiteFilter",
	HanningFilter:       "HanningFilter",
	HammingFilter:       "HammingFilter",
	BlackmanFilter:      "BlackmanFilter",
	GaussianFilter:      "GaussianFilter",
	QuadraticFilter:     "QuadraticFilter",
	CubicFilter:         "CubicFilter",
	CatromFilter:        "CatromFilter",
	MitchellFilter:      "MitchellFilter",
	JincFilter:          "JincFilter",
	SincFilter:          "SincFilter",
	SincFastFilter:      "SincFastFilter",
	KaiserFilter:        "KaiserFilter",
	WelshFilter:         "WelshFilter",
	ParzenFilter:        "ParzenFilter",
	BohmanFilter:        "BohmanFilter",
	BartlettFilter:      "BartlettFilter",
	LagrangeFilter:      "LagrangeFilter",
	LanczosFilter:       "LanczosFilter",
	LanczosSharpFilter:  "LanczosSharpFilter",
	Lanczos2Filter:      "Lanczos2Filter",
	Lanczos2SharpFilter: "Lanczos2SharpFilter",
	RobidouxFilter:      "RobidouxFilter",
	RobidouxSharpFilter: "RobidouxSharpFilter",
	CosineFilter:        "CosineFilter",
	SplineFilter:        "SplineFilter",
	LanczosRadiusFilter: "LanczosRadiusFilter",
	SentinelFilter:      "SentinelFilter",
}

func (cst *FilterType) String() string {
	if v, ok := filterTypeStrings[FilterType(*cst)]; ok {
		return v
	}
	return fmt.Sprintf("UnknownFilterType[%d]", *cst)
}
