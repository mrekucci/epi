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

// testStackInterface tests Stack interface method implementation.
func testStackInterface(t *testing.T, s Stack, tests []stackTest) error {
	// Test pop of empty stack.
	if got, want := s.Pop(), interface{}(nil); got != want {
		return fmt.Errorf("s.Pop() = %v; want %v", got, want)
	}

	// Test len of empty stack.
	if err := checkLength(s, 0); err != nil {
		return err
	}

	// Test push.
	for i, test := range tests {
		if err := s.Push(test.e); err != test.err { // Test push error.
			return fmt.Errorf("s.Push(%v) = %v; want %v", test.e, err, test.err)
		}
		if test.err == nil { // Check len when test doesn't have an error.
			if err := checkLength(s, len(tests[:i+1])); err != nil {
				return fmt.Errorf("s.Push(%v) got %v", test.e, err)
			}
		}
	}

	// Test pop all pushed elms.
	for i := range tests {
		test := tests[len(tests)-i-1]
		if test.err == nil {
			if got, want := s.Pop(), test.e; got != want {
				return fmt.Errorf("s.Pop() = %v; want %v", got, want)
			}
			if err := checkLength(s, len(tests[:len(tests)-i-1])); err != nil {
				return fmt.Errorf("s.Pop() got %v", err)
			}
		}
	}

	return nil
}

type stackTest struct {
	e   interface{}
	err error
}

func TestIntStack(t *testing.T) {
	ifaceTests := []stackTest{
		{0, nil},
		{1, nil},
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
	if err := testStackInterface(t, new(IntStack), ifaceTests); err != nil {
		t.Error(err)
	}
}
