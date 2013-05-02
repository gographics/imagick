package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
)

type EvaluateOperator int

const (
	UndefinedEvaluateOperator           EvaluateOperator = C.UndefinedEvaluateOperator
	AddEvaluateOperator                 EvaluateOperator = C.AddEvaluateOperator
	AndEvaluateOperator                 EvaluateOperator = C.AndEvaluateOperator
	DivideEvaluateOperator              EvaluateOperator = C.DivideEvaluateOperator
	LeftShiftEvaluateOperator           EvaluateOperator = C.LeftShiftEvaluateOperator
	MaxEvaluateOperator                 EvaluateOperator = C.MaxEvaluateOperator
	MinEvaluateOperator                 EvaluateOperator = C.MinEvaluateOperator
	MultiplyEvaluateOperator            EvaluateOperator = C.MultiplyEvaluateOperator
	OrEvaluateOperator                  EvaluateOperator = C.OrEvaluateOperator
	RightShiftEvaluateOperator          EvaluateOperator = C.RightShiftEvaluateOperator
	SetEvaluateOperator                 EvaluateOperator = C.SetEvaluateOperator
	SubtractEvaluateOperator            EvaluateOperator = C.SubtractEvaluateOperator
	XorEvaluateOperator                 EvaluateOperator = C.XorEvaluateOperator
	PowEvaluateOperator                 EvaluateOperator = C.PowEvaluateOperator
	LogEvaluateOperator                 EvaluateOperator = C.LogEvaluateOperator
	ThresholdEvaluateOperator           EvaluateOperator = C.ThresholdEvaluateOperator
	ThresholdBlackEvaluateOperator      EvaluateOperator = C.ThresholdBlackEvaluateOperator
	ThresholdWhiteEvaluateOperator      EvaluateOperator = C.ThresholdWhiteEvaluateOperator
	GaussianNoiseEvaluateOperator       EvaluateOperator = C.GaussianNoiseEvaluateOperator
	ImpulseNoiseEvaluateOperator        EvaluateOperator = C.ImpulseNoiseEvaluateOperator
	LaplacianNoiseEvaluateOperator      EvaluateOperator = C.LaplacianNoiseEvaluateOperator
	MultiplicativeNoiseEvaluateOperator EvaluateOperator = C.MultiplicativeNoiseEvaluateOperator
	PoissonNoiseEvaluateOperator        EvaluateOperator = C.PoissonNoiseEvaluateOperator
	UniformNoiseEvaluateOperator        EvaluateOperator = C.UniformNoiseEvaluateOperator
	CosineEvaluateOperator              EvaluateOperator = C.CosineEvaluateOperator
	SineEvaluateOperator                EvaluateOperator = C.SineEvaluateOperator
	AddModulusEvaluateOperator          EvaluateOperator = C.AddModulusEvaluateOperator
	MeanEvaluateOperator                EvaluateOperator = C.MeanEvaluateOperator
	AbsEvaluateOperator                 EvaluateOperator = C.AbsEvaluateOperator
	ExponentialEvaluateOperator         EvaluateOperator = C.ExponentialEvaluateOperator
	MedianEvaluateOperator              EvaluateOperator = C.MedianEvaluateOperator
	SumEvaluateOperator                 EvaluateOperator = C.SumEvaluateOperator
)

var magickEvaluateOperatorStrings = map[EvaluateOperator]string{
	UndefinedEvaluateOperator:           "UndefinedEvaluateOperator",
	AddEvaluateOperator:                 "AddEvaluateOperator",
	AndEvaluateOperator:                 "AndEvaluateOperator",
	DivideEvaluateOperator:              "DivideEvaluateOperator",
	LeftShiftEvaluateOperator:           "LeftShiftEvaluateOperator",
	MaxEvaluateOperator:                 "MaxEvaluateOperator",
	MinEvaluateOperator:                 "MinEvaluateOperator",
	MultiplyEvaluateOperator:            "MultiplyEvaluateOperator",
	OrEvaluateOperator:                  "OrEvaluateOperator",
	RightShiftEvaluateOperator:          "RightShiftEvaluateOperator",
	SetEvaluateOperator:                 "SetEvaluateOperator",
	SubtractEvaluateOperator:            "SubtractEvaluateOperator",
	XorEvaluateOperator:                 "XorEvaluateOperator",
	PowEvaluateOperator:                 "PowEvaluateOperator",
	LogEvaluateOperator:                 "LogEvaluateOperator",
	ThresholdEvaluateOperator:           "ThresholdEvaluateOperator",
	ThresholdBlackEvaluateOperator:      "ThresholdBlackEvaluateOperator",
	ThresholdWhiteEvaluateOperator:      "ThresholdWhiteEvaluateOperator",
	GaussianNoiseEvaluateOperator:       "GaussianNoiseEvaluateOperator",
	ImpulseNoiseEvaluateOperator:        "ImpulseNoiseEvaluateOperator",
	LaplacianNoiseEvaluateOperator:      "LaplacianNoiseEvaluateOperator",
	MultiplicativeNoiseEvaluateOperator: "MultiplicativeNoiseEvaluateOperator",
	PoissonNoiseEvaluateOperator:        "PoissonNoiseEvaluateOperator",
	UniformNoiseEvaluateOperator:        "UniformNoiseEvaluateOperator",
	CosineEvaluateOperator:              "CosineEvaluateOperator",
	SineEvaluateOperator:                "SineEvaluateOperator",
	AddModulusEvaluateOperator:          "AddModulusEvaluateOperator",
	MeanEvaluateOperator:                "MeanEvaluateOperator",
	AbsEvaluateOperator:                 "AbsEvaluateOperator",
	ExponentialEvaluateOperator:         "ExponentialEvaluateOperator",
	MedianEvaluateOperator:              "MedianEvaluateOperator",
	SumEvaluateOperator:                 "SumEvaluateOperator",
}

func (cst *EvaluateOperator) String() string {
	if v, ok := magickEvaluateOperatorStrings[EvaluateOperator(*cst)]; ok {
		return v
	}
	return fmt.Sprintf("UnknownEvaluateOperator[%d]", *cst)
}
