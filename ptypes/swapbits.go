// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package ptypes

// SwapBits swaps the bits of x at indices i and j, and returns the result.
// The time and space complexity is O(1), independent of word size.
func SwapBits(x, i, j uint64) uint64 {
	if (x >> i & 1) != (x >> j & 1) {
		x ^= (1<<i | 1<<j)
	}
	return x
}
