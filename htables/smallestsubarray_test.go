// Copyright (c) 2016, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package htables

import "testing"

type subarrayCoveringSetFn func([]string, map[string]bool) (int, int)

func testSubarrayCoveringSetFn(t *testing.T, fn subarrayCoveringSetFn, fnName string) {
	for _, test := range []struct {
		paragraph  []string
		keywords   map[string]bool
		start, end int
	}{
		{nil, map[string]bool{"none": true}, -1, -1},
		{[]string{"a", "b"}, nil, -1, -1},
		{[]string{"a", "b"}, map[string]bool{"none": true}, -1, -1},
		{[]string{"a", "b"}, map[string]bool{"a": true, "b": true}, 0, 1},
		{[]string{"a", "b"}, map[string]bool{"b": true, "a": true}, 0, 1},
		{[]string{"a", "b"}, map[string]bool{"a": true, "b": true, "c": true}, -1, -1},
		{
			[]string{"apple", "banana", "apple", "apple", "dog", "cat", "apple", "dog", "banana", "apple", "cat", "dog"},
			map[string]bool{"banana": true, "cat": true},
			8, 10,
		},
	} {
		start, end := fn(test.paragraph, test.keywords)
		if start != test.start || end != test.end {
			t.Errorf("%s(%v, %v) = %d, %d; want %d, %d",
				fnName, test.paragraph, test.keywords, start, end, test.start, test.end)
		}
	}
}

func TestSubarrayCoveringSet(t *testing.T) {
	testSubarrayCoveringSetFn(t, SubarrayCoveringSet, "SubarrayCoveringSet")
}

func TestSubarrayCoveringSetNaive(t *testing.T) {
	testSubarrayCoveringSetFn(t, SubarrayCoveringSetNaive, "SubarrayCoveringSetNaive")
}

// TODO: benchmark!
