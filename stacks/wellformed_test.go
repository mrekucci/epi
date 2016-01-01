// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package stacks

import "testing"

func TestIsWellFormed(t *testing.T) {
	for _, test := range []struct {
		in   string
		want bool
	}{
		{"()", true},
		{"[]", true},
		{"{}", true},
		{"[{}([])]", true},
		{"([]){()}", true},
		{"[()[]{()()}]", true},
		{"(", false},
		{"[", false},
		{"{", false},
		{"{)", false},
		{"[()[]{()()", false},
		{"0}", false},
		{"[a", false},
		{"(/)", false},
	} {
		if got := IsWellFormed(test.in); got != test.want {
			t.Errorf("IsWellFormed(%q) = %t; want %t", test.in, got, test.want)
		}
	}
}

func benchIsWellFormed(b *testing.B, size int) {
	const brackets = "()[]{}([])" // 10 brackets.
	buf := make([]byte, size)
	for i := 0; i < len(buf); i += copy(buf[i:], brackets) {
		// Copy the brackets size/len(brackets) times into the buf.
	}
	s := string(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsWellFormed(s)
	}
}

func BenchmarkIsWellFormed1e1(b *testing.B) { benchIsWellFormed(b, 1e1) }
func BenchmarkIsWellFormed1e3(b *testing.B) { benchIsWellFormed(b, 1e3) }
func BenchmarkIsWellFormed1e5(b *testing.B) { benchIsWellFormed(b, 1e5) }
