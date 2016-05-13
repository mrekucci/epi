// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package strings

// RLEEncode encode successive repeated characters
// by the repetition count and the character.
// "", false will be returned if s contains digit.
// Note: this implementation works only for ASCII table characters.
// The time complexity is O(n) where n is the length of the string s.
func RLEEncode(s string) (es string, ok bool) {
	var e []byte
	cnt := 1
	for i := 0; i < len(s); i++ {
		switch {
		case s[i] > 127 || '0' <= s[i] && s[i] <= '9':
			return "", false
		case len(s) == i+1 || s[i] != s[i+1]:
			e = append(append(e, IntToString(int64(cnt))...), s[i])
			cnt = 1
		default:
			cnt++
		}
	}

	return string(e), true
}

// RLEDecode decode string s encoded as the repetition count
// of successive repeated characters and the character.
// "", false will be returned if encoding of s is broken.
// Note: this implementation works only for ASCII table characters.
// The time complexity is O(n) where n is the length of the string s.
func RLEDecode(s string) (ds string, ok bool) {
	if len(s) == 1 {
		return "", false
	}

	var d []byte
	num := true
	cnt := 0
	for _, c := range s {
		switch {
		case c > 127 || num && !('0' <= c && c <= '9'):
			return "", false
		case '0' <= c && c <= '9':
			cnt = cnt*10 + int(c-'0')
			num = false
		default:
			for cnt > 0 {
				d = append(d, byte(c))
				cnt--
			}
			num = true
		}
	}

	return string(d), true
}
