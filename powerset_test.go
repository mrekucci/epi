// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package epi

import (
	"math/rand"
	"reflect"
	"testing"
)

var powerSetTests = []struct {
	in   []interface{}
	want []interface{}
	ok   bool
}{
	{
		[]interface{}{"A", "B"},
		[]interface{}{
			[]interface{}(nil),
			[]interface{}{"A"},
			[]interface{}{"B"},
			[]interface{}{"A", "B"}},
		true,
	},
	{
		[]interface{}{"A", "B", "C"},
		[]interface{}{
			[]interface{}(nil),
			[]interface{}{"A"},
			[]interface{}{"B"},
			[]interface{}{"A", "B"},
			[]interface{}{"C"},
			[]interface{}{"A", "C"},
			[]interface{}{"B", "C"},
			[]interface{}{"A", "B", "C"}},
		true,
	},
	{
		[]interface{}(nil), // Will be initialized in init function to 32 or 64 elements according to architecture.
		[]interface{}(nil),
		false,
	},
}

func init() {
	for i := range powerSetTests {
		test := &powerSetTests[i]
		if !test.ok {
			l := intSize
			for i := 0; i < l; i++ {
				test.in = append(test.in, "0")
			}
		}
	}
}

func TestPowerSet(t *testing.T) {
	for _, tt := range powerSetTests {
		got, ok := PowerSet(tt.in)
		if !reflect.DeepEqual(got, tt.want) || ok != tt.ok {
			t.Errorf("PowerSet(%#v) = %#v, %t; want %#v, %t", tt.in, got, ok, tt.want, tt.ok)
		}
	}
}

func benchPowerSet(b *testing.B, size int) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		ints := rand.Perm(size)
		data := make([]interface{}, len(ints))
		for j := range data {
			data[j] = ints[j]
		}
		b.StartTimer()
		PowerSet(data)
		b.StopTimer()
	}
}

func BenchmarkPowerSetIntSizeDiv8(b *testing.B) { benchPowerSet(b, intSize/8) }
func BenchmarkPowerSetIntSizeDiv6(b *testing.B) { benchPowerSet(b, intSize/6) }
func BenchmarkPowerSetIntSizeDiv4(b *testing.B) { benchPowerSet(b, intSize/4) }
