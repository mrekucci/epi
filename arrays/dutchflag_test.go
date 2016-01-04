// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package arrays

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestRearrange(t *testing.T) {
	for _, test := range []struct {
		an   [5]int
		i    int
		want [5]int
	}{
		{[...]int{5, 4, 3, 2, 1}, 0, [...]int{4, 3, 2, 1, 5}},
		{[...]int{5, 4, 3, 2, 1}, 1, [...]int{1, 3, 2, 4, 5}},
		{[...]int{5, 4, 3, 2, 1}, 2, [...]int{1, 2, 3, 4, 5}},
		{[...]int{5, 4, 3, 2, 1}, 3, [...]int{1, 2, 3, 4, 5}},
		{[...]int{5, 4, 3, 2, 1}, 4, [...]int{1, 3, 2, 4, 5}},
		{[...]int{4, 3, 3, 5, 5}, 0, [...]int{3, 3, 4, 5, 5}},
	} {
		got := test.an
		Rearrange(got[:], test.i)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Rearrange(%d, %d) got %d; want %d", test.an, test.i, got, test.want)
		}
	}
}

func benchRearrange(b *testing.B, size int) {
	b.StopTimer()
	p, q := size/3, size
	ints := rand.New(rand.NewSource(int64(size))).Perm(size)
	data := make([]int, size)
	for i := 0; i < b.N; i++ {
		copy(data, ints)
		j := data[i%size]%(q-p) + p
		b.StartTimer()
		Rearrange(data, j)
		b.StopTimer()
	}
}

func BenchmarkRearrange1e2(b *testing.B) { benchRearrange(b, 1e3) }
func BenchmarkRearrange1e4(b *testing.B) { benchRearrange(b, 1e5) }
func BenchmarkRearrange1e6(b *testing.B) { benchRearrange(b, 1e7) }
