// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package epi

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestPowerSet(t *testing.T) {
	for _, test := range []struct {
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
			[]interface{}(nil), // Will be initialized, before the test will run, to 32 or 64 elements (according to the architecture).
			[]interface{}(nil),
			false,
		},
	} {
		if !test.ok {
			l := intSize
			for i := 0; i < l; i++ {
				test.in = append(test.in, "0")
			}
		}

		if got, ok := PowerSet(test.in); !reflect.DeepEqual(got, test.want) || ok != test.ok {
			t.Errorf("PowerSet(%#v) = %#v, %t; want %#v, %t", test.in, got, ok, test.want, test.ok)
		}
	}
}

func benchPowerSet(b *testing.B, size int) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]interface{}, size)
		for j, p := range rand.New(rand.NewSource(int64(i))).Perm(size) {
			data[j] = p
		}
		b.StartTimer()
		PowerSet(data)
		b.StopTimer()
	}
}

func BenchmarkPowerSetIntSizeDiv8(b *testing.B) { benchPowerSet(b, intSize/8) }
func BenchmarkPowerSetIntSizeDiv6(b *testing.B) { benchPowerSet(b, intSize/6) }
func BenchmarkPowerSetIntSizeDiv4(b *testing.B) { benchPowerSet(b, intSize/4) }
