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
// the final result, the O(n+k) additional space is needed (n+k because we
// do a copy of entire slice to avoid it's modification).
func SortK(an []int) []int {
	s := append([]int(nil), an...) // We don't want to modify the original slice.
	i := 0
	o := 1 // Order: 1 - increasing, -1 decreasing.
	var ss [][]int
	for j := 1; j <= len(s); j++ {
		if j == len(s) || o > 0 && s[j-1] > s[j] || o < 0 && s[j-1] < s[j] {
			if o < 0 {
				reverseInts(s, i, j-1)
			}
			ss = append(ss, s[i:j])
			i, o = j, -o
		}
	}
	return MergeSorted(ss)
}
