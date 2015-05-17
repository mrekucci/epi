package ch5

import "math"

// PowerSet returns a power set of s.
// The length of s is limited to size of int.
// When is crossed then nil interface and false is returned.
func PowerSet(s []interface{}) ([]interface{}, bool) {
	var ps []interface{}
	if len(s) >= (32 << (^uint(0) >> 63)) {
		return ps, false
	}

	for i := 0; i < (1 << uint(len(s))); i++ {
		x := i
		var ss []interface{}
		for x > 0 {
			bp := int(math.Log2(float64(x & ^(x - 1)))) // Compute least significant bit position of int.
			ss = append(ss, s[bp])
			x &= (x - 1) // This ensure that the iteration count will be the same as number of 1 bits in x.
			// Add if statement x == 0 which will execute on sub-set end.
		}
		ps = append(ps, ss)
	}
	return ps, true
}
