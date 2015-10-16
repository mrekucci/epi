// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package queues

import "errors"

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
