// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package lists

// CreateCycle returns cycled linked list created from data.
// The cycle in the list is created to node on index ci of data.
// The returned node represents the reference to the start of the
// cycle or nil is returned when there is no cycle in returned list.
// If ci < 0 || ci >= len(data), then a list without cycle is returned.
func CreateCycle(data []interface{}, ci int) (*List, *Node) {
	if ci < 0 || ci >= len(data) {
		return NewFromSlice(data), nil
	}

	l := new(List)
	var csn *Node
	for i := 0; i <= ci; i++ {
		csn = l.Insert(&Node{Data: data[i]})
	}
	ln := csn
	for i := ci + 1; i < len(data); i++ {
		ln = l.Insert(&Node{Data: data[i]})
	}
	ln.next = csn // Create a cycle from the ln to the csn.

	return l, csn
}

// HasCycle returns *Node that represents the start of
// the cycle in l or nil when there is no cycle in l.
// The time complexity is O(n); the space efficiency is O(1).
func HasCycle(l *List) *Node {
	s, f := l.head, l.head
	for f != nil && f.next != nil && f.next.next != nil {
		s, f = s.next, f.next.next // f diverge 2x faster then s: 2s == f.
		if f == s {                // Cycle detected (we don't know whether the s is at the start or inside the cycle).
			s = l.head
			for f != s { // It'll take i iteration to find the starting node (where i is the index of node in the list where the cycle starts).
				s, f = s.next, f.next
			}
			return s
		}
	}
	return nil
}

// HasCycleAlt returns *Node that represents the start of
// the cycle in l or nil when there is no cycle in l.
// Slower alternative to HasCycle. The time complexity is O(n)
// with big constant value; the space efficiency is O(n).
func HasCycleAlt(l *List) *Node {
	ns := make(map[*Node]bool)
	for n := l.head; n != nil; n = n.next {
		if ns[n] {
			return n
		}
		ns[n] = true
	}
	return nil
}
