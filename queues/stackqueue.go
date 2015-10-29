// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package queues

import "github.com/mrekucci/epi/stacks"

// IntStackQueue is an implementation of the Queue interface
// for integer values implemented by stack.
type IntStackQueue struct {
	enq, deq stacks.IntStack
}

// Enqueue inserts e element at the back of the queue.
// The time complexity is O(1) amortized.
func (q *IntStackQueue) Enqueue(e interface{}) { q.enq.Push(e) }

// Dequeue removes and returns the front integer element from this queue.
// The time complexity is O(1) amortized.
func (q *IntStackQueue) Dequeue() interface{} {
	switch {
	case q.Len() == 0:
		return nil
	case q.deq.Len() == 0:
		for q.enq.Len() != 1 {
			q.deq.Push(q.enq.Pop())
		}
		return q.enq.Pop()
	}
	return q.deq.Pop()
}

// Len returns the length of this queue.
// The time complexity is O(1).
func (q *IntStackQueue) Len() int { return q.enq.Len() + q.deq.Len() }

// NewIntStackQueue returns a new initialized *IntStackQueue.
func NewIntStackQueue() Queue {
	return &IntStackQueue{stacks.IntStack{}, stacks.IntStack{}}
}
