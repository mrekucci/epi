// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package recursion

// genPerm generates all the permutation of an into the res.
func genPerm(i int, an []int, perm [][]int) [][]int {
	if i == len(an)-1 { // Base case. Copy actual permutation of an to the result.
		return append(perm, append([]int(nil), an...))
	}
	for j := i; j < len(an); j++ {
		an[i], an[j] = an[j], an[i]
		perm = genPerm(i+1, an, perm) // To an[i] generate all permutations of an[i+1:].
		an[i], an[j] = an[j], an[i]
	}
	return perm
}

// Permutations returns all the possible permutations of an.
// The time complexity is O(n!); The O(1) additional space is needed.
// The an slice may be modified during the execution.
func Permutations(an []int) [][]int {
	return genPerm(0, an, nil)
}
