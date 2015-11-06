// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package heaps

// reverseInts reverse elements of an[i:j] in an.
func reverseInts(an []int, i, j int) {
	for i < j {
		an[i], an[j] = an[j], an[i]
		i++
		j--
	}
}

// SortK sorts k-increasing-decreasing slice an and returns the result.
// The time complexity is O(n*log(k)). Beyond the space needed to write
// the final result, the O(k) additional space is needed.
// The an can be modified during the function execution.
func SortK(an []int) []int {
	i := 0
	o := 1 // Order: 1 - increasing, -1 decreasing.
	var ss [][]int
	for j := 1; j <= len(an); j++ {
		if j == len(an) || o > 0 && an[j-1] > an[j] || o < 0 && an[j-1] < an[j] {
			if o < 0 {
				reverseInts(an, i, j-1)
			}
			ss = append(ss, an[i:j])
			i, o = j, -o
		}
	}
	return MergeSorted(ss)
}
