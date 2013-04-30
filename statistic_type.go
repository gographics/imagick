package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
)

type StatisticType int

const (
	UndefinedStatistic StatisticType = iota
	GradientStatistic
	MaximumStatistic
	MeanStatistic
	MedianStatistic
	MinimumStatistic
	ModeStatistic
	NonpeakStatistic
	StandardDeviationStatistic
)

var statisticTypeStrings = map[StatisticType]string{
	UndefinedStatistic:         "UndefinedStatistic",
	GradientStatistic:          "GradientStatistic",
	MaximumStatistic:           "MaximumStatistic",
	MeanStatistic:              "MeanStatistic",
	MedianStatistic:            "MedianStatistic",
	MinimumStatistic:           "MinimumStatistic",
	ModeStatistic:              "ModeStatistic",
	NonpeakStatistic:           "NonpeakStatistic",
	StandardDeviationStatistic: "StandardDeviationStatistic",
}

func (ct *StatisticType) String() string {
	if v, ok := statisticTypeStrings[StatisticType(*ct)]; ok {
		return v
	}
	return fmt.Sprintf("UnknownCompression[%d]", *ct)
}
