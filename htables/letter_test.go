// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package htables

import (
	"math/rand"
	"testing"

	"github.com/mrekucci/epi/internal/epiutil"
)

func TestIsLetterSubset(t *testing.T) {
	for _, test := range []struct {
		l, m string
		want bool
	}{
		{" ", "", false},
		{"a", "", false},
		{"", "a", true},
		{"abc", "cba", true},
		{"Hello World", "Hello New World, and Googbye Old World", true},
		{"Hello New World, and Googbye Old World", "Hello World", false},
	} {
		if got := IsLetterSubset(test.l, test.m); got != test.want {
			t.Errorf("IsLetterConstructible(%q, %q) = %t; want %t", test.l, test.m, got, test.want)
		}
	}

}

func benchIsLetterSubset(b *testing.B, size int) {
	m := epiutil.RandStr(size, "ab cd ef gh ij kl mn op qr st uv wx yz", rand.NewSource(int64(size)))
	l := m[:size/2]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsLetterSubset(l, m)
	}
}

func BenchmarkIsLetterSubset1e2(b *testing.B) { benchIsLetterSubset(b, 1e2) }
func BenchmarkIsLetterSubset1e4(b *testing.B) { benchIsLetterSubset(b, 1e4) }
func BenchmarkIsLetterSubset1e6(b *testing.B) { benchIsLetterSubset(b, 1e6) }
