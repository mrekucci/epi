// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package stacks

var parenthesis = map[rune]rune{'(': ')', '[': ']', '{': '}'}

// IsWellFormed returns true is given string of brackets is well formed.
// The time complexity is O(n), where n is the number of characters in s.
// The space complexity is O(n).
func IsWellFormed(s string) bool {
	stack := new(RuneStack)
	for _, c := range s {
		switch p, ok := parenthesis[c]; {
		case ok:
			stack.Push(p)
		case stack.Len() == 0 || c != stack.Pop().(rune):
			return false
		}
	}
	if stack.Len() != 0 {
		return false
	}
	return true
}
