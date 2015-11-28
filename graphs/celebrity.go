// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package graphs

// FindCelebrity returns an index into 2D slice f that represents a celebrity
// on the party who doesn't know no one. -1 is returned if such a celebrity
// doesn't exists in f or if a person exists who doesn't know the celebrity.
// The time complexity is O(n), and O(1) additional space is needed.
func FindCelebrity(f [][]bool) int {
	r, c := 0, 1 // c starts at 1 'cause that's start of diagonal where A<->A, B<->B, C<->C, ...
	for c < len(f) {
		if f[r][c] {
			r, c = c, c+1 // All candidates less then c are not celebrity candidates.
		} else {
			c++ // r is still a celebrity candidate but c is not.
		}
	}
	for _, status := range f[r] { // Check if selected candidate is really a celebrity.
		if status {
			return -1
		}
	}
	return r
}
