// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package lists

import "testing"

type hasCycleFn func(l *List) *Node

func testHasCycleFn(t *testing.T, fn hasCycleFn, fnName string) {
	for _, test := range []struct {
		l []interface{}
		i int
	}{
		// Test empty list.
		{[]interface{}{}, -1},

		// Test lists without cycle.
		{[]interface{}{0}, -1},
		{[]interface{}{0, 1, 2, 3, 4, 5, 6}, -1},
		{[]interface{}{0, 1, 2, 3, 4, 5, 6}, 7},

		// Test lists with cycle.
		{[]interface{}{0}, 0},
		{[]interface{}{0, 1, 2, 3, 4, 5, 6}, 0},
		{[]interface{}{0, 1, 2, 3, 4, 5, 6}, 3},
		{[]interface{}{0, 1, 2, 3, 4, 5, 6}, 6},
	} {
		l, want := CreateCycle(test.l, test.i)
		if got := fn(l); got != want {
			t.Errorf("%s(%v) = %v; want %v", fnName, test.l, got, want)
		}
	}
}

func TestHasCycle(t *testing.T)    { testHasCycleFn(t, HasCycle, "HasCycle") }
func TestHasCycleAlt(t *testing.T) { testHasCycleFn(t, HasCycleAlt, "HasCycleAlt") }

func benchHasCycleFn(b *testing.B, size int, fn hasCycleFn, fnName string) {
	data := make([]interface{}, size) // We don't care about content but about pointers.
	l, _ := CreateCycle(data, 0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fn(l)
	}
}

func BenchmarkHasCycle1e2(b *testing.B)    { benchHasCycleFn(b, 1e2, HasCycle, "HasCycle") }
func BenchmarkHasCycleAlt1e2(b *testing.B) { benchHasCycleFn(b, 1e2, HasCycleAlt, "HasCycleAlt") }
func BenchmarkHasCycle2e2(b *testing.B)    { benchHasCycleFn(b, 2e2, HasCycle, "HasCycle") }
func BenchmarkHasCycleAlt2e2(b *testing.B) { benchHasCycleFn(b, 2e2, HasCycleAlt, "HasCycleAlt") }
func BenchmarkHasCycle3e2(b *testing.B)    { benchHasCycleFn(b, 3e2, HasCycle, "HasCycle") }
func BenchmarkHasCycleAlt3e2(b *testing.B) { benchHasCycleFn(b, 3e2, HasCycleAlt, "HasCycleAlt") }
