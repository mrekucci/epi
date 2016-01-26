// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package ptypes

import "math"

// ReverseInt returns reversed x digits.
// 0, false is returned if x cannot be reversed because of int64 overflow.
// The time complexity is O(n), where n is the number of digits in x.
// The space complexity is O(1).
func ReverseInt(x int64) (r int64, ok bool) {
	const cutoff = math.MaxInt64/10 + 1 // The first smallest number such that cutoff*10 > MaxInt64.
	var n uint64

	neg := x < 0
	u := uint64(x)
	if neg {
		u = -u
	}

	for u > 0 {
		if n >= cutoff { // Check if n*10 overflows.
			return 0, false // TODO: cover this in tests!
		}
		n = n*10 + u%10
		if neg && n > -math.MinInt64 || !neg && n > math.MaxInt64 { // -n < math.MinInt64 || n > math.MaxInt64
			return 0, false
		}
		u /= 10
	}

	r = int64(n)
	if neg {
		r = -r
	}
	return r, true
}
