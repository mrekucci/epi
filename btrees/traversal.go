// Copyright (c) 2016, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package btrees

// PBTree represents a binary tree which has also reference to its parent.
type PBTree struct {
	Data   interface{}
	parent *PBTree
	left   *PBTree
	right  *PBTree
}

// InorderTraversal traverse the left subtree, visit the root, then traverse the right subtree.
// The time complexity is O(n), and O(1) additional space is needed.
func InorderTraversal(tree *PBTree) (w []interface{}) {
	var prev, next *PBTree
	for curr := tree; curr != nil; prev, curr = curr, next {
		switch prev {
		case curr.parent: // Going down: we came down to curr from prev.
			if curr.left != nil { // Go left 'til it is possible.
				next = curr.left
			} else {
				w = append(w, curr.Data)
				// check and go to the right if the curr.right subtree is not empty. Otherwise go to the curr.parent.
				if curr.right != nil {
					next = curr.right
				} else {
					next = curr.parent
				}
			}
		case curr.left: // Going up: previously processed is the left child of current.
			w = append(w, curr.Data)
			// check and go to the right if the curr.right subtree is not empty. Otherwise go to the curr.parent.
			if curr.right != nil {
				next = curr.right
			} else {
				next = curr.parent
			}
		default: // Done with both children, now go up.
			next = curr.parent
		}
	}

	return w
}
