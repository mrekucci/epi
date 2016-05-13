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
func DepthOrder(t *IntBTree) (order [][]int) {
	var level []int
	pq := new(arrayQueue)
	pq.Enqueue(t) // Add root.
	cnt := pq.Len() // Number of elements on the same level.
	for pq.Len() != 0 {
		n := pq.Dequeue().(*IntBTree)
		cnt--
		if n != nil {
			pq.Enqueue(n.left)
			pq.Enqueue(n.right)
			level = append(level, n.Data)
		}
		if cnt == 0 {
			cnt = pq.Len() // Set count to the number of elements that should be processed on next level.
			if len(level) != 0 {
				order = append(order, level)
				level = []int(nil)
			}
		}
	}
	return order
}
