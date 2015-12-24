// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package greedy

import "sort"

// PairTasks compute an optimum set of pairs from given tasks
// duration such that the maximum pair sum is minimum.
// The time complexity is O(n*log(n)). Beyond the space needed
// to write the final result, O(1) additional space is needed.
func PairTasks(tasks []int) (pairs [][]int) {
	if len(tasks) == 0 {
		return pairs
	}

	sort.Ints(tasks)
	pairs = make([][]int, len(tasks)/2)
	// If len(tasks) is odd we will ignore the last element and add it after the iteration.
	for i, j := 0, 2*len(pairs)-1; i < j; i, j = i+1, j-1 {
		pairs[i] = []int{tasks[i], tasks[j]}
	}

	if len(tasks)%2 != 0 {
		pairs = append(pairs, []int{tasks[len(tasks)-1]})
	}
	return pairs
}
