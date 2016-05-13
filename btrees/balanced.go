// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package btrees

// IsBalanced returns true if t is a balanced binary tree.
// The time complexity is O(n). The O(h) additional space
// is needed (where h is the height of the tree).
func IsBalanced(t *BTree) bool {
	var checkBalance func(t *BTree) (height int, isBalanced bool)
	checkBalance = func(t *BTree) (height int, isBalanced bool) {
		if t == nil {
			return -1, true // Base case.
		}

		// Post-order walk.
		lh, lb := checkBalance(t.left)
		if !lb {
			return lh, lb
		}
		rh, rb := checkBalance(t.right)
		if !rb {
			return rh, rb
		}

		d := lh - rh
		max := lh
		if lh < rh {
			max = rh
		}
		return max + 1, d <= 1 && d >= -1

	}
	_, b := checkBalance(t)
	return b
}
