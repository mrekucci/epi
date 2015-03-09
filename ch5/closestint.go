package ch5

// intWeight returns number of bits set to 1 in x.
func intWeight(x uint64) (w int) {
	for x > 0 {
		w++
		x &= (x - 1)
	}
	return w
}

// closestInt returns integer closest to x with the same weight.
func closestInt(x uint64) (uint64, bool) {
	for i := uint(0); i < 63; i++ {
		if (x>>i)&1 != (x >> (i + 1) & 1) {
			x ^= 1<<i | 1<<(i+1)
			return x, true
		}
	}
	return 0, false // If all bits are 0 or 1.
}
