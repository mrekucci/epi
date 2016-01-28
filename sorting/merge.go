// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package sorting

// MergeInPlace extends x by length of y and merges sorted slices x and y
// into the extended x slice. The modified header of x slice is returned.
// The time complexity is O(n+m), where n=len(x), m=len(y).
// The O(1) additional space is needed beyond the expansion of x by m.
func MergeInPlace(x, y []int) []int {
	if len(y) == 0 {
		return x
	}

	i, j := len(x)-1, len(y)-1
	x = append(x, y...)
	k := i + j + 1
	for i >= 0 && j >= 0 {
		ex, ey := x[i], y[j]
		if ex > ey {
			x[k] = ex
			i--
		} else {
			x[k] = ey
			j--
		}
		k--
	}

	// Append what's left.
	for j >= 0 {
		x[k] = y[j]
		k--
		j--
	}
	return x
}
