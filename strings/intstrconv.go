// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package strings

import (
	"errors"
	"math"
)

// ErrSyntax indicates that a value does not have the right syntax.
var ErrSyntax = errors.New("StringToInt: invalid syntax")

// ErrRange indicates that a value is out of range.
var ErrRange = errors.New("StringToInt: value out of range")

// StringToInt converts number represented by string with base 10 to integer.
func StringToInt(s string) (int64, error) {
	const cutoff = math.MaxInt64/10 + 1 // The first smallest number such that cutoff*10 > MaxInt64.

	if len(s) == 0 {
		return 0, ErrSyntax
	}

	neg := false
	if s[0] == '+' {
		s = s[1:]
	} else if s[0] == '-' {
		neg = true
		s = s[1:]
	}

	var u uint64
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0, ErrSyntax
		}

		if u >= cutoff { // Check if u*10 overflows.
			return 0, ErrRange
		}
		u *= 10

		u += uint64(c-'0')
		if neg && u > -math.MinInt64 || !neg && u > math.MaxInt64 { // Check for overflows: -n < math.MinInt64 || n > math.MaxInt64
			return 0, ErrRange
		}
	}

	n := int64(u)
	if neg {
		n = -n
	}
	return n, nil
}

// IntToString converts integer to string.
func IntToString(n int64) string {
	if n == 0 {
		return "0"
	}

	var s [19 + 1]byte // 19 is max digits of int64; +1 for sign.
	i := len(s)

	neg := n < 0
	u := uint64(n)
	if neg {
		u = -u // uint64(^n + 1)
	}

	for u > 0 {
		i--
		s[i] = byte(u%10 + '0')
		u /= 10
	}

	if neg {
		i--
		s[i] = '-'
	}

	return string(s[i:])
}
