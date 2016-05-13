// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package strings

import (
	"math/rand"
	"testing"
	"time"

	"github.com/mrekucci/epi/internal/epiutil"
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

func benchIndexFn(b *testing.B, size int, fn indexFn) {
	s := epiutil.RandStr(size, "☺ abcdefghijklmnopqrstuvwxyz 世界", rand.NewSource(time.Now().UnixNano()))
	p := s[size/3 : size*2/3]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fn(s, p)
	}
}

func BenchmarkIndexNaive1e4(b *testing.B) { benchIndexFn(b, 1e4, IndexNaive) }
func BenchmarkIndexRK1e4(b *testing.B)    { benchIndexFn(b, 1e4, IndexRK) }
func BenchmarkIndexNaive1e6(b *testing.B) { benchIndexFn(b, 1e6, IndexNaive) }
func BenchmarkIndexRK1e6(b *testing.B)    { benchIndexFn(b, 1e6, IndexRK) }
func BenchmarkIndexNaive1e8(b *testing.B) { benchIndexFn(b, 1e8, IndexNaive) }
func BenchmarkIndexRK1e8(b *testing.B)    { benchIndexFn(b, 1e8, IndexRK) }
