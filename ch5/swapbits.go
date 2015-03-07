package ch5

// swapBits swaps the bits of x at indices i and j, and returns the result.
func swapBits(x uint64, i, j uint64) uint64 {
	if (x >> i & 1) != (x >> j & 1) {
		x ^= (1<<i | 1<<j)
	}
	return x
}
