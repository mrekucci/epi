// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package strings

// reverseBytes reverse elements of a[s:e] in a.
func reverseBytes(a []byte, s, e int) {
	for e > s {
		a[s], a[e] = a[e], a[s]
		s++
		e--
	}
}

// ReverseWords returns a new string containing the words from s in reverse order.
func ReverseWords(s string) string {
	r := []byte(s)

	reverseBytes(r, 0, len(s)-1) // Reverse whole sentence.
	p := 0
	for q := p; q < len(r); q++ { // Reverse each world in the reversed sentence.
		if r[q] == ' ' {
			reverseBytes(r, p, q-1) // q-1 exclude the ' ' character from reversal.
			p = q + 1
		}
	}
	reverseBytes(r, p, len(r)-1) // Reverse the last word.

	return string(r)
}
