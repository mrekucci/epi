// Copyright (c) 2016, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package sorting

// Node is a node of a linked list.
type Node struct {
	Data int
	next *Node
}

// SortList sorts linked list using insertion sort algorithm.
// The time complexity is O(n*n); The space complexity is O(1).
func SortList(list *Node) *Node {
	head, suc := &Node{next: list}, list
	// The sublist from start up to and including suc is sorted in asc order.
	for suc != nil && suc.next != nil {
		if suc.Data > suc.next.Data {
			dst, pre := suc.next, head
			for pre.next.Data < dst.Data {
				pre = pre.next
			}
			pre.next, suc.next, dst.next = dst, dst.next, pre.next
		} else {
			suc = suc.next
		}
	}
	return head.next
}

// TODO: implement the merge sort solution for linked list.
