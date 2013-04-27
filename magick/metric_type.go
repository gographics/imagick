package magick

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
	UndefinedMetric MetricType = iota
	AbsoluteErrorMetric
	MeanAbsoluteErrorMetric
	MeanErrorPerPixelMetric
	MeanSquaredErrorMetric
	PeakAbsoluteErrorMetric
	PeakSignalToNoiseRatioMetric
	RootMeanSquaredErrorMetric
	NormalizedCrossCorrelationErrorMetric
	FuzzErrorMetric
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
