// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package ptypes

// reverseBitsTable is a reverse bits cache for all 16-bit non-negative integers.
var reverseBitsTable = createReverseBitsTable()

// createReverseBitsTable computes and returns reverse
// bits for all 16-bit non-negative integers.
func createReverseBitsTable() (bt [1 << 16]uint16) {
	for i := 0; i < len(bt); i++ {
		bt[i] = uint16(ReverseBits(uint64(i)) >> 48)
	}
	return bt
}

// ReverseBits reverse bits in x.
// The time complexity is O(n) where n is the word size.
// The space complexity is O(1).
func ReverseBits(x uint64) uint64 {
	for i := uint64(0); i < 64/2; i++ {
		x = SwapBits(x, i, 63-i)
	}
	return x
}

// ReverseBitsLookup reverse bits in x.
// The time complexity is O(n/l) where n is the word size and l is the width
// of a word of cache key. The space complexity is O(1) beyond the 1<<l space
// is needed to cache precomputed results, which is constant.
func ReverseBitsLookup(x uint64) uint64 {
	return uint64(reverseBitsTable[(x>>48)&0xffff]) |
		uint64(reverseBitsTable[(x>>32)&0xffff])<<16 |
		uint64(reverseBitsTable[(x>>16)&0xffff])<<32 |
		uint64(reverseBitsTable[x&0xffff])<<48
}
