// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package bstrees

// FindFirstGreaterK returns the first node in tree greater then k
// (when encountering in-order) or nil if such a node isn't present.
// The time complexity is O(log(n)). The O(1) additional space is needed.
func FindFirstGreaterK(tree *BSTree, k int) *BSTree {
	var g *BSTree
	node := tree
	for node != nil {
		if k < node.Data.(int) {
			g, node = node, node.left
		} else {
			node = node.right
		}
	}
	return g
}
