// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package strings

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/mrekucci/epi/internal/epiutil"
)

func TestPhoneMnemonics(t *testing.T) {
	for _, test := range []struct {
		in   string
		want []string
		ok   bool
	}{
		{"", nil, false},
		{"0a1", nil, false},
		{"0", []string{"0"}, true},
		{"1", []string{"1"}, true},
		{"2", []string{"A", "B", "C"}, true},
		{"3", []string{"D", "E", "F"}, true},
		{"4", []string{"G", "H", "I"}, true},
		{"5", []string{"J", "K", "L"}, true},
		{"6", []string{"M", "N", "O"}, true},
		{"7", []string{"P", "Q", "R", "S"}, true},
		{"8", []string{"T", "U", "V"}, true},
		{"9", []string{"W", "X", "Y", "Z"}, true},
		{"12", []string{"1A", "1B", "1C"}, true},
		{"23", []string{"AD", "AE", "AF", "BD", "BE", "BF", "CD", "CE", "CF"}, true},
		{"33", []string{"DD", "DE", "DF", "ED", "EE", "EF", "FD", "FE", "FF"}, true},
		{"123", []string{"1AD", "1AE", "1AF", "1BD", "1BE", "1BF", "1CD", "1CE", "1CF"}, true},
	} {
		if got, ok := PhoneMnemonics(test.in); !reflect.DeepEqual(got, test.want) || ok != test.ok {
			t.Errorf("PhoneMnemonics(%q) = %q %t; want %q, %t", test.in, got, ok, test.want, test.ok)
		}
	}
}

func benchPhoneMnemonics(b *testing.B, size int) {
	number := epiutil.RandStr(size, "0123456789", rand.NewSource(int64(size)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PhoneMnemonics(number)
	}
}

func BenchmarkPhoneMnemonics1e0(b *testing.B) { benchPhoneMnemonics(b, 1e0) }
func BenchmarkPhoneMnemonics5e0(b *testing.B) { benchPhoneMnemonics(b, 5e0) }
func BenchmarkPhoneMnemonics1e1(b *testing.B) { benchPhoneMnemonics(b, 1e1) }
