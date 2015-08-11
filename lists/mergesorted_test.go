// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package lists

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestMergeSorted(t *testing.T) {
	const (
		maxInt = int(^uint(0) >> 1)
		minInt = -maxInt - 1
	)
	for _, test := range []struct {
		l, f []interface{}
		want []interface{}
		ok   bool
	}{
		{[]interface{}{0}, []interface{}{0}, []interface{}{0, 0}, true},
		{[]interface{}{0}, []interface{}{-1, 1}, []interface{}{-1, 0, 1}, true},
		{[]interface{}{2, 5, 7}, []interface{}{3, 11}, []interface{}{2, 3, 5, 7, 11}, true},
		{[]interface{}{1, 3, 5}, []interface{}{2, 4, 6}, []interface{}{1, 2, 3, 4, 5, 6}, true},
		{[]interface{}{maxInt}, []interface{}{minInt}, []interface{}{minInt, maxInt}, true},
		{[]interface{}{1, "b", 3}, []interface{}{0}, []interface{}{0, 1}, false},
		{[]interface{}{0}, []interface{}{1, "b", 3}, []interface{}{0, 1}, false},
	} {
		l, f := NewFromSlice(test.l), NewFromSlice(test.f)
		m, ok := MergeSorted(l, f)
		if got := m.ToSlice(); ok != test.ok || !reflect.DeepEqual(got, test.want) {
			t.Errorf("MergeSorted(%v, %v) = %v, %t; want %v, %t", test.l, test.f, got, ok, test.want, test.ok)
		}
	}
}

func benchMergeSorted(b *testing.B, size int) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		var s []interface{}
		for _, n := range rand.New(rand.NewSource(int64(i))).Perm(size) {
			s = append(s, n)
		}
		l := NewFromSlice(s[:size/2])
		f := NewFromSlice(s[size/2:])
		b.StartTimer()
		m, ok := MergeSorted(l, f)
		b.StopTimer()
		if !ok || !reflect.DeepEqual(m.ToSlice(), s) {
			b.Error("MergeSorted did not merge lists properly")
		}
	}
}

func BenchmarkMergeSorted1e1(b *testing.B) { benchMergeSorted(b, 1e1) }
func BenchmarkMergeSorted1e2(b *testing.B) { benchMergeSorted(b, 1e2) }
func BenchmarkMergeSorted1e3(b *testing.B) { benchMergeSorted(b, 1e3) }
