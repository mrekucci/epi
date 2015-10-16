// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package queues

// IntArrayQueue is an implementation of the Queue interface
// for integer values implemented by array.
type IntArrayQueue struct {
	head, tail int
	len        int
	an         []int
}

// Enqueue inserts e element at the back of the queue.
// An error is returned if e is not of type int.
// The time complexity is O(1) amortized.
func (q *IntArrayQueue) Enqueue(e interface{}) error {
	v, ok := e.(int)
	if !ok {
		return ErrType
	}
	if q.Len() == cap(q.an) { // Resize.
		an := make([]int, 2*cap(q.an))
		n := copy(an, q.an[q.head:])
		copy(an[n:], q.an[:q.tail])
		q.head = 0
		q.tail = q.Len()
		q.an = an
	}
	q.an[q.tail] = v
	q.tail = (q.tail + 1) % cap(q.an)
	q.len++
	return nil
}

// Dequeue removes and returns the front integer element from this queue.
// The time complexity is O(1)
func (q *IntArrayQueue) Dequeue() interface{} {
	if q.Len() == 0 {
		return nil
	}
	e := q.an[q.head]
	q.head = (q.head + 1) % cap(q.an)
	q.len--
	return e
}

// Len returns the length of this queue.
// The time complexity is O(1)
func (q *IntArrayQueue) Len() int {
	return q.len
}

// NewIntArrayQueue returns a new *IntArrayQueue with given size.
// The size is set automatically to 2 if given size is less then 2.
func NewIntArrayQueue(size int) *IntArrayQueue {
	if size < 2 {
		size = 2
	}
	return &IntArrayQueue{an: make([]int, size)}
}
