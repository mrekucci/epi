// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package lists

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestEvenOddMerge(t *testing.T) {
	for _, test := range []struct {
		l    []interface{}
		want []interface{}
		ok   bool
	}{
		{nil, nil, true},
		{[]interface{}{0}, []interface{}{0}, true},
		{[]interface{}{1}, []interface{}{1}, true},
		{[]interface{}{1, 2}, []interface{}{1, 2}, false},
		{[]interface{}{0, 1}, []interface{}{0, 1}, true},
		{[]interface{}{0, 2, 4, 6}, []interface{}{0, 2, 4, 6}, true},
		{[]interface{}{0, 3, 5, 7}, []interface{}{0, 3, 5, 7}, true},
		{[]interface{}{0, 2, 5, 6}, []interface{}{0, 2, 6, 5}, true},
		{[]interface{}{0, 1, 3, 4}, []interface{}{0, 4, 1, 3}, true},
		{[]interface{}{0, 3, 5, 6}, []interface{}{0, 6, 3, 5}, true},
		{[]interface{}{0, 1, 2, 3, 4}, []interface{}{0, 2, 4, 1, 3}, true},
		{[]interface{}{0, -1, -2, 3, 4}, []interface{}{0, -2, 4, -1, 3}, true},
		{[]interface{}{0, 3, 4, -1, -2}, []interface{}{0, 4, -2, 3, -1}, true},
	} {
		l := NewFromSlice(test.l)
		ok := EvenOddMerge(l)
		if got := l.ToSlice(); ok != test.ok || !reflect.DeepEqual(got, test.want) {
			t.Errorf("NewFromSlice(%v) = %t updated list %v; want %t updated list %v", test.l, ok, got, test.ok, test.want)
		}
	}
}

func benchEvenOddMerge(b *testing.B, size int) {
	ints := make([]interface{}, size)
	ints[0] = 0
	for i, n := range rand.New(rand.NewSource(int64(size))).Perm(size - 1) {
		ints[i+1] = n // i+1 because the firs value in list has to be 0.
	}
	l := NewFromSlice(ints)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EvenOddMerge(l)
	}
}

func BenchmarkEvenOddMerge1e1(b *testing.B) { benchEvenOddMerge(b, 1e1) }
func BenchmarkEvenOddMerge1e3(b *testing.B) { benchEvenOddMerge(b, 1e3) }
func BenchmarkEvenOddMerge1e5(b *testing.B) { benchEvenOddMerge(b, 1e5) }
