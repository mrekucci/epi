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
		var ss []interface{}
		// x == 0 indicates sub-set end.
		// x &= (x - 1) ensures that the iteration count will be the same as number of bits set to 1 in x.
		for x := i; x > 0; x &= (x - 1) {
			lsb, i := x&-x, 0                    // x&-x is same as x&^(x - 1).
			for p := 1; lsb&p == 0; p = p << 1 { // lsb must always be greater then 0, which is always true 'cause x > 0.
				i++ // Compute the index of x's least significant bit.
			}
			ss = append(ss, s[i])
		}
		ps = append(ps, ss)
	}

	return ps, true
}
