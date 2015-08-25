// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package stacks

import (
	"strings"

	"github.com/mrekucci/epi/ptypes"
)

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
	for _, t := range strings.Split(rpn, ",") {
		var o interface{}
		if eval, ok := operations[t]; ok {
			a, b := s.Pop().(int), s.Pop().(int) // Always be int.
			o = eval(a, b)
		} else {
			n, err := ptypes.StringToInt(t)
			if err != nil {
				return 0, err
			}
			o = int(n)
		}
		s.Push(o) // o is always int so push should never returned an error.
	}
	return s.Pop().(int), nil
}
