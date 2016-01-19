// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package bsearch

import (
	"math"
)

// cmp compares a and b with a tolerance and returns:
//
// 	-1: a < b
//	 0: a == b
//	 1: a > b
//
// Note: works only for positive numbers.
func cmp(a, b float64) int {
	const epsilon = 0.00001       // The tolerance.
	switch diff := (a - b) / b; { // a-b ensures that the diff will be always negative when a < b
	case diff < -epsilon: // Check tolerance if a < b
		return -1
	case diff > epsilon: // Check tolerance if a > b
		return 1
	}
	return 0
}

// SquareRootReal returns square root of x.
// The time complexity is O(x/t), where t is the tolerance.
// The space complexity is O(1).
func SquareRootReal(x float64) float64 {
	switch {
	case x < 0:
		return math.NaN()
	case x == 0 || math.IsNaN(x) || math.IsInf(x, 1):
		return x
	}

	l, r := 1.0, x // For x >= 1.0, x <= x*x, so, we search in interval: <1.0, x>.
	if x < 1.0 {   // For x < 1.0, x > x*x, so, we search in interval: <x, 1.0>.
		l, r = r, l
	}

	for cmp(l, r) == -1 {
		m := l + (r-l)/2
		switch cmp(m*m, x) {
		case 0:
			return m
		case 1:
			r = m
		case -1:
			l = m
		}
	}

	return l
}
