// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

// Package epiutil is an utility class with functions
// common to all packages in the epi project.
package epiutil

import "math/rand"

// RandStr returns a string of length n constructed
// from pseudo-randomly selected characters from t.
// The pseudo-randomness uses random values from s.
func RandStr(n int, t string, s rand.Source) string {
	l := len(t) - 1
	chars := make([]byte, n)
	if l > 0 {
		for i, p := range rand.New(s).Perm(n) {
			chars[i] = t[p%l]
		}
	}
	return string(chars)
}
