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

type queueTest struct {
	e   interface{}
	err error
}

// checkLength returns an error if q.Len() != len.
func checkLength(q Queue, len int) error {
	if q.Len() != len {
		return fmt.Errorf("q.Len() = %d; want %d", q.Len(), len)
	}
	return nil
}

// checkEnqueue returns an error if enqueue test has failed.
func checkEnqueue(q Queue, test queueTest, cnt *int) error {
	if err := q.Enqueue(test.e); err != test.err { // Test enqueue error.
		return fmt.Errorf("q.Enqueue(%v) = %v; want %v", test.e, err, test.err)
	}
	if test.err == nil { // Check len when test doesn't expect an error.
		*cnt++
		if err := checkLength(q, *cnt); err != nil {
			return fmt.Errorf("q.Enqueue(%v) got %v", test.e, err)
		}
	}
	return nil
}

// checkDequeue returns an error if dequeue test has failed.
func checkDequeue(q Queue, test queueTest, cnt *int) error {
	if test.err != nil {
		return nil
	}
	if got, want := q.Dequeue(), test.e; got != want {
		return fmt.Errorf("q.Dequeue() = %v; want %v", got, want)
	}
	*cnt--
	if err := checkLength(q, *cnt); err != nil {
		return fmt.Errorf("q.Dequeue() got %v", err)
	}
	return nil
}

// testQueueInterface tests Queue interface method implementation.
func testQueueInterface(t *testing.T, q Queue, tests []queueTest) {
	enqueued := 1
	// Test dequeue of empty queue.
	if err := checkDequeue(q, queueTest{nil, nil}, &enqueued); err != nil {
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
	for i := enqueued; i < len(tests); i++ {
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
