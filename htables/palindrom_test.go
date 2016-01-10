// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package htables

import (
	"math/rand"
	"testing"

	"github.com/mrekucci/epi/internal/epiutil"
)

func TestCanFormPalindrome(t *testing.T) {
	for _, test := range []struct {
		in   string
		want bool
	}{
		{"", true},
		{"ab", false},
		{"aab", true},
		{"aaab", false},
		{"aaabccc", false},
		{"aaabcc", false},
		{"edified", true},
		{"界☺世", false},
		{"世世☺", true},
	} {
		if got := CanFormPalindrome(test.in); got != test.want {
			t.Errorf("CanFormPalindrome(%q) = %t; want %t", test.in, got, test.want)
		}
	}
}

func benchCanFormPalindrome(b *testing.B, size int) {
	s := epiutil.RandStr(size, "abcdefghijklmnopqrstuvwxyz", rand.NewSource(int64(size)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CanFormPalindrome(s)
	}
}

func BenchmarkCanFormPalindrome1e2(b *testing.B) { benchCanFormPalindrome(b, 1e2) }
func BenchmarkCanFormPalindrome1e4(b *testing.B) { benchCanFormPalindrome(b, 1e4) }
func BenchmarkCanFormPalindrome1e6(b *testing.B) { benchCanFormPalindrome(b, 1e6) }
