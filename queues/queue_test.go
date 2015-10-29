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

// checkEnqueue returns an error if enqueue test has failed.
func checkEnqueue(q Queue, e interface{}, cnt *int) error {
	q.Enqueue(e)
	*cnt++
	if err := checkLength(q, *cnt); err != nil {
		return fmt.Errorf("q.Enqueue(%v) got %v", e, err)
	}
	return nil
}

// checkDequeue returns an error if dequeue test has failed.
func checkDequeue(q Queue, want interface{}, cnt *int) error {
	if got := q.Dequeue(); got != want {
		return fmt.Errorf("q.Dequeue() = %v; want %v", got, want)
	}
	*cnt--
	if err := checkLength(q, *cnt); err != nil {
		return fmt.Errorf("q.Dequeue() got %v", err)
	}
	return nil
}

// testQueueInterface tests Queue interface method implementation.
func testQueueInterface(t *testing.T, q Queue, tests []interface{}) {
	enqueued := 1
	// Test dequeue of empty queue.
	if err := checkDequeue(q, nil, &enqueued); err != nil {
		t.Error(err)
	}
	// Test enqueue first 2/3 of the elements.
	for i := 0; i < 2*len(tests)/3; i++ {
		if err := checkEnqueue(q, tests[i], &enqueued); err != nil {
			t.Error(err)
		}
	}
	// Test dequeue first 1/3 of the elements from queue.
	for i := 0; i < len(tests)/3; i++ {
		if err := checkDequeue(q, tests[i], &enqueued); err != nil {
			t.Error(err)
		}
	}
	// Test enqueue last 2/3 of the elements.
	for i := len(tests) / 3; i < len(tests); i++ {
		if err := checkEnqueue(q, tests[i], &enqueued); err != nil {
			t.Error(err)
		}
	}
	// Test dequeue middle 1/3 of the element from queue.
	for i := len(tests) / 3; i < 2*len(tests)/3; i++ {
		if err := checkDequeue(q, tests[i], &enqueued); err != nil {
			t.Error(err)
		}
	}
	// Test dequeue last 2/3 of the element from queue.
	for i := len(tests) / 3; i < len(tests); i++ {
		if err := checkDequeue(q, tests[i], &enqueued); err != nil {
			t.Error(err)
		}
	}
}
