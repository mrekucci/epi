// Copyright (c) 2016, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package heaps

import (
	"math/big"
	"reflect"
	"testing"
)

func TestMedianStream(t *testing.T) {
	in, out := make(chan int), make(chan *big.Rat)
	go MedianStream(in, out)
	for _, p := range []struct {
		in   int
		want *big.Rat
	}{
		{1, big.NewRat(1, 1)},
		{0, big.NewRat(1+0, 2)},
		{3, big.NewRat(1, 1)},
		{5, big.NewRat(1+3, 2)},
		{2, big.NewRat(2, 1)},
		{0, big.NewRat(1+2, 2)},
		{1, big.NewRat(1, 1)},
	} {
		in <- p.in
		got := <-out
		if !reflect.DeepEqual(got, p.want) {
			t.Errorf("%v MedianStream: got %v; want %v", p, got, p.want)
		}
	}
	close(in)
}

func benchMedianStream(b *testing.B, size int) {
	in, out := make(chan int), make(chan *big.Rat)
	go MedianStream(in, out)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for v := 0; v < size; v++ {
			in <- v
			<-out
		}
	}
}

func BenchmarkMedian1e2(b *testing.B) { benchMedianStream(b, 1e2) }
func BenchmarkMedian1e4(b *testing.B) { benchMedianStream(b, 1e4) }
func BenchmarkMedian1e6(b *testing.B) { benchMedianStream(b, 1e6) }
