// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package stacks

import (
	"reflect"
	"testing"
)

func TestEvalRPN(t *testing.T) {
	for _, test := range []struct {
		rpn  string
		want int
		err  error
	}{
		// Add operation tests.
		{"2,3,+", 5, nil},
		{"-2,3,+", 1, nil},
		{"2,-3,+", -1, nil},
		{"-2,-3,+", -5, nil},

		// Sub operation tests.
		{"6,3,-", -3, nil},
		{"-6,3,-", 9, nil},
		{"6,-3,-", -9, nil},
		{"-6,-3,-", 3, nil},

		// Mul operation tests.
		{"2,3,*", 6, nil},
		{"-2,3,*", -6, nil},
		{"2,-3,*", -6, nil},
		{"-2,-3,*", 6, nil},

		// Div operation tests.
		{"3,6,/", 2, nil},
		{"-3,6,/", -2, nil},
		{"3,-6,/", -2, nil},
		{"-3,-6,/", 2, nil},
		{"2,5,/", 2, nil},

		// Mix of operations tests.
		{"1,2,+,3,4,*,-", 9, nil},
		{"1,2,3,4,5,+,-,*,/,3,4,*,+", 24, nil},

		// Error tests.
		{"-", 0, &ErrParse{ErrParseOp, 0, "-"}},
		{"x,0,-", 0, &ErrParse{ErrParseNum, 0, "x"}},
		{"0,0,&", 0, &ErrParse{ErrParseNum, 4, "&"}},
		{"0,9223372036854775808,+", 0, &ErrParse{ErrParseNum, 2, "9223372036854775808"}},
	} {
		got, err := EvalRPN(test.rpn)
		if !reflect.DeepEqual(got, test.want) || !reflect.DeepEqual(err, test.err) {
			t.Errorf("EvalRPN(%q) = %v, %v; want %v, %v", test.rpn, got, err, test.want, test.err)
		}
	}
}

func benchEvalRPN(b *testing.B, size int) {
	var rpn, o string
	for j := size / 2; j >= 0; j-- { // Construct rpn according to the size.
		rpn += "0,0,+"
		o += ",+"
	}
	rpn += o
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EvalRPN(rpn)
	}
}

func BenchmarkEvalRPN1e1(b *testing.B) { benchEvalRPN(b, 1e1) }
func BenchmarkEvalRPN1e2(b *testing.B) { benchEvalRPN(b, 1e2) }
func BenchmarkEvalRPN1e3(b *testing.B) { benchEvalRPN(b, 1e3) }
