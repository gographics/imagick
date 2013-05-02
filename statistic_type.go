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
	UndefinedStatistic         StatisticType = C.UndefinedStatistic
	GradientStatistic          StatisticType = C.GradientStatistic
	MaximumStatistic           StatisticType = C.MaximumStatistic
	MeanStatistic              StatisticType = C.MeanStatistic
	MedianStatistic            StatisticType = C.MedianStatistic
	MinimumStatistic           StatisticType = C.MinimumStatistic
	ModeStatistic              StatisticType = C.ModeStatistic
	NonpeakStatistic           StatisticType = C.NonpeakStatistic
	StandardDeviationStatistic StatisticType = C.StandardDeviationStatistic
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
