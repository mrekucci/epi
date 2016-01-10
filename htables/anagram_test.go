// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package htables

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/mrekucci/epi/internal/epiutil"
)

// compareGroups returns true if the content of the groups
// (regardless of the order) a and b is identical.
func compareGroups(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}
	b = append([][]string(nil), b...)
	for _, ea := range a {
		for i, eb := range b {
			if reflect.DeepEqual(ea, eb) {
				b[i], b = b[len(b)-1], b[:len(b)-1]
				break
			}
		}
	}
	if len(b) != 0 {
		return false
	}
	return true
}

func TestGroupAnagrams(t *testing.T) {
	for _, test := range []struct {
		in   []string
		want [][]string
	}{
		{[]string(nil), [][]string(nil)},
		{[]string{"a"}, [][]string(nil)},
		{[]string{"a", "b"}, [][]string(nil)},
		{[]string{"ab", "ba"}, [][]string{[]string{"ab", "ba"}}},
		{[]string{"debitcard", "elvis", "silent", "badcredit", "lives", "freedom", "listen", "levis", "money"},
			[][]string{
				[]string{"debitcard", "badcredit"},
				[]string{"elvis", "lives", "levis"},
				[]string{"silent", "listen"}}},
	} {
		if got := GroupAnagrams(test.in); !compareGroups(got, test.want) {
			t.Errorf("GroupAnagrams(%q)", test.in)
			t.Errorf("got  %q", got)
			t.Errorf("want %q", test.want)
		}
	}
}

func benchGroupAnagrams(b *testing.B, n, m int) {
	words := make([]string, n)
	for i := 0; i < n; i++ {
		words[i] = epiutil.RandStr(m, "abcdefghijklmnopqrstuvwxyz", rand.NewSource(int64(i)))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GroupAnagrams(words)
	}
}

func BenchmarkGroupAnagramsN1e1M10(b *testing.B) { benchGroupAnagrams(b, 1e1, 10) }
func BenchmarkGroupAnagramsN1e1M20(b *testing.B) { benchGroupAnagrams(b, 1e1, 20) }
func BenchmarkGroupAnagramsN1e1M30(b *testing.B) { benchGroupAnagrams(b, 1e1, 30) }
func BenchmarkGroupAnagramsN1e2M10(b *testing.B) { benchGroupAnagrams(b, 1e2, 10) }
func BenchmarkGroupAnagramsN1e2M20(b *testing.B) { benchGroupAnagrams(b, 1e2, 20) }
func BenchmarkGroupAnagramsN1e2M30(b *testing.B) { benchGroupAnagrams(b, 1e2, 30) }
func BenchmarkGroupAnagramsN1e3M10(b *testing.B) { benchGroupAnagrams(b, 1e3, 10) }
func BenchmarkGroupAnagramsN1e3M20(b *testing.B) { benchGroupAnagrams(b, 1e3, 20) }
func BenchmarkGroupAnagramsN1e3M30(b *testing.B) { benchGroupAnagrams(b, 1e3, 30) }
