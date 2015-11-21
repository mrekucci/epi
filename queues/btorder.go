// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package queues

// IntBTree represents a binary tree with int key.
type IntBTree struct {
	Data  int
	left  *IntBTree
	right *IntBTree
}

// DepthOrder returns a slice consisting of keys belonging to the
// same level (order is from left to right) of the binary tree t.
// The time complexity is O(n). The O(m) additional space is
// needed, where m is the maximum number of nodes at any level.
func DepthOrder(t *IntBTree) [][]int {
	var r [][]int
	var l []int
	pq := new(arrayQueue)
	pq.Enqueue(t) // Add root.
	c := pq.Len() // Number of elements on the same level.
	for pq.Len() != 0 {
		n := pq.Dequeue()
		c--
		if n != (*IntBTree)(nil) {
			st := n.(*IntBTree)
			pq.Enqueue(st.left)
			pq.Enqueue(st.right)
			l = append(l, st.Data)
		}
		if c == 0 {
			c = pq.Len() // Set count to the number of elements that should be processed on next level.
			if len(l) != 0 {
				r = append(r, l)
				l = []int(nil)
			}
		}
	}
	return r
}
