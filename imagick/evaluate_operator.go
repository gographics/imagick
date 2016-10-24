// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type EvaluateOperator int

const (
	EVAL_OP_UNDEFINED            EvaluateOperator = C.UndefinedEvaluateOperator
	EVAL_OP_ADD                  EvaluateOperator = C.AddEvaluateOperator
	EVAL_OP_AND                  EvaluateOperator = C.AndEvaluateOperator
	EVAL_OP_DIVIDE               EvaluateOperator = C.DivideEvaluateOperator
	EVAL_OP_LEFT_SHIFT           EvaluateOperator = C.LeftShiftEvaluateOperator
	EVAL_OP_MAX                  EvaluateOperator = C.MaxEvaluateOperator
	EVAL_OP_MIN                  EvaluateOperator = C.MinEvaluateOperator
	EVAL_OP_MULTIPLY             EvaluateOperator = C.MultiplyEvaluateOperator
	EVAL_OP_OR                   EvaluateOperator = C.OrEvaluateOperator
	EVAL_OP_RIGHT_SHIFT          EvaluateOperator = C.RightShiftEvaluateOperator
	EVAL_OP_SET                  EvaluateOperator = C.SetEvaluateOperator
	EVAL_OP_SUBTRACT             EvaluateOperator = C.SubtractEvaluateOperator
	EVAL_OP_XOR                  EvaluateOperator = C.XorEvaluateOperator
	EVAL_OP_POW                  EvaluateOperator = C.PowEvaluateOperator
	EVAL_OP_LOG                  EvaluateOperator = C.LogEvaluateOperator
	EVAL_OP_THRESHOLD            EvaluateOperator = C.ThresholdEvaluateOperator
	EVAL_OP_THRESHOLD_BLACK      EvaluateOperator = C.ThresholdBlackEvaluateOperator
	EVAL_OP_THRESHOLD_WHITE      EvaluateOperator = C.ThresholdWhiteEvaluateOperator
	EVAL_OP_GAUSSIAN_NOISE       EvaluateOperator = C.GaussianNoiseEvaluateOperator
	EVAL_OP_IMPULSE_NOISE        EvaluateOperator = C.ImpulseNoiseEvaluateOperator
	EVAL_OP_LAPLACIAN_NOISE      EvaluateOperator = C.LaplacianNoiseEvaluateOperator
	EVAL_OP_MULTIPLICATIVE_NOISE EvaluateOperator = C.MultiplicativeNoiseEvaluateOperator
	EVAL_OP_POISSON_NOISE        EvaluateOperator = C.PoissonNoiseEvaluateOperator
	EVAL_OP_UNIFORM_NOISE        EvaluateOperator = C.UniformNoiseEvaluateOperator
	EVAL_OP_COSINE               EvaluateOperator = C.CosineEvaluateOperator
	EVAL_OP_SINE                 EvaluateOperator = C.SineEvaluateOperator
	EVAL_OP_ADD_MODULUS          EvaluateOperator = C.AddModulusEvaluateOperator
	EVAL_OP_MEAN                 EvaluateOperator = C.MeanEvaluateOperator
	EVAL_OP_ABS                  EvaluateOperator = C.AbsEvaluateOperator
	EVAL_OP_EXPONENTIAL          EvaluateOperator = C.ExponentialEvaluateOperator
	EVAL_OP_MEDIAN               EvaluateOperator = C.MedianEvaluateOperator
	EVAL_OP_SUM                  EvaluateOperator = C.SumEvaluateOperator
)
