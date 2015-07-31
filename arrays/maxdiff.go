// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package arrays

// Integer limit values.
const (
	maxInt = int(^uint(0) >> 1)
	minInt = -maxInt - 1
)

// max returns the larger of x or y.
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// min returns the smaller of x or y.
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// MinBatteryCap returns battery capacity needed
// for given heights to be traversed.
//
// This algorithm finds the biggest difference
// (in ASC order) between two points in O(n) time.
func MinBatteryCap(heights []int) (int, bool) {
	if len(heights) == 0 {
		return 0, true
	}

	cap := 0
	hm := heights[0] // We don't set it to maxInt, 'cause we wanna avoid to return false overflow indication when 1st element is negative.

	// We assume that z is the vertical dimension (height). Energy usage depends on the change in height.
	for _, h := range heights {
		if (h > 0 && -hm > maxInt-h) || (h < 0 && -hm < minInt-h) { // Check for overflow.
			return 0, false
		}
		cap = max(cap, h-hm)
		hm = min(hm, h)
	}

	return cap, true
}
