package epi

// intSize is the size in bits of an int.
const intSize = 32 << (^uint(0) >> 63)

// Integer limit values.
const (
	maxInt = int(^uint(0) >> 1)
	minInt = -maxInt - 1
)
