// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package sorting

import (
	"math/rand"
	"testing"

	"github.com/mrekucci/epi/internal/epiutil"
)

func TestCountOccurrences(t *testing.T) {
	for _, test := range []struct {
		in   string
		want string
	}{
		{"", ""},
		{"0", "(0, 1)"},
		{"a", "(a, 1)"},
		{"ba", "(a, 1), (b, 1)"},
		{"bcdacebe", "(a, 1), (b, 2), (c, 2), (d, 1), (e, 2)"},
		{"bcdacebe", "(a, 1), (b, 2), (c, 2), (d, 1), (e, 2)"},
		{"界☺世界☺世界☺世", "(☺, 3), (世, 3), (界, 3)"},
	} {
		if got := CountOccurrences(test.in); got != test.want {
			t.Errorf("CountOccurrences(%q) = %q; want %q", test.in, got, test.want)
		}
	}
}

func benchCountOccurrences(b *testing.B, size int) {
	s := epiutil.RandStr(size, "0123456789abcdefghijklmnopqrstuvwxyz", rand.NewSource(int64(size)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CountOccurrences(s)
	}
}

func BenchmarkCountOccurrences1e2(b *testing.B) { benchCountOccurrences(b, 1e2) }
func BenchmarkCountOccurrences1e4(b *testing.B) { benchCountOccurrences(b, 1e4) }
func BenchmarkCountOccurrences1e6(b *testing.B) { benchCountOccurrences(b, 1e6) }
