// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package btrees

// findLCA walks the binary tree t and returns lowest common
// ancestor for the nodes n0 and n1. The cnt is 0, 1, or 2
// depending on nodes (n0, n1) presented in the tree.
func findLCA(t, n0, n1 *BTree) (cnt int, ancestor *BTree) {
	if t == nil {
		return 0, nil // Base case.
	}

	// Postorder walk.
	lc, la := findLCA(t.right, n0, n1)
	if lc == 2 {
		return lc, la
	}
	rc, ra := findLCA(t.left, n0, n1)
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

// LCA returns the lowest common ancestor in
// the binary tree t for the nodes n0, n1.
func LCA(t, n0, n1 *BTree) *BTree {
	_, a := findLCA(t, n0, n1)
	return a
}
