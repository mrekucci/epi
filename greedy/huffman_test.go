// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package greedy

import (
	"reflect"
	"testing"
)

func TestHuffmanEncoding(t *testing.T) {
	for _, test := range []struct {
		in   []*Symbol
		want []*Symbol
	}{
		{nil, nil},
		{
			[]*Symbol{
				&Symbol{'a', 8.17, ""}},
			[]*Symbol{
				&Symbol{'a', 8.17, ""}},
		},
		{
			[]*Symbol{
				&Symbol{'a', 8.17, ""},
				&Symbol{'b', 1.49, ""},
				&Symbol{'c', 2.78, ""}},
			[]*Symbol{
				&Symbol{'a', 8.17, "1"},
				&Symbol{'b', 1.49, "00"},
				&Symbol{'c', 2.78, "01"}},
		},
		{
			[]*Symbol{
				&Symbol{'a', 8.17, ""},
				&Symbol{'b', 1.49, ""},
				&Symbol{'c', 2.78, ""},
				&Symbol{'d', 4.253, ""},
				&Symbol{'e', 12.702, ""},
				&Symbol{'f', 2.228, ""},
				&Symbol{'g', 2.015, ""},
				&Symbol{'h', 6.094, ""},
				&Symbol{'i', 6.966, ""},
				&Symbol{'j', 0.153, ""},
				&Symbol{'k', 0.772, ""},
				&Symbol{'l', 4.025, ""},
				&Symbol{'m', 2.406, ""},
				&Symbol{'n', 6.749, ""},
				&Symbol{'o', 7.507, ""},
				&Symbol{'p', 1.929, ""},
				&Symbol{'q', 0.095, ""},
				&Symbol{'r', 5.987, ""},
				&Symbol{'s', 6.327, ""},
				&Symbol{'t', 9.056, ""},
				&Symbol{'u', 2.758, ""},
				&Symbol{'v', 0.978, ""},
				&Symbol{'w', 2.360, ""},
				&Symbol{'x', 0.150, ""},
				&Symbol{'y', 1.974, ""},
				&Symbol{'z', 0.074, ""}},
			[]*Symbol{
				&Symbol{'a', 8.17, "1110"},
				&Symbol{'b', 1.49, "110000"},
				&Symbol{'c', 2.78, "01001"},
				&Symbol{'d', 4.253, "11111"},
				&Symbol{'e', 12.702, "100"},
				&Symbol{'f', 2.228, "00101"},
				&Symbol{'g', 2.015, "110011"},
				&Symbol{'h', 6.094, "0110"},
				&Symbol{'i', 6.966, "1011"},
				&Symbol{'j', 0.153, "001001011"},
				&Symbol{'k', 0.772, "0010011"},
				&Symbol{'l', 4.025, "11110"},
				&Symbol{'m', 2.406, "00111"},
				&Symbol{'n', 6.749, "1010"},
				&Symbol{'o', 7.507, "1101"},
				&Symbol{'p', 1.929, "110001"},
				&Symbol{'q', 0.095, "001001001"},
				&Symbol{'r', 5.987, "0101"},
				&Symbol{'s', 6.327, "0111"},
				&Symbol{'t', 9.056, "000"},
				&Symbol{'u', 2.758, "01000"},
				&Symbol{'v', 0.978, "001000"},
				&Symbol{'w', 2.360, "00110"},
				&Symbol{'x', 0.150, "001001010"},
				&Symbol{'y', 1.974, "110010"},
				&Symbol{'z', 0.074, "001001000"}},
		},
	} {
		got := append([]*Symbol(nil), test.in...) // Make a copy to avoid modification of the original.
		if HuffmanEncoding(got); !reflect.DeepEqual(got, test.want) {
			t.Errorf("HuffmanEncoding(%v): got %v; want %v", test.in, got, test.want)
			for i, s := range got {
				if !reflect.DeepEqual(s, test.want[i]) {
					t.Errorf("symbol %q got %v; want %v", s.c, s, test.want[i])
				}
			}
		}
	}
}

func BenchmarkHuffmanEncoding(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HuffmanEncoding([]*Symbol{
			&Symbol{'a', 8.17, ""},
			&Symbol{'b', 1.49, ""},
			&Symbol{'c', 2.78, ""},
			&Symbol{'d', 4.253, ""},
			&Symbol{'e', 12.702, ""},
			&Symbol{'f', 2.228, ""},
			&Symbol{'g', 2.015, ""},
			&Symbol{'h', 6.094, ""},
			&Symbol{'i', 6.966, ""},
			&Symbol{'j', 0.153, ""},
			&Symbol{'k', 0.772, ""},
			&Symbol{'l', 4.025, ""},
			&Symbol{'m', 2.406, ""},
			&Symbol{'n', 6.749, ""},
			&Symbol{'o', 7.507, ""},
			&Symbol{'p', 1.929, ""},
			&Symbol{'q', 0.095, ""},
			&Symbol{'r', 5.987, ""},
			&Symbol{'s', 6.327, ""},
			&Symbol{'t', 9.056, ""},
			&Symbol{'u', 2.758, ""},
			&Symbol{'v', 0.978, ""},
			&Symbol{'w', 2.360, ""},
			&Symbol{'x', 0.150, ""},
			&Symbol{'y', 1.974, ""},
			&Symbol{'z', 0.074, ""}})
	}
}
