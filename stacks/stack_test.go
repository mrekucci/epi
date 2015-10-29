// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package stacks

import (
	"fmt"
	"testing"
)

// Integer limit values.
const (
	maxInt = int(^uint(0) >> 1)
	minInt = -maxInt - 1
)

// checkLength returns an error if s.Len() != len.
func checkLength(s Stack, len int) error {
	if s.Len() != len {
		return fmt.Errorf("s.Len() = %d; want %d", s.Len(), len)
	}
	return nil
}

func checkPush(s Stack, e interface{}, cnt *int) error {
	s.Push(e)
	*cnt++
	if err := checkLength(s, *cnt); err != nil {
		return fmt.Errorf("s.Push(%v) got %v", e, err)
	}
	return nil
}

func checkPop(s Stack, want interface{}, cnt *int) error {
	if got := s.Pop(); got != want {
		return fmt.Errorf("s.Pop() = %v; want %v", got, want)
	}
	*cnt--
	if err := checkLength(s, *cnt); err != nil {
		return fmt.Errorf("s.Pop() got %v", err)
	}
	return nil
}

// testStackInterface tests Stack interface method implementation.
func testStackInterface(t *testing.T, s Stack, tests []interface{}) {
	pushed := 1
	// Test pop of empty stack.
	if err := checkPop(s, nil, &pushed); err != nil {
		t.Error(err)
	}
	// Test push first 2/3 of the elements.
	for i := 0; i < 2*len(tests)/3; i++ {
		if err := checkPush(s, tests[i], &pushed); err != nil {
			t.Error(err)
		}
	}
	// Test pop first 1/3 of the elements from stack.
	for i := 2*len(tests)/3 - 1; i >= len(tests); i-- {
		if err := checkPop(s, tests[i], &pushed); err != nil {
			t.Error(err)
		}
	}
	// Test push last 2/3 of the elements.
	for i := 2 * len(tests) / 3; i < len(tests); i++ {
		if err := checkPush(s, tests[i], &pushed); err != nil {
			t.Error(err)
		}
	}
	// Test pop all the element from stack.
	for i := len(tests) - 1; i >= 0; i-- {
		if err := checkPop(s, tests[i], &pushed); err != nil {
			t.Error(err)
		}
	}
}

func TestIntStack(t *testing.T) {
	testStackInterface(t, new(IntStack), []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, minInt, maxInt})
}
