package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type DirectionType int

const (
	DIRECTION_LEFT_TO_RIGHT	DirectionType = C.LeftToRightDirection
	DIRECTION_RIGHT_TO_LEFT DirectionType = C.RightToLeftDirection
)
