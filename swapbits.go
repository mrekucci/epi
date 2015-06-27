// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package epi

// SwapBits swaps the bits of x at indices i and j, and returns the result.
func SwapBits(x uint64, i, j uint64) uint64 {
	if (x >> i & 1) != (x >> j & 1) {
		x ^= (1<<i | 1<<j)
	}
	return x
}
