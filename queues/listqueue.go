// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package queues

type node struct {
	data int
	next *node
}

// IntListQueue is an implementation of the Queue interface
// for integer values implemented by circular linked list.
type IntListQueue struct {
	head, tail *node
	len        int
}

// Enqueue inserts e element at the back of the queue.
// The time complexity is O(1)
func (q *IntListQueue) Enqueue(e interface{}) {
	n := &node{data: e.(int)}
	if q.head == nil {
		q.head = n
		n.next = q.head
		q.tail = q.head
	} else {
		q.tail.next = n
		n.next = q.head
		q.tail = n
	}
	q.len++
}

// Dequeue removes and returns the front integer element from this queue.
// The time complexity is O(1)
func (q *IntListQueue) Dequeue() interface{} {
	switch q.len {
	case 0:
		return nil
	case 1:
		v := q.head.data
		q.head = nil
		q.tail = nil
		q.len = 0
		return v
	}
	n := q.head
	q.head = n.next
	q.tail.next = q.head
	n.next = nil // Avoid memory leaks.
	q.len--
	return n.data
}

// Len returns the length of this queue.
// The time complexity is O(1)
func (q *IntListQueue) Len() int { return q.len }
