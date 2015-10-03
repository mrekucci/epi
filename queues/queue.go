// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package queues

import (
	"container/list"
	"errors"
)

// ErrType indicates that a value is no of the expected type.
var ErrType = errors.New("queues: unexpected type")

// Queue interface define a basic queue operations.
type Queue interface {
	// Insert element e at the back of the queue.
	// An error is returned if e was not added.
	Enqueue(e interface{}) error

	// Remove and return the front element from the queue.
	Dequeue() interface{}

	// Length of the queue.
	Len() int
}

// IntListQueue is an implementation of the Queue interface
// for integer values backed by doubly linked list.
type IntListQueue struct {
	l *list.List
}

// Enqueue inserts e e at the back of the queue.
// An error is returned if e is not of type int.
// The time complexity is O(1)
func (q *IntListQueue) Enqueue(e interface{}) error {
	if _, ok := e.(int); !ok {
		return ErrType
	}
	q.l.PushBack(e)
	return nil
}

// Dequeue removes and returns the front integer element from this queue.
// The time complexity is O(1)
func (q *IntListQueue) Dequeue() interface{} {
	if q.Len() == 0 {
		return nil
	}
	return q.l.Remove(q.l.Front())
}

// Len returns the length of this queue.
// The time complexity is O(1)
func (q *IntListQueue) Len() int {
	return q.l.Len()
}

// NewIntListQueue returns an initialized IntListQueue.
func NewIntListQueue() *IntListQueue {
	q := new(IntListQueue)
	q.l = list.New()
	return q
}
