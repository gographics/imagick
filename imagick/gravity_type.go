// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type GravityType int

const (
	GRAVITY_UNDEFINED  GravityType = C.UndefinedGravity
	GRAVITY_FORGET     GravityType = C.ForgetGravity
	GRAVITY_NORTH_WEST GravityType = C.NorthWestGravity
	GRAVITY_NORTH      GravityType = C.NorthGravity
	GRAVITY_NORTH_EAST GravityType = C.NorthEastGravity
	GRAVITY_WEST       GravityType = C.WestGravity
	GRAVITY_CENTER     GravityType = C.CenterGravity
	GRAVITY_EAST       GravityType = C.EastGravity
	GRAVITY_SOUTH_WEST GravityType = C.SouthWestGravity
	GRAVITY_SOUTH      GravityType = C.SouthGravity
	GRAVITY_SOUTH_EAST GravityType = C.SouthEastGravity
)
