// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package strings

// ReverseItr returns reversed string s.
// Note it uses iterative way to reverse the string.
// The time complexity is O(n), and O(1) additional space is needed.
func ReverseItr(s string) string {
	r := []rune(s)
	i, j := 0, len(r)-1
	for i < j {
		r[i], r[j] = r[j], r[i]
		i++
		j--
	}
	return string(r)
}

// ReverseRecAux returns reversed string s.
// The time and the space complexity is O(n).
// Note: it uses recursive way with auxiliary variables to reverse the string.
// Warning: may cause stack overflow for long strings.
func ReverseRecAux(s string) string {
	r := []rune(s)
	var reverse func(i, j int)
	reverse = func(i, j int) {
		if i < j {
			r[i], r[j] = r[j], r[i]
			reverse(i+1, j-1)
		}
	}
	reverse(0, len(r)-1)
	return string(r)
}

// ReverseRecPure returns reversed string s.
// The time and the space complexity is O(n).
// Note: it uses pure recursive way to reverse the string.
// Warning: may cause stack overflow for long strings.
func ReverseRecPure(s string) string {
	var reverse func(r []rune) []rune
	reverse = func(r []rune) []rune {
		if len(r) == 0 {
			return r
		}
		return append(reverse(r[1:]), r[0])
	}
	return string(reverse([]rune(s)))
}
