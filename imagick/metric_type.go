// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type MetricType int

const (
	METRIC_UNDEFINED                          MetricType = C.UndefinedErrorMetric
	METRIC_ABSOLUTE_ERROR                     MetricType = C.AbsoluteErrorMetric
	METRIC_FUZZ_ERROR                         MetricType = C.FuzzErrorMetric
	METRIC_MEAN_ABSOLUTE_ERROR                MetricType = C.MeanAbsoluteErrorMetric
	METRIC_MEAN_ERROR_PER_PIXEL               MetricType = C.MeanErrorPerPixelErrorMetric
	METRIC_MEAN_SQUARED_ERROR                 MetricType = C.MeanSquaredErrorMetric
	METRIC_NORMALIZED_CROSS_CORRELATION_ERROR MetricType = C.NormalizedCrossCorrelationErrorMetric
	METRIC_PEAK_ABSOLUTE_ERROR                MetricType = C.PeakAbsoluteErrorMetric
	METRIC_PEAK_SIGNAL_TO_NOISE_RATIO         MetricType = C.PeakSignalToNoiseRatioErrorMetric
	METRIC_PERCEPTUAL_HASH_ERROR              MetricType = C.PerceptualHashErrorMetric
	METRIC_ROOT_MEAN_SQUARED_ERROR            MetricType = C.RootMeanSquaredErrorMetric
)
