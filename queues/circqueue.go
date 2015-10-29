// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package queues

// arrayQueue is an implementation of the Queue interface implemented by array.
type arrayQueue struct {
	head, tail int
	len        int
	an         []interface{}
}

// Enqueue inserts e element at the back of the queue.
// The time complexity is O(1) amortized.
func (q *arrayQueue) Enqueue(e interface{}) {
	if q.Len() == cap(q.an) { // Resize.
		an := make([]interface{}, 2*cap(q.an)+1)
		n := copy(an, q.an[q.head:])
		copy(an[n:], q.an[:q.tail])
		q.head = 0
		q.tail = q.Len()
		q.an = an
	}
	q.an[q.tail] = e
	q.tail = (q.tail + 1) % cap(q.an)
	q.len++
}

// Dequeue removes and returns the front element from this queue.
// The time complexity is O(1)
func (q *arrayQueue) Dequeue() interface{} {
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
func (q *arrayQueue) Len() int { return q.len }

// NewArrayQueue returns a new Queue with given size.
func NewArrayQueue(size int) Queue { return &arrayQueue{an: []interface{}{}} }
