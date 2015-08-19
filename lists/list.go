// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package lists

// Node is a node of a linked list.
type Node struct {
	Data interface{}
	next *Node
}

// List represents a linked list.
type List struct {
	head *Node
	len  int
}

// Insert inserts node n at the end of this list.
// The complexity is Θ(n).
func (l *List) Insert(n *Node) *Node {
	l.len++
	if l.head == nil {
		l.head = n
		return n
	}
	p := l.head
	for p.next != nil {
		p = p.next
	}
	p.next = n
	return n
}

// Remove removes node n from this list.
// The complexity is O(n)
func (l *List) Remove(n *Node) *Node {
	switch {
	case n == nil || l.head == nil: // Nothing to remove, or list is empty.
		return nil
	case l.head == n: // If n is the first node remove it.
		l.head = l.head.next
		n.next = nil
		l.len--
		return n
	}
	for p := l.head; p.next != nil; p = p.next {
		if p.next == n { // If n is the next node remove it.
			p.next = p.next.next
			n.next = nil
			l.len--
			return n
		}
	}
	return nil
}

// First returns the first Node.Data of this list or nil.
// The complexity is Θ(1).
func (l *List) First() interface{} { return l.head.Data }

// Last returns last Node.Data of this list.
// The complexity is Θ(n).
func (l *List) Last() interface{} {
	p := l.head
	for p.next != nil {
		p = p.next
	}
	return p.Data
}

// Len returns the number of elements of list l.
// The complexity is Θ(1).
func (l *List) Len() int { return l.len }

// ToSlice returns data of this list as a slice.
// The complexity is Θ(n)
func (l *List) ToSlice() (s []interface{}) {
	for p := l.head; p != nil; p = p.next {
		s = append(s, p.Data)
	}
	return s
}

// NewFromSlice returns an initialized list filled with given data.
func NewFromSlice(data []interface{}) *List {
	l := new(List)
	for _, v := range data {
		l.Insert(&Node{Data: v})
	}
	return l
}

// PopInt takes the first value from l and returns it with its node.
// A zero value for node is returned when l is empty or the first
// value doesn't contain any data. The ok is set to false when the
// first value isn't of type int.
func PopInt(l *List) (int, *Node, bool) {
	n := l.head
	if n == nil || n.Data == nil {
		return 0, n, true
	}
	i, ok := n.Data.(int)
	if !ok {
		return i, nil, ok
	}
	return i, n, ok
}
