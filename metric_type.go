package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
)

type MetricType int

const (
	UndefinedMetric                       MetricType = C.UndefinedMetric
	AbsoluteErrorMetric                   MetricType = C.AbsoluteErrorMetric
	MeanAbsoluteErrorMetric               MetricType = C.MeanAbsoluteErrorMetric
	MeanErrorPerPixelMetric               MetricType = C.MeanErrorPerPixelMetric
	MeanSquaredErrorMetric                MetricType = C.MeanSquaredErrorMetric
	PeakAbsoluteErrorMetric               MetricType = C.PeakAbsoluteErrorMetric
	PeakSignalToNoiseRatioMetric          MetricType = C.PeakSignalToNoiseRatioMetric
	RootMeanSquaredErrorMetric            MetricType = C.RootMeanSquaredErrorMetric
	NormalizedCrossCorrelationErrorMetric MetricType = C.NormalizedCrossCorrelationErrorMetric
	FuzzErrorMetric                       MetricType = C.FuzzErrorMetric
)

var metricTypeStrings = map[MetricType]string{
	UndefinedMetric:                       "UndefinedMetric",
	AbsoluteErrorMetric:                   "AbsoluteErrorMetric",
	MeanAbsoluteErrorMetric:               "MeanAbsoluteErrorMetric",
	MeanErrorPerPixelMetric:               "MeanErrorPerPixelMetric",
	MeanSquaredErrorMetric:                "MeanSquaredErrorMetric",
	PeakAbsoluteErrorMetric:               "PeakAbsoluteErrorMetric",
	PeakSignalToNoiseRatioMetric:          "PeakSignalToNoiseRatioMetric",
	RootMeanSquaredErrorMetric:            "RootMeanSquaredErrorMetric",
	NormalizedCrossCorrelationErrorMetric: "NormalizedCrossCorrelationErrorMetric",
	FuzzErrorMetric:                       "FuzzErrorMetric",
}

func (cst *MetricType) String() string {
	if v, ok := metricTypeStrings[MetricType(*cst)]; ok {
		return v
	}
	return fmt.Sprintf("UnknownNoise[%d]", *cst)
}
