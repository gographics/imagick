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
	UndefinedFilter     FilterType = C.UndefinedFilter
	PointFilter         FilterType = C.PointFilter
	BoxFilter           FilterType = C.BoxFilter
	TriangleFilter      FilterType = C.TriangleFilter
	HermiteFilter       FilterType = C.HermiteFilter
	HanningFilter       FilterType = C.HanningFilter
	HammingFilter       FilterType = C.HammingFilter
	BlackmanFilter      FilterType = C.BlackmanFilter
	GaussianFilter      FilterType = C.GaussianFilter
	QuadraticFilter     FilterType = C.QuadraticFilter
	CubicFilter         FilterType = C.CubicFilter
	CatromFilter        FilterType = C.CatromFilter
	MitchellFilter      FilterType = C.MitchellFilter
	JincFilter          FilterType = C.JincFilter
	SincFilter          FilterType = C.SincFilter
	SincFastFilter      FilterType = C.SincFastFilter
	KaiserFilter        FilterType = C.KaiserFilter
	WelshFilter         FilterType = C.WelshFilter
	ParzenFilter        FilterType = C.ParzenFilter
	BohmanFilter        FilterType = C.BohmanFilter
	BartlettFilter      FilterType = C.BartlettFilter
	LagrangeFilter      FilterType = C.LagrangeFilter
	LanczosFilter       FilterType = C.LanczosFilter
	LanczosSharpFilter  FilterType = C.LanczosSharpFilter
	Lanczos2Filter      FilterType = C.Lanczos2Filter
	Lanczos2SharpFilter FilterType = C.Lanczos2SharpFilter
	RobidouxFilter      FilterType = C.RobidouxFilter
	RobidouxSharpFilter FilterType = C.RobidouxSharpFilter
	CosineFilter        FilterType = C.CosineFilter
	SplineFilter        FilterType = C.SplineFilter
	LanczosRadiusFilter FilterType = C.LanczosRadiusFilter
	SentinelFilter      FilterType = C.SentinelFilter
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
