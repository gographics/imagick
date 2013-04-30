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
	UndefinedEvaluateOperator EvaluateOperator = iota
	AddEvaluateOperator
	AndEvaluateOperator
	DivideEvaluateOperator
	LeftShiftEvaluateOperator
	MaxEvaluateOperator
	MinEvaluateOperator
	MultiplyEvaluateOperator
	OrEvaluateOperator
	RightShiftEvaluateOperator
	SetEvaluateOperator
	SubtractEvaluateOperator
	XorEvaluateOperator
	PowEvaluateOperator
	LogEvaluateOperator
	ThresholdEvaluateOperator
	ThresholdBlackEvaluateOperator
	ThresholdWhiteEvaluateOperator
	GaussianNoiseEvaluateOperator
	ImpulseNoiseEvaluateOperator
	LaplacianNoiseEvaluateOperator
	MultiplicativeNoiseEvaluateOperator
	PoissonNoiseEvaluateOperator
	UniformNoiseEvaluateOperator
	CosineEvaluateOperator
	SineEvaluateOperator
	AddModulusEvaluateOperator
	MeanEvaluateOperator
	AbsEvaluateOperator
	ExponentialEvaluateOperator
	MedianEvaluateOperator
	SumEvaluateOperator
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
