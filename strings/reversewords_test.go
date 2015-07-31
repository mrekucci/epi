// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package strings

import (
	"math/rand"
	"testing"

	"github.com/mrekucci/epi/epiutil"
)

func TestReverseWords(t *testing.T) {
	for _, test := range []struct {
		in   string
		want string
	}{
		{},
		{"a", "a"},
		{"b ", " b"},
		{"a b c", "c b a"},
		{"a b c d", "d c b a"},
		{"Alice like Bob", "Bob like Alice"},
		{"The quick brown fox jumps over the lazy dog", "dog lazy the over jumps fox brown quick The"},

		// Test more then one space.
		{"aa  bb", "bb  aa"},
	} {
		if got := ReverseWords(test.in); got != test.want {
			t.Errorf("ReverseWorlds(%q) = %q; want %q", test.in, got, test.want)
		}
	}
}

func benchReverseWords(b *testing.B, size int) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		s := epiutil.RandStr(size, "The quick brown fox jumps over the lazy dog", rand.NewSource(int64(i))) // Pangram.
		b.StartTimer()
		ReverseWords(s)
		b.StopTimer()
	}
}

func BenchmarkReverseWords1e1(b *testing.B) { benchReverseWords(b, 1e1) }
func BenchmarkReverseWords1e2(b *testing.B) { benchReverseWords(b, 1e2) }
func BenchmarkReverseWords1e3(b *testing.B) { benchReverseWords(b, 1e3) }
