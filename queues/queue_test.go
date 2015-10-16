// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package queues

import (
	"fmt"
	"testing"
)

// Integer limit values.
const (
	maxInt = int(^uint(0) >> 1)
	minInt = -maxInt - 1
)

// checkLength returns an error if q.Len() != len.
func checkLength(q Queue, len int) error {
	if q.Len() != len {
		return fmt.Errorf("q.Len() = %d; want %d", q.Len(), len)
	}
	return nil
}

// testQueueInterface tests Queue interface method implementation.
func testQueueInterface(t *testing.T, q Queue, tests []queueTest) error {
	// Test dequeue of empty queue.
	if got, want := q.Dequeue(), interface{}(nil); got != want {
		return fmt.Errorf("q.Dequeue() = %v; want %v", got, want)
	}

	// Test len of empty queue.
	if err := checkLength(q, 0); err != nil {
		return err
	}

	// Test enqueue.
	enqueued := 0
	for _, test := range tests {
		if err := q.Enqueue(test.e); err != test.err { // Test enqueue error.
			return fmt.Errorf("q.Enqueue(%v) = %v; want %v", test.e, err, test.err)
		}
		if test.err == nil { // Check len when test doesn't expect an error.
			enqueued++
			if err := checkLength(q, enqueued); err != nil {
				return fmt.Errorf("q.Enqueue(%v) got %v", test.e, err)
			}
		}
	}

	// Test dequeue of all enqueued elements.
	for _, test := range tests {
		if test.err == nil {
			if got, want := q.Dequeue(), test.e; got != want {
				return fmt.Errorf("q.Dequeue() = %v; want %v", got, want)
			}
			enqueued--
			if err := checkLength(q, enqueued); err != nil {
				return fmt.Errorf("q.Dequeue() got %v", err)
			}
		}
	}

	return nil
}

type queueTest struct {
	e   interface{}
	err error
}
