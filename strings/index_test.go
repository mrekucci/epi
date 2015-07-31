// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package strings

import (
	"math/rand"
	"testing"

	"github.com/mrekucci/epi/epiutil"
)

type indexFn func(s, p string) int

func testIndexFn(t *testing.T, fn indexFn, fnName string) {
	for _, test := range []struct {
		s, p string
		want int
	}{
		{"", "", 0},
		{"", "x", -1},
		{"", "xxx", -1},
		{"xx", "xxx", -1},
		{"xxx", "xxx", 0},
		{"xxYxYxxYxxx", "Y", 2},
		{"xxYxYxxYxxx", "Yxx", 4},
		{"xxxyy", "xxyy", 1},
		{"xx☺yy", "x☺y", 1},
		{"xx世界yy", "世界", 2},
		{"xxxyyyzzz", "yyy", 3},
		{"xxYxYYxYxxYxxx", "Yxx", 7},
		{"xyz", "", 0},
		{"xyz", "y", 1},
		{"x", "y", -1},
		{"x", "x", 0},
		{"xyz", "x", 0},
		{"xyz", "y", 1},
		{"xyz", "z", 2},
		{"xyz", "-", -1},
	} {
		if got := fn(test.s, test.p); got != test.want {
			t.Errorf("%s(%q, %q) = %d; want %d", fnName, test.s, test.p, got, test.want)
		}
	}
}

func TestIndexNaive(t *testing.T) { testIndexFn(t, IndexNaive, "IndexNaive") }
func TestIndexRK(t *testing.T)    { testIndexFn(t, IndexRK, "IndexRK") }

func benchIndexFn(b *testing.B, size int, fn indexFn, fnName string) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		s := epiutil.RandStr(size, "☺ abcdefghijklmnopqrstuvwxyz 世界", rand.NewSource(int64(i)))
		o, c := size/3, size*2/3
		p := s[o:c]
		b.StartTimer()
		r := fn(s, p)
		b.StopTimer()
		if r != o {
			b.Errorf("%s did not find the index of the substring", fnName)
		}
	}
}

func BenchmarkIndexNaive1e4(b *testing.B) { benchIndexFn(b, 1e4, IndexNaive, "IndexNaive") }
func BenchmarkIndexRK1e4(b *testing.B)    { benchIndexFn(b, 1e4, IndexRK, "IndexRK") }
func BenchmarkIndexNaive1e6(b *testing.B) { benchIndexFn(b, 1e6, IndexNaive, "IndexNaive") }
func BenchmarkIndexRK1e6(b *testing.B)    { benchIndexFn(b, 1e6, IndexRK, "IndexRK") }
func BenchmarkIndexNaive1e8(b *testing.B) { benchIndexFn(b, 1e8, IndexNaive, "IndexNaive") }
func BenchmarkIndexRK1e8(b *testing.B)    { benchIndexFn(b, 1e8, IndexRK, "IndexRK") }
