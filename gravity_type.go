package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
)

type GravityType int

const (
	UndefinedGravity GravityType = iota
)

const (
	ForgetGravity GravityType = iota
	NorthWestGravity
	NorthGravity
	NorthEastGravity
	WestGravity
	CenterGravity
	EastGravity
	SouthWestGravity
	SouthGravity
	SouthEastGravity
	StaticGravity
)

var gravityTypeStrings = map[GravityType]string{
	UndefinedGravity: "UndefinedGravity",
	NorthWestGravity: "NorthWestGravity",
	NorthGravity:     "NorthGravity",
	NorthEastGravity: "NorthEastGravity",
	WestGravity:      "WestGravity",
	CenterGravity:    "CenterGravity",
	EastGravity:      "EastGravity",
	SouthWestGravity: "SouthWestGravity",
	SouthGravity:     "SouthGravity",
	SouthEastGravity: "SouthEastGravity",
	StaticGravity:    "StaticGravity",
}

func (cst *GravityType) String() string {
	if v, ok := gravityTypeStrings[GravityType(*cst)]; ok {
		return v
	}
	return fmt.Sprintf("UnknownGravityType[%d]", *cst)
}
