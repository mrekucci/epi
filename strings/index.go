// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package strings

// IndexNaive returns an start index of pattern p in the string s.
// It uses sequential search method for matching a pattern p in the string s.
// Let n be len(s) and m be len(p), then the matching time is O((n−m+1)*m)
// without any pre-processing time.
func IndexNaive(s, p string) int {
	m := len(p)
	if m == 0 {
		return 0
	}
	for i := m; i <= len(s); i++ {
		if s[i-m:i] == p {
			return i - m
		}
	}
	return -1
}

// IndexRK returns an start index of pattern p in the string s.
// It uses Rabin-Karp method for matching a pattern p in the string s.
// Let n be len(s) and m be len(p), then the time for pre-processing is Θ(m).
// And the time for matching:
// 	- average: Θ(n+m+1)
// 	- worst: Θ((n−m+1)*m)
//
func IndexRK(s, p string) int {
	const b = 977 // We use prime number as a base to reduce number of collisions in hash.

	n := len(s)
	m := len(p)

	switch {
	case m == 0:
		return 0
	case m == n:
		if p == s {
			return 0
		}
		return -1
	case m > n:
		return -1
	}

	// We don't manually reduce the computed hash values (and base power) by modulus.
	// Instead we use the fact that overflow of integer multiplication and addition acts as a modulus.
	// For example, overflow of the two unit32 values: a*b == (a*b)%(2**32-1).
	// The increased number of potential collisions, when modulus isn't a large
	// prime, is outweighed by the fact that we wouldn't need to explicitly use
	// modular reduction for every addition and multiplication. This results in
	// a better performance characteristics.

	// Pre-processing.
	var bp, ph, sh uint32 = 1, 0, 0
	for i := 0; i < m; i++ {
		if i > 0 { // Base power: bp = b**(m-1)%(2**32-1).
			bp *= b // Implicit modulus 2**32-1.
		}
		ph = b*ph + uint32(p[i]) // Implicit modulus 2**32-1.
		sh = b*sh + uint32(s[i]) // Implicit modulus 2**32-1.
	}

	// Matching.
	for i := m; i <= n; i++ {
		if ph == sh {
			if s[i-m:i] == p {
				return i - m
			}
		}
		if i < n {
			// Slide right the rolling hash window (include implicit modulus 2**32-1).
			sh -= bp * uint32(s[i-m]) // Remove previous value.
			sh = b*sh + uint32(s[i])  // Add next value.
		}
	}

	return -1
}
