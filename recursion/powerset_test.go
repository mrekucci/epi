// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package recursion

import (
	"math/rand"
	"reflect"
	"testing"
)

type powerSetFn func(s []interface{}) ([][]interface{}, bool)

type powerSetTest struct {
	in   []interface{}
	want [][]interface{}
	ok   bool
}

func testPowerSetFn(t *testing.T, fn powerSetFn, fnName string, tests []powerSetTest) {
	for _, test := range tests {
		if got, ok := fn(test.in); !reflect.DeepEqual(got, test.want) || ok != test.ok {
			t.Errorf("%s(%v) = %v, %t; want %v, %t", fnName, test.in, got, ok, test.want, test.ok)
		}
	}
}

func TestPowerSetRec(t *testing.T) {
	tests := []powerSetTest{
		{[]interface{}{"A", "B"}, [][]interface{}{{"A", "B"}, {"A"}, {"B"}, nil}, true},
		{[]interface{}{"A", "B", "C"}, [][]interface{}{{"A", "B", "C"}, {"A", "B"}, {"A", "C"}, {"A"}, {"B", "C"}, {"B"}, {"C"}, nil}, true},
		{[]interface{}(nil), [][]interface{}{nil}, true}}
	testPowerSetFn(t, PowerSetRec, "PowerSetRec", tests)
}

func TestPowerSetItr(t *testing.T) {
	tests := []powerSetTest{
		{[]interface{}{"A", "B"}, [][]interface{}{nil, {"A"}, {"B"}, {"A", "B"}}, true},
		{[]interface{}{"A", "B", "C"}, [][]interface{}{nil, {"A"}, {"B"}, {"A", "B"}, {"C"}, {"A", "C"}, {"B", "C"}, {"A", "B", "C"}}, true},
		{[]interface{}(nil), [][]interface{}{nil}, true},
		{make([]interface{}, intSize), [][]interface{}(nil), false}}
	testPowerSetFn(t, PowerSetItr, "PowerSetItr", tests)
}

func benchPowerSetFn(b *testing.B, size int, fn powerSetFn) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]interface{}, size)
		for j, p := range rand.New(rand.NewSource(int64(i))).Perm(size) {
			data[j] = p
		}
		b.StartTimer()
		fn(data)
		b.StopTimer()
	}
}

func BenchmarkPowerSetRec4(b *testing.B) { benchPowerSetFn(b, 4, PowerSetRec) }
func BenchmarkPowerSetItr4(b *testing.B) { benchPowerSetFn(b, 4, PowerSetItr) }
func BenchmarkPowerSetRec6(b *testing.B) { benchPowerSetFn(b, 6, PowerSetRec) }
func BenchmarkPowerSetItr6(b *testing.B) { benchPowerSetFn(b, 6, PowerSetItr) }
func BenchmarkPowerSetRec8(b *testing.B) { benchPowerSetFn(b, 8, PowerSetRec) }
func BenchmarkPowerSetItr8(b *testing.B) { benchPowerSetFn(b, 8, PowerSetItr) }
