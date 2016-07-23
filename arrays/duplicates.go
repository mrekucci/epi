// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package arrays

// DeleteDuplicates updates sorted slice xs so that all duplicates have been
// removed and the reaming elements have been shifted left to fill the emptied
// indices. The function returns the number of valid elements in xs.
// The time complexity is O(n), and O(1) additional space is needed.
func DeleteDuplicates(xs []int) (d int) {
	if len(xs) == 0 {
		return 0
	}

	j := 1
	for i := 1; i < len(xs); i++ {
		if xs[i-1] != xs[i] {
			xs[j] = xs[i]
			j++
		}
	}
	return j
}
