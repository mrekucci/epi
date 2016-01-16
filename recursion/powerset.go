// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package recursion

// genPowerSet generates a power set of s[i:] into the ps.
func genPowerSet(i int, s, branch []interface{}, ps [][]interface{}) [][]interface{} {
	if i == len(s) {
		return append(ps, append([]interface{}(nil), branch...))
	}
	ps = genPowerSet(i+1, s, append(branch, s[i]), ps) // Generate all subsets that contain s[i].
	ps = genPowerSet(i+1, s, branch, ps)               // Generate all subsets that do not contain s[i].
	return ps
}

// PowerSetRec returns a power set of s (it uses recursion to generate the set).
// The time complexity is O(n*(2**n)). The space complexity is O(2**n)
// The returned boolean value is always true.
func PowerSetRec(s []interface{}) (ps [][]interface{}, ok bool) {
	return genPowerSet(0, s, nil, nil), true
}

// PowerSetItr returns a power set of s (it uses mapping to integer bits to generate the set).
// The time complexity is O(n*(2**n)). The space complexity is O(2**n)
// The nil, false is returned if the length of s is equal or bigger
// then size of int of actual architecture.
func PowerSetItr(s []interface{}) (ps [][]interface{}, ok bool) {
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
