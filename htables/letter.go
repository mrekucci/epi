// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package htables

// IsLetterSubset returns true if the letter can be constructed from characters
// appearing in the magazine. The time complexity is O(m+n) where m=len(letter)
// and n=len(magazine). The space complexity is O(l), where l is the number of
// distinct characters in the letter.
func IsLetterSubset(letter, magazine string) bool {
	if len(letter) > len(magazine) {
		return false
	}

	f := make(map[rune]int)
	for _, w := range letter {
		f[w]++
	}

	for _, w := range magazine {
		switch cnt, ok := f[w]; {
		case cnt == 1:
			delete(f, w)
		case ok:
			f[w]--
		}
		if len(f) == 0 {
			return true
		}
	}
	return false
}
