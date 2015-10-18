// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package stacks

import (
	"errors"
	"fmt"
)

// ErrType indicates that a value is no of the expected type.
var ErrType = errors.New("stacks: unexpected type")

// Stack interface define a basic stack operations.
type Stack interface {
	// Push adds e on top of the stack. An error is returned if e was not added.
	Push(e interface{}) error

	// Remove and return last added element from this stack.
	Pop() interface{}

	// Length of the stack.
	Len() int
}

// IntStack is an implementation of the Stack interface for integer values.
type IntStack []int

// Push adds e on top of the stack. An error is returned if e is not of type int.
// The time complexity is O(1) amortized.
func (s *IntStack) Push(e interface{}) error {
	v, ok := e.(int)
	if !ok {
		return ErrType
	}
	*s = append(*s, v)
	fmt.Println(*s)
	return nil
}

// Pop removes and returns the last added integer element from this stack.
// The time complexity is O(1) amortized.
func (s *IntStack) Pop() (e interface{}) {
	if s.Len() == 0 {
		return nil
	}
	e = (*s)[s.Len()-1]
	if cap(*s) > 64 && float64(s.Len()/cap(*s)) < 0.75 { // Free memory when the length of a slice shrunk enough.
		*s = append([]int(nil), (*s)[:s.Len()-1]...)
	} else {
		*s = (*s)[:s.Len()-1]
	}
	fmt.Println(*s)
	return e
}

// Len returns the length of this stack.
func (s *IntStack) Len() int {
	return len(*s)
}
