// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package recursion

// genPerm generates all the permutation of an into the res.
func genPerm(i int, an []int, res *[][]int) {
	if i == len(an)-1 { // Base case. Copy actual permutation of an to the result.
		*res = append(*res, append([]int(nil), an...))
	} else {
		for j := i; j < len(an); j++ {
			an[i], an[j] = an[j], an[i]
			genPerm(i+1, an, res) // To an[i] generate all permutations of an[i+1:].
			an[i], an[j] = an[j], an[i]
		}
	}
}

// Permutations returns all the possible permutations of an.
// The time complexity is O(n!); The O(1) additional space is needed.
// The an slice may be modified during the execution.
func Permutations(an []int) (p [][]int) {
	if len(an) == 0 {
		return nil
	}
	genPerm(0, an, &p)
	return p
}
