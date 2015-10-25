// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package recursion

// PowerSet returns a power set of s.
// The length of s must be less then size of int.
// If the size is equal or bigger, then nil interface and false is returned.
func PowerSet(s []interface{}) (ps []interface{}, ok bool) {
	if len(s) >= intSize {
		return ps, false
	}

	for i := 0; i < (1 << uint(len(s))); i++ {
		x := i
		var ss []interface{}
		for x > 0 {
			lsb := x & -x // x & -x is same as x & ^(x - 1).

			// Compute the index of x's least significant bit.
			i := 0
			p := 1
			for lsb&p == 0 { // lsb must always be greater then 0, which is always true 'cause x > 0.
				p <<= 1
				i++
			}

			ss = append(ss, s[i])
			x &= (x - 1) // This ensures that the iteration count will be the same as number of 1 bits in x.
			// x == 0 indicates sub-set end.
		}
		ps = append(ps, ss)
	}

	return ps, true
}
