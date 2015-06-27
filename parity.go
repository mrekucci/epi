// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package epi

import "math"

// parityTable is a parity cache for all 16-bit non-negative integers.
var parityTable = initParityTable()

// initParityTable computes and returns parities for all 16-bit non-negative integers.
func initParityTable() []uint16 {
	pt := make([]uint16, 1<<16)
	for i := 0; i <= math.MaxUint16; i++ {
		pt[i] = Parity(uint16(i))
	}
	return pt
}

// Parity returns 1 if the number of bits set to 1 in x is odd, otherwise O.
func Parity(x uint16) (res uint16) {
	for x > 0 {
		res ^= 1
		x &= (x - 1)
	}
	return res
}

// ParityLookup returns 1 if the number of bits set to 1 in x is odd, otherwise O.
// This function is good for computing the parity of a very large number of
// 64-bit non-negative integers.
func ParityLookup(x uint64) uint16 {
	return parityTable[(x>>48)&0xffff] ^
		parityTable[(x>>32)&0xffff] ^
		parityTable[(x>>16)&0xffff] ^
		parityTable[x&0xffff]
}
