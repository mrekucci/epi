// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package btrees

// LCA returns the lowest common ancestor in
// the binary tree t for the nodes n0, n1.
func LCA(t, n0, n1 *BTree) *BTree {
	var findLCA func(t *BTree) (cnt int, ancestor *BTree)
	findLCA = func(t *BTree) (cnt int, ancestor *BTree) {
		if t == nil {
			return 0, nil // Base case.
		}

		// Post-order walk.
		lc, la := findLCA(t.left)
		if lc == 2 {
			return lc, la
		}
		rc, ra := findLCA(t.right)
		if rc == 2 {
			return rc, ra
		}

		cnt = lc + rc
		if t == n0 {
			cnt++
		}
		if t == n1 {
			cnt++
		}
		if cnt == 2 {
			ancestor = t
		}
		return cnt, ancestor
	}

	_, a := findLCA(t)
	return a
}
