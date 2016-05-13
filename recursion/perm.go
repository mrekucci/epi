// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package recursion

// Permutations returns all the possible permutations of an.
// The time and space complexity is O(n!),
// The an slice may be modified during the execution.
func Permutations(an []int) (p [][]int) {
	// genPerm generates all the permutation of an into the p.
	var genPerm func(i int)
	genPerm = func(i int) {
		if i == len(an)-1 { // Base case. Copy actual permutation of an to the result.
			p = append(p, append([]int(nil), an...))
			return
		}
		for j := i; j < len(an); j++ {
			an[i], an[j] = an[j], an[i]
			genPerm(i+1) // To an[i] generate all permutations of an[i+1:].
			an[i], an[j] = an[j], an[i]
		}
	}

	genPerm(0)
	return p
}
