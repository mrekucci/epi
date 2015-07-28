// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package epi

import "testing"

func TestRandStr(t *testing.T) {
	for _, test := range []struct {
		n int
		t string
	}{
		{1, ""},
		{1, "a"},
		{1, "abc"},
		{10, "ab"},
		{10, "abcdefghijklmnopqrstuvwxyz"},
		{100, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"},
	} {
		if got := len(randStr(test.n, test.t)); got != test.n {
			t.Errorf("len(randStr(%d, %q) = %d; want %d", test.n, test.t, got, test.n)
		}
	}
}

func benchRandStr(b *testing.B, size int) {
	for i := 0; i < b.N; i++ {
		randStr(size, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	}
}

func BenchmarkRandStr1e2(b *testing.B) { benchRandStr(b, 1e2) }
func BenchmarkRandStr1e4(b *testing.B) { benchRandStr(b, 1e4) }
func BenchmarkRandStr1e6(b *testing.B) { benchRandStr(b, 1e6) }
