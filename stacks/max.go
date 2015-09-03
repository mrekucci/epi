// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package stacks

// maxCnt holds a max value and its count as a pair.
type maxCnt struct {
	max int
	cnt int
}

// IntStackMax is an implementation of the Stack interface for integer
// values with additional Max method. It uses O(n) additional space in
// worst-case (elements are sorted ASC) to support Max O(1) time complexity.
type IntStackMax struct {
	IntStack
	m []maxCnt // Records maxim through push and delete them through pop.
}

// Push adds e on top of the stack. An error is returned if e is not of type int.
// The time complexity is O(1)
func (s *IntStackMax) Push(e interface{}) error {
	if err := s.IntStack.Push(e); err != nil {
		return err
	}

	v := e.(int)
	if len(s.m) == 0 {
		s.m = append(s.m, maxCnt{v, 1})
		return nil
	}

	switch peek := s.m[len(s.m)-1]; {
	case peek.max == v:
		peek.cnt++
		s.m[len(s.m)-1] = peek
	case peek.max < v:
		s.m = append(s.m, maxCnt{v, 1})
	}
	return nil
}

// Pop removes and returns the last added integer element from this stack.
// The time complexity is O(1)
func (s *IntStackMax) Pop() (e interface{}) {
	if e = s.IntStack.Pop(); e != nil {
		if v, peek := e.(int), s.m[len(s.m)-1]; v == peek.max {
			if peek.cnt == 1 {
				s.m = s.m[:len(s.m)-1]
			} else {
				peek.cnt--
				s.m[len(s.m)-1] = peek
			}
		}
	}
	return e
}

// Max returns a value of the biggest element in this stack.
// The time complexity is O(1)
func (s *IntStackMax) Max() interface{} {
	if len(s.m) == 0 {
		return nil
	}
	return s.m[len(s.m)-1].max
}
