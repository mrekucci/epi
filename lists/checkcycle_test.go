// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package lists

import "testing"

// createCycleList returns cycled linked list created from data.
// The cycle in the list is created to node on index ci of data.
// The returned node represents the reference to the start of the
// cycle or nil is returned when there is no cycle in returned list.
// If ci < 0 || ci >= len(data), then a list without cycle is returned.
func createCycleList(data []interface{}, ci int) (*List, *Node) {
	if ci < 0 || ci >= len(data) {
		return NewFromSlice(data), nil
	}

	l := new(List)
	var csn *Node
	for i := 0; i <= ci; i++ {
		csn = l.Insert(&Node{Data: data[i]})
	}
	ln := csn
	for i := ci + 1; i < len(data); i++ {
		ln = l.Insert(&Node{Data: data[i]})
	}
	ln.next = csn // Create a cycle from the ln to the csn.

	return l, csn
}

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
		l, want := createCycleList(test.l, test.i)
		if got := fn(l); got != want {
			t.Errorf("%s(%v) = %v; want %v", fnName, test.l, got, want)
		}
	}
}

func TestHasCycle(t *testing.T)    { testHasCycleFn(t, HasCycle, "HasCycle") }
func TestHasCycleAlt(t *testing.T) { testHasCycleFn(t, HasCycleAlt, "HasCycleAlt") }

func benchHasCycleFn(b *testing.B, size int, fn hasCycleFn, fnName string) {
	b.StopTimer()
	data := make([]interface{}, size) // We don't care about content but about pointers.
	for i := 0; i < b.N; i++ {
		l, n := createCycleList(data, 0)
		b.StartTimer()
		csn := fn(l)
		b.StopTimer()
		if n != csn {
			b.Errorf("%s did not find the cycle", fnName)
		}
	}
}

func BenchmarkHasCycle1e2(b *testing.B)    { benchHasCycleFn(b, 1e2, HasCycle, "HasCycle") }
func BenchmarkHasCycleAlt1e2(b *testing.B) { benchHasCycleFn(b, 1e2, HasCycleAlt, "HasCycleAlt") }
func BenchmarkHasCycle2e2(b *testing.B)    { benchHasCycleFn(b, 2e2, HasCycle, "HasCycle") }
func BenchmarkHasCycleAlt2e2(b *testing.B) { benchHasCycleFn(b, 2e2, HasCycleAlt, "HasCycleAlt") }
func BenchmarkHasCycle3e2(b *testing.B)    { benchHasCycleFn(b, 3e2, HasCycle, "HasCycle") }
func BenchmarkHasCycleAlt3e2(b *testing.B) { benchHasCycleFn(b, 3e2, HasCycleAlt, "HasCycleAlt") }
