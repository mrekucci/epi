package ch5

import "math"

// reverseBits reverse bits in x.
func reverseBits(x uint16) uint16 {
	rx := uint64(x)
	i := uint64(0)
	for i < 16/2 {
		rx = swapBits(rx, i, (15 - i))
		i++
	}
	return uint16(rx)
}

// reverseBitsTable is a reverse bits cache for all 16-bit nonnegative integers.
var reverseBitsTable = initReverseBitsTable()

// initReverseBitsTable computes and returns reverse
// bits for all 16-bit nonnegative integers.
func initReverseBitsTable() []uint16 {
	bt := make([]uint16, 1<<16)
	for i := 0; i <= math.MaxUint16; i++ {
		bt[i] = reverseBits(uint16(i))
	}
	return bt
}

func reverseBitsLookup(x uint64) uint64 {
	return uint64(reverseBitsTable[(x>>48)&0xffff]) |
		uint64(reverseBitsTable[(x>>32)&0xffff])<<16 |
		uint64(reverseBitsTable[(x>>16)&0xffff])<<32 |
		uint64(reverseBitsTable[x&0xffff])<<48
}
