// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package strings

// ReverseItr returns reversed string s.
// Note it uses iterative way to reverse the string.
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

// reverseAux recursively call it self with auxiliary variables to reverse r.
func reverseAux(r []rune, i, j int) {
	if i < j {
		r[i], r[j] = r[j], r[i]
		reverseAux(r, i+1, j-1)
	}
}

// ReverseRecAux returns reversed string s.
// Note: it uses recursive way with auxiliary variables to reverse the string.
// Warning: may cause stack overflow for long strings.
func ReverseRecAux(s string) string {
	r := []rune(s)
	reverseAux(r, 0, len(r)-1)
	return string(r)
}

// reversePure recursively call it self to reverse r.
func reversePure(r []rune) []rune {
	if len(r) == 0 {
		return r
	}
	return append(reversePure(r[1:]), r[0])
}

// ReverseRecPure returns reversed string s.
// Note: it uses pure recursive way to reverse the string.
// Warning: may cause stack overflow for long strings.
func ReverseRecPure(s string) string {
	return string(reversePure([]rune(s)))
}
