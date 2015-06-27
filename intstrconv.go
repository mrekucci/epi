// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package epi

import (
	"errors"
	"math"
)

var errSyntax = errors.New("ch5.stringToInt: invalid syntax")
var errRange = errors.New("ch5.stringToInt: value out of range")

// StringToInt converts number represented by string with base 10 to integer.
func StringToInt(s string) (int64, error) {
	const cutoff = math.MaxInt64/10 + 1 // The first smallest number such that cutoff*10 > MaxInt64.

	if len(s) == 0 {
		return 0, errSyntax
	}

	neg := false
	if s[0] == '+' {
		s = s[1:]
	} else if s[0] == '-' {
		neg = true
		s = s[1:]
	}

	var u uint64
	for i := range s {
		if s[i] < '0' || s[i] > '9' {
			return 0, errSyntax
		}

		if u >= cutoff {
			// u*10 overflows.
			return 0, errRange
		}
		u *= 10

		nu := u + uint64(s[i]-'0')
		if neg && nu > -math.MinInt64 || !neg && nu > math.MaxInt64 {
			// u+v overflows.
			return 0, errRange
		}
		u = nu
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

	var s [19 + 1]byte // 10 is max digits of int64; +1 for sign.
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
