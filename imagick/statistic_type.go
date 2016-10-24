// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type StatisticType int

const (
	STATISTIC_UNDEFINED          StatisticType = C.UndefinedStatistic
	STATISTIC_GRADIENT           StatisticType = C.GradientStatistic
	STATISTIC_MAXIMUM            StatisticType = C.MaximumStatistic
	STATISTIC_MEAN               StatisticType = C.MeanStatistic
	STATISTIC_MEDIAN             StatisticType = C.MedianStatistic
	STATISTIC_MINIMUM            StatisticType = C.MinimumStatistic
	STATISTIC_MODE               StatisticType = C.ModeStatistic
	STATISTIC_NONPEAK            StatisticType = C.NonpeakStatistic
	STATISTIC_STANDARD_DEVIATION StatisticType = C.StandardDeviationStatistic
)
