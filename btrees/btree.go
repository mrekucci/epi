// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package btrees

// BTree represents a binary tree.
type BTree struct {
	Data  interface{}
	left  *BTree
	right *BTree
}

// BTreeP represents a binary tree which has also reference to its parent.
type BTreeP struct {
	Data   interface{}
	parent *BTreeP
	left   *BTreeP
	right  *BTreeP
}

// Height returns the height of the binary tree t.
// The time complexity is O(n). The O(h) additional
// space is needed (where h is the height of the tree).
func Height(t *BTree) int {
	hl, hr := -1, -1
	if t != nil {
		hl = 1 + Height(t.right)
		hr = 1 + Height(t.left)
	}
	if hl > hr {
		return hl
	}
	return hr
}

// Walk returns elements of the binary tree t in given traversal order.
func Walk(t *BTree, order func(t *BTree, w *[]interface{})) []interface{} {
	var w []interface{}
	order(t, &w)
	return w
}

// Preorder visit the root, traverse the left
// subtree, then traverse the right subtree.
// The time complexity is O(n). The O(h) additional
// space is needed (where h is the height of the tree).
func Preorder(t *BTree, w *[]interface{}) {
	if t != nil {
		*w = append(*w, t.Data)
		Preorder(t.left, w)
		Preorder(t.right, w)
	}
}

// Inorder traverse the left subtree, visit the
// root, then traverse the right subtree.
// The time complexity is O(n). The O(h) additional
// space is needed (where h is the height of the tree).
func Inorder(t *BTree, w *[]interface{}) {
	if t != nil {
		Inorder(t.left, w)
		*w = append(*w, t.Data)
		Inorder(t.right, w)
	}
}

// Postorder traverse the left subtree, traverse the
// right subtree, and then visit the root.
// The time complexity is O(n). The O(h) additional
// space is needed (where h is the height of the tree).
func Postorder(t *BTree, w *[]interface{}) {
	if t != nil {
		Postorder(t.left, w)
		Postorder(t.right, w)
		*w = append(*w, t.Data)
	}
}
