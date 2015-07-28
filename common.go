// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package epi

import "math/rand"

// intSize is the size in bits of an int.
const intSize = 32 << (^uint(0) >> 63)

// Integer limit values.
const (
	maxInt = int(^uint(0) >> 1)
	minInt = -maxInt - 1
)

// randStr returns a string of length n constructed
// from pseudo-randomly selected characters from t.
func randStr(n int, t string) string {
	q := len(t) - 1
	ints := rand.New(rand.NewSource(1238)).Perm(n)
	chars := make([]byte, len(ints))
	for j, n := range ints {
		chars[j] = t[n%q]
	}
	return string(chars)
}
