// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package bstrees

// FindFirstK returns the first occurrence of k in tree
// (when encountering in-order) or nil if k isn't present.
// The time complexity is O(h) where h is the tree height.
// The O(1) additional space is needed.
func FindFirstK(tree *BSTree, k int) *BSTree {
	var c *BSTree
	node := tree
	for node != nil {
		switch v := node.Data.(int); {
		case v > k:
			node = node.left
		case v == k:
			c, node = node, node.left
		case v < k:
			node = node.right
		}
	}
	return c
}
