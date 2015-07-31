// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package ptypes

import "math"

// reverseBitsTable is a reverse bits cache for all 16-bit non-negative integers.
var reverseBitsTable = initReverseBitsTable()

// initReverseBitsTable computes and returns reverse
// bits for all 16-bit non-negative integers.
func initReverseBitsTable() []uint16 {
	bt := make([]uint16, 1<<16)
	for i := 0; i <= math.MaxUint16; i++ {
		bt[i] = ReverseBits(uint16(i))
	}
	return bt
}

// ReverseBits reverse bits in x.
func ReverseBits(x uint16) uint16 {
	rx := uint64(x)
	i := uint64(0)
	for i < 16/2 {
		rx = SwapBits(rx, i, (15 - i))
		i++
	}
	return uint16(rx)
}

// ReverseBitsLookup reverse bits in x.
// This function is good for reverse bits of a very
// large number of 64-bit non-negative integers.
func ReverseBitsLookup(x uint64) uint64 {
	return uint64(reverseBitsTable[(x>>48)&0xffff]) |
		uint64(reverseBitsTable[(x>>32)&0xffff])<<16 |
		uint64(reverseBitsTable[(x>>16)&0xffff])<<32 |
		uint64(reverseBitsTable[x&0xffff])<<48
}
