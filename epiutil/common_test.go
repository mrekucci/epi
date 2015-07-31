// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package epiutil

import (
	"math/rand"
	"testing"
)

func TestRandStr(t *testing.T) {
	for _, test := range []struct {
		n int
		t string
		s rand.Source
	}{
		{1, "", rand.NewSource(1)},
		{1, "a", rand.NewSource(2)},
		{1, "abc", rand.NewSource(3)},
		{10, "ab", rand.NewSource(4)},
		{10, "abcdefghijklmnopqrstuvwxyz", rand.NewSource(5)},
		{100, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", rand.NewSource(6)},
	} {
		if got := len(RandStr(test.n, test.t, test.s)); got != test.n {
			t.Errorf("len(randStr(%d, %q, %v) = %d; want %d", test.n, test.t, test.s, got, test.n)
		}
	}
}

func benchRandStr(b *testing.B, size int) {
	for i := 0; i < b.N; i++ {
		RandStr(size, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", rand.NewSource(int64(i)))
	}
}

func BenchmarkRandStr1e2(b *testing.B) { benchRandStr(b, 1e2) }
func BenchmarkRandStr1e4(b *testing.B) { benchRandStr(b, 1e4) }
func BenchmarkRandStr1e6(b *testing.B) { benchRandStr(b, 1e6) }
