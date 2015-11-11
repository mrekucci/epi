// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package htables

// CanFormPalindrome returns true if s can represent a palindrome.
// The time complexity is O(n) and O(c) additional space is needed, where
// c is the number of distinct characters in the string s.
func CanFormPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	f := make(map[rune]int)
	for _, c := range s {
		f[c]++
	}

	var odd int
	for _, cnt := range f {
		if cnt%2 != 0 {
			if odd++; odd != 1 {
				return false
			}
		}
	}
	return true
}
