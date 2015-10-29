// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package queues

// Queue interface define a basic queue operations.
type Queue interface {
	// Insert element e at the back of the queue.
	Enqueue(e interface{})

	// Remove and return the front element from the queue.
	Dequeue() interface{}

	// Length of the queue.
	Len() int
}
