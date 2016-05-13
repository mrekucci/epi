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
	if t == nil {
		return -1
	}
	hl := 1 + Height(t.left)
	hr := 1 + Height(t.right)
	if hl > hr {
		return hl
	}
	return hr
}

// Walk returns elements of the binary tree t in given traversal order.
func Walk(t *BTree, order func(t *BTree, w []interface{}) []interface{}) []interface{} {
	return order(t, nil)
}

// Preorder visit the root, traverse the left
// subtree, then traverse the right subtree.
// The time complexity is O(n). The O(h) additional
// space is needed (where h is the height of the tree).
func Preorder(t *BTree, w []interface{}) []interface{} {
	if t != nil {
		w = append(w, t.Data)
		w = Preorder(t.left, w)
		w = Preorder(t.right, w)
	}
	return w
}

// Inorder traverse the left subtree, visit the
// root, then traverse the right subtree.
// The time complexity is O(n). The O(h) additional
// space is needed (where h is the height of the tree).
func Inorder(t *BTree, w []interface{}) []interface{} {
	if t != nil {
		w = Inorder(t.left, w)
		w = append(w, t.Data)
		w = Inorder(t.right, w)
	}
	return w
}

// Postorder traverse the left subtree, traverse the
// right subtree, and then visit the root.
// The time complexity is O(n). The O(h) additional
// space is needed (where h is the height of the tree).
func Postorder(t *BTree, w []interface{}) []interface{} {
	if t != nil {
		w = Postorder(t.left, w)
		w = Postorder(t.right, w)
		w = append(w, t.Data)
	}
	return w
}
