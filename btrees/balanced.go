// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package btrees

// state represents the state of the binary tree.
type state struct {
	isBalanced bool
	height     int
}

// checkState returns the status of the binary tree t.
func checkState(t *BTree) state {
	if t == nil {
		return state{true, -1} // Base case.
	}

	// Postorder walk.
	lt := checkState(t.left)
	if !lt.isBalanced {
		return lt
	}
	rt := checkState(t.right)
	if !rt.isBalanced {
		return rt
	}

	d := lt.height - rt.height
	max := lt.height
	if lt.height < rt.height {
		max = rt.height
	}
	return state{d <= 1 && d >= -1, max + 1}

}

// IsBalanced returns true if t is a balanced binary tree.
// The time complexity is O(n). The O(h) additional space
// is needed (where h is the height of the tree).
func IsBalanced(t *BTree) bool { return checkState(t).isBalanced }
