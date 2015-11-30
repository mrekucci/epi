// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package graphs

const (
	white = true
	black = false
)

// Coordinate represents concrete position in 2D space.
type Coordinate struct {
	x, y int
}

// 2D coordinate shift by one.
var (
	right = Coordinate{0, 1}
	left  = Coordinate{0, -1}
	down  = Coordinate{1, 0}
	up    = Coordinate{-1, 0}
)

