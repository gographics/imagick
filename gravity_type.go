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
	UndefinedGravity GravityType = C.UndefinedGravity
	ForgetGravity    GravityType = C.ForgetGravity
	NorthWestGravity GravityType = C.NorthWestGravity
	NorthGravity     GravityType = C.NorthGravity
	NorthEastGravity GravityType = C.NorthEastGravity
	WestGravity      GravityType = C.WestGravity
	CenterGravity    GravityType = C.CenterGravity
	EastGravity      GravityType = C.EastGravity
	SouthWestGravity GravityType = C.SouthWestGravity
	SouthGravity     GravityType = C.SouthGravity
	SouthEastGravity GravityType = C.SouthEastGravity
	StaticGravity    GravityType = C.StaticGravity
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
