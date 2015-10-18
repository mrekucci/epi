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

type stackTest struct {
	e   interface{}
	err error
}

// checkLength returns an error if s.Len() != len.
func checkLength(s Stack, len int) error {
	if s.Len() != len {
		return fmt.Errorf("s.Len() = %d; want %d", s.Len(), len)
	}
	return nil
}

func checkPush(s Stack, test stackTest, cnt *int) error {
	if err := s.Push(test.e); err != test.err { // Test push error.
		return fmt.Errorf("s.Push(%v) = %v; want %v", test.e, err, test.err)
	}
	if test.err == nil { // Check length when test doesn't have an error.
		*cnt++
		if err := checkLength(s, *cnt); err != nil {
			return fmt.Errorf("s.Push(%v) got %v", test.e, err)
		}
	}
	return nil
}

func checkPop(s Stack, test stackTest, cnt *int) error {
	if test.err == nil {
		if got, want := s.Pop(), test.e; got != want {
			return fmt.Errorf("s.Pop() = %v; want %v", got, want)
		}
		*cnt--
		if err := checkLength(s, *cnt); err != nil {
			return fmt.Errorf("s.Pop() got %v", err)
		}
	}
	return nil
}

// testStackInterface tests Stack interface method implementation.
func testStackInterface(t *testing.T, s Stack, tests []stackTest) {
	pushed := 1
	// Test pop of empty stack.
	if err := checkPop(s, stackTest{nil, nil}, &pushed); err != nil {
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
	ifaceTests := []stackTest{
		{0, nil},
		{1, nil},
		{2, nil},
		{3, nil},
		{4, nil},
		{5, nil},
		{6, nil},
		{7, nil},
		{8, nil},
		{9, nil},
		{minInt, nil},
		{maxInt, nil},
		{"x", ErrType},
	}
	testStackInterface(t, new(IntStack), ifaceTests)
}
