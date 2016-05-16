// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package recursion

// AllSubsets generates all size k subsets of {1,2,...,n}.
// The number of all generated subsets is n!/(k!(n-k)!), which
// is combination without repetition.
// The time and the space complexity is O(c*k), where c=n!/(k!(n-k)!).
func AllSubsets(n, k int) (subset [][]int) {
	if k == 0 {
		return nil
	}
	set := make([]int, k)
	var combinations func(o, p int)
	combinations = func(o, p int) {
		if p == k {
			subset = append(subset, append([]int(nil), set...))
			return
		}
		for i := o; i <= n; i++ { // First execution of this cycle generates the first values in every subset.
			set[p] = i
			combinations(i+1, p+1)
		}
	}
	combinations(1, 0)
	return subset
}
