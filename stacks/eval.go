// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package stacks

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// ErrParseOp indicates that a parsed token should be a valid operand.
var ErrParseOp = errors.New("expect a valid operand")

// ErrParseNum indicates that a parsed token should be a valid integer number.
var ErrParseNum = errors.New("expect a valid integer number")

// ErrParse represents a parsing error with additional records of failed parsing.
type ErrParse struct {
	error
	Index int
	Token string
}

func (e *ErrParse) Error() string {
	return fmt.Sprintf("EvalRPN: parsed token %q at index %d: %v", e.Token, e.Index, e.error)
}

// evalFn is an evaluation function template.
type evalFn func(a, b int) int

// operation defines possible operations of rpn expression.
var operations = map[string]evalFn{
	"+": func(a, b int) int { return a + b },
	"-": func(a, b int) int { return a - b },
	"*": func(a, b int) int { return a * b },
	"/": func(a, b int) int { return a / b },
}

// EvalRPN parses rpn expression and returns the number that
// the expression evaluates to. An rpn expression can be evaluated
// uniquely to an integer. The operations are not check for overflows.
// An error is returned if rpn protocol is malformed.
func EvalRPN(rpn string) (int, error) {
	s := new(IntStack)
	for i, t := range strings.Split(rpn, ",") {
		var o interface{}
		if eval, ok := operations[t]; ok {
			a, b := s.Pop(), s.Pop()
			if a == nil || b == nil {
				return 0, &ErrParse{ErrParseOp, i * 2, t} // i*2 'cause we stripped "," from rpn.
			}
			o = eval(a.(int), b.(int))
		} else {
			n, err := strconv.Atoi(t)
			if err != nil {
				return 0, &ErrParse{ErrParseNum, i * 2, t} // i*2 'cause we stripped "," from rpn.
			}
			o = int(n)
		}
		s.Push(o) // o is always int so push should never returned an error.
	}
	return s.Pop().(int), nil
}
