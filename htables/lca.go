// Copyright (c) 2016, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package htables

// BTreeP represents a binary tree which has also reference to its parent.
type BTreeP struct {
	Data   interface{}
	parent *BTreeP
	left   *BTreeP
	right  *BTreeP
}

// LCA returns the lowest common ancestor in the binary tree t for the nodes n0, n1.
// The time complexity is O(d0+d1), where d0 is the distance from LCA to the n0,
// and d1 is the distance from LCA to the n1, The space complexity is also O(d0+d1).
// In the worst case, when n0 and n1 are children and the LCA is root, the time
// and the space complexity is O(h), where h is the height of the tree t.
func LCA(t, n0, n1 *BTreeP) *BTreeP {
	set := make(map[*BTreeP]bool)
	for n0 != nil || n1 != nil {
		if n0 != nil {
			if set[n0] {
				return n0
			}
			set[n0] = true
			n0 = n0.parent
		}
		if n1 != nil {
			if set[n1] {
				return n1
			}
			set[n1] = true
			n1 = n1.parent
		}
	}
	return nil
}
