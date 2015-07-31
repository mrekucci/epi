// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package strings

import (
	"math/rand"
	"testing"

	"github.com/mrekucci/epi/epiutil"
)

func TestRLEEncode(t *testing.T) {
	for _, test := range []struct {
		in   string
		want string
		ok   bool
	}{
		{"", "", true},
		{"a", "1a", true},
		{"A", "1A", true},
		{"aa", "2a", true},
		{"AA", "2A", true},
		{"ab", "1a1b", true},
		{"AB", "1A1B", true},
		{"aab", "2a1b", true},
		{"AAB", "2A1B", true},
		{"abb", "1a2b", true},
		{"ABB", "1A2B", true},
		{"aaaabcccaa", "4a1b3c2a", true},
		{"AAAABCCCAA", "4A1B3C2A", true},
		{"aaaaaaaaaaaaaaaaaaaab", "20a1b", true},
		{"AAAAAAAAAAAAAAAAAAAAb", "20A1b", true},

		{"1", "", false},
		{"a1", "", false},
	} {
		if got, ok := RLEEncode(test.in); got != test.want || ok != test.ok {
			t.Errorf("RLEEncode(%s) = %s, %t; want %s, %t", test.in, got, ok, test.want, test.ok)
		}
	}
}

func benchRLEEncode(b *testing.B, size int) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		s := epiutil.RandStr(size, "abcdefghijklmnopqrstuvwxyz", rand.NewSource(int64(i)))
		b.StartTimer()
		RLEEncode(s)
		b.StopTimer()
	}
}

func BenchmarkRLEEncode1e1(b *testing.B) { benchRLEEncode(b, 1e1) }
func BenchmarkRLEEncode1e2(b *testing.B) { benchRLEEncode(b, 1e2) }
func BenchmarkRLEEncode1e3(b *testing.B) { benchRLEEncode(b, 1e3) }

func TestRLEDecode(t *testing.T) {
	for _, test := range []struct {
		in   string
		want string
		ok   bool
	}{
		{"", "", true},
		{"1a", "a", true},
		{"1A", "A", true},
		{"2a", "aa", true},
		{"2A", "AA", true},
		{"1a1b", "ab", true},
		{"1A1B", "AB", true},
		{"2a1b", "aab", true},
		{"2A1B", "AAB", true},
		{"1a2b", "abb", true},
		{"1A2B", "ABB", true},
		{"4a1b3c2a", "aaaabcccaa", true},
		{"4A1B3C2A", "AAAABCCCAA", true},
		{"20a1b", "aaaaaaaaaaaaaaaaaaaab", true},
		{"20A1b", "AAAAAAAAAAAAAAAAAAAAb", true},

		{"1", "", false},
		{"a", "", false},
		{"a1", "", false},
		{"1ab", "", false},
	} {
		if got, ok := RLEDecode(test.in); got != test.want || ok != test.ok {
			t.Errorf("RLEDecode(%s) = %s, %t; want %s, %t", test.in, got, ok, test.want, test.ok)
		}
	}
}

func benchRLEDecode(b *testing.B, size int) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		s, ok := RLEEncode(epiutil.RandStr(size, "abcdefghijklmnopqrstuvwxyz", rand.NewSource(int64(i))))
		if !ok {
			b.Errorf("RLEEncode did not encode string properly")
		}
		b.StartTimer()
		RLEDecode(s)
		b.StopTimer()
	}
}

func BenchmarkRLEDecode1e1(b *testing.B) { benchRLEDecode(b, 1e1) }
func BenchmarkRLEDecode1e2(b *testing.B) { benchRLEDecode(b, 1e2) }
func BenchmarkRLEDecode1e3(b *testing.B) { benchRLEDecode(b, 1e3) }
