// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package lists

import (
	"math/rand"
	"reflect"
	"sort"
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
		if test.ok && (l.head != nil || f.head != nil) {
			t.Errorf("MergeSorted(%v, %v) got l.First() = %v, f.First() = %v; want <nil>, <nil>", test.l, test.f, l.First(), f.First())
		}
		if got := m.ToSlice(); ok != test.ok || !reflect.DeepEqual(got, test.want) {
			t.Errorf("MergeSorted(%v, %v) = %v, %t; want %v, %t", test.l, test.f, got, ok, test.want, test.ok)
		}
	}
}

func benchMergeSorted(b *testing.B, size int) {
	b.StopTimer()
	s := make([]interface{}, size)
	s1 := s[:size/2]
	s2 := s[size/2:]
	data := rand.New(rand.NewSource(int64(size))).Perm(size)
	sort.Ints(data)
	for i := 1; i < size; i += 2 {
		s1[i/2] = data[i-1]
		s2[i/2] = data[i]
	}
	for i := 0; i < b.N; i++ {
		l := NewFromSlice(s1)
		f := NewFromSlice(s2)
		b.StartTimer()
		MergeSorted(l, f)
		b.StopTimer()
	}
}

func BenchmarkMergeSorted1e2(b *testing.B) { benchMergeSorted(b, 1e2) }
func BenchmarkMergeSorted1e3(b *testing.B) { benchMergeSorted(b, 1e3) }
func BenchmarkMergeSorted1e4(b *testing.B) { benchMergeSorted(b, 1e4) }
