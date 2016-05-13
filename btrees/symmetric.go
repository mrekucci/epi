// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package btrees

// IsSymmetric returns true if t is a symmetric binary tree.
func IsSymmetric(t *BTree) bool {
	if t == nil {
		return true
	}

	var checkSymmetry func(l, r *BTree) bool
	checkSymmetry = func(l, r *BTree) bool {
		switch {
		case l == nil && r == nil:
			return true
		case l != nil && r != nil:
			return l.Data == r.Data &&
				checkSymmetry(l.left, r.right) &&
				checkSymmetry(l.right, r.left)
		}
		return false
	}
	return checkSymmetry(t.left, t.right)
}
