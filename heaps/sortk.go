// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package heaps

// reverseInts reverse elements of xs[i:j] in xs.
func reverseInts(xs []int, i, j int) {
	for i < j {
		xs[i], xs[j] = xs[j], xs[i]
		i++
		j--
	}
}

// SortK sorts k-increasing-decreasing slice xs and returns the result.
// The time complexity is O(n*log(k)). Beyond the space needed to write
// the final result, the O(k) additional space is needed.
// The xs can be modified during the function execution.
func SortK(xs []int) []int {
	i, o := 0, 1 // o - Order: 1 - increasing, -1 - decreasing.
	var ss [][]int
	for j := 1; j <= len(xs); j++ {
		if j == len(xs) || o > 0 && xs[j-1] > xs[j] || o < 0 && xs[j-1] < xs[j] {
			if o < 0 {
				reverseInts(xs, i, j-1)
			}
			ss = append(ss, xs[i:j])
			i, o = j, -o
		}
	}
	return MergeSorted(ss)
}
