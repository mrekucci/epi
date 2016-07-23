// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package recursion

// Permutations returns all the possible permutations of xs.
// The time and space complexity is O(n!),
// The xs slice may be modified during the execution.
func Permutations(xs []int) (p [][]int) {
	// generate generates all the permutation of xs into the p.
	var generate func(i int)
	generate = func(i int) {
		if i == len(xs)-1 { // Base case. Copy actual permutation of xs to the result.
			p = append(p, append([]int(nil), xs...))
			return
		}
		for j := i; j < len(xs); j++ {
			xs[i], xs[j] = xs[j], xs[i]
			generate(i + 1) // To xs[i] generate all permutations of xs[i+1:].
			xs[i], xs[j] = xs[j], xs[i]
		}
	}

	generate(0)
	return p
}
