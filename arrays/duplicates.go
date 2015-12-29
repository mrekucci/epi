// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package arrays

// DeleteDuplicates updates sorted slice an so that all duplicates have been
// removed and the reaming elements have been shifted left to fill the emptied
// indices. The function returns the number of valid elements in an.
// The time complexity is O(n), and O(1) additional space is needed.
func DeleteDuplicates(an []int) (d int) {
	if len(an) == 0 {
		return 0
	}

	j := 1
	for i := 1; i < len(an); i++ {
		if an[i-1] != an[i] {
			an[j] = an[i]
			j++
		}
	}
	return j
}
