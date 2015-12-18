// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package ptypes

// parityTable is a parity cache for all 16-bit non-negative integers.
var parityTable = initParityTable()

// initParityTable computes and returns parities for all 16-bit non-negative integers.
func initParityTable() []uint16 {
	pt := make([]uint16, 1<<16)
	for i := 0; i < len(pt); i++ {
		pt[i] = Parity(int64(i))
	}
	return pt
}

// Parity returns 1 if the number of bits set to 1 in x is odd, otherwise O.
// The time complexity is O(log(n)) where n is the word size.
// The space complexity is O(1).
func Parity(x int64) (p uint16) {
	x ^= x >> 32
	x ^= x >> 16
	x ^= x >> 8
	x ^= x >> 4
	x ^= x >> 2
	x ^= x >> 1
	return uint16(x & 1)
}

// Parity returns 1 if the number of bits set to 1 in x is odd, otherwise O.
// The time complexity is O(k) where k is the number of bits in x set to 1.
// The space complexity is O(1).
func ParityAlt(x int64) (p uint16) {
	for x != 0 {
		p ^= 1
		x &= (x - 1)
	}
	return p
}

// ParityLookup returns 1 if the number of bits set to 1 in x is odd, otherwise O.
// The time complexity is O(n/l) where n is the word size and l is the width of a
// world for which we have cached results. The space complexity is O(1) beyond the
// 1<<l space needed to cache results which is constant.
func ParityLookup(x int64) uint16 {
	return parityTable[(x>>48)&0xffff] ^ parityTable[(x>>32)&0xffff] ^ parityTable[(x>>16)&0xffff] ^ parityTable[x&0xffff]
}
