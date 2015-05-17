package ch5

import "math"

// parity returns 1 if the number of bits set to 1 in x is odd, otherwise O.
func parity(x uint16) (res uint16) {
	for x > 0 {
		res ^= 1
		x &= (x - 1)
	}
	return res
}

// parityTable is a parity cache for all 16-bit non-negative integers.
var parityTable = initParityTable()

// initParityTable computes and returns parities for all 16-bit non-negative integers.
func initParityTable() []uint16 {
	pt := make([]uint16, 1<<16)
	for i := 0; i <= math.MaxUint16; i++ {
		pt[i] = parity(uint16(i))
	}
	return pt
}

// parityLookup returns 1 if the number of bits set to 1 in x is odd, otherwise O.
// This function is good for computing the parity of a very large number of
// 64-bit non-negative integers.
func parityLookup(x uint64) uint16 {
	return parityTable[(x>>48)&0xffff] ^
		parityTable[(x>>32)&0xffff] ^
		parityTable[(x>>16)&0xffff] ^
		parityTable[x&0xffff]
}
