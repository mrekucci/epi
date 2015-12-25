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
				&Symbol{rune('a'), 8.17, ""}},
			[]*Symbol{
				&Symbol{rune('a'), 8.17, ""}},
		},
		{
			[]*Symbol{
				&Symbol{rune('a'), 8.17, ""},
				&Symbol{rune('b'), 1.49, ""},
				&Symbol{rune('c'), 2.78, ""}},
			[]*Symbol{
				&Symbol{rune('a'), 8.17, "1"},
				&Symbol{rune('b'), 1.49, "00"},
				&Symbol{rune('c'), 2.78, "01"}},
		},
		{
			[]*Symbol{
				&Symbol{rune('a'), 8.17, ""},
				&Symbol{rune('b'), 1.49, ""},
				&Symbol{rune('c'), 2.78, ""},
				&Symbol{rune('d'), 4.253, ""},
				&Symbol{rune('e'), 12.702, ""},
				&Symbol{rune('f'), 2.228, ""},
				&Symbol{rune('g'), 2.015, ""},
				&Symbol{rune('h'), 6.094, ""},
				&Symbol{rune('i'), 6.966, ""},
				&Symbol{rune('j'), 0.153, ""},
				&Symbol{rune('k'), 0.772, ""},
				&Symbol{rune('l'), 4.025, ""},
				&Symbol{rune('m'), 2.406, ""},
				&Symbol{rune('n'), 6.749, ""},
				&Symbol{rune('o'), 7.507, ""},
				&Symbol{rune('p'), 1.929, ""},
				&Symbol{rune('q'), 0.095, ""},
				&Symbol{rune('r'), 5.987, ""},
				&Symbol{rune('s'), 6.327, ""},
				&Symbol{rune('t'), 9.056, ""},
				&Symbol{rune('u'), 2.758, ""},
				&Symbol{rune('v'), 0.978, ""},
				&Symbol{rune('w'), 2.360, ""},
				&Symbol{rune('x'), 0.150, ""},
				&Symbol{rune('y'), 1.974, ""},
				&Symbol{rune('z'), 0.074, ""}},
			[]*Symbol{
				&Symbol{rune('a'), 8.17, "1110"},
				&Symbol{rune('b'), 1.49, "110000"},
				&Symbol{rune('c'), 2.78, "01001"},
				&Symbol{rune('d'), 4.253, "11111"},
				&Symbol{rune('e'), 12.702, "100"},
				&Symbol{rune('f'), 2.228, "00101"},
				&Symbol{rune('g'), 2.015, "110011"},
				&Symbol{rune('h'), 6.094, "0110"},
				&Symbol{rune('i'), 6.966, "1011"},
				&Symbol{rune('j'), 0.153, "001001011"},
				&Symbol{rune('k'), 0.772, "0010011"},
				&Symbol{rune('l'), 4.025, "11110"},
				&Symbol{rune('m'), 2.406, "00111"},
				&Symbol{rune('n'), 6.749, "1010"},
				&Symbol{rune('o'), 7.507, "1101"},
				&Symbol{rune('p'), 1.929, "110001"},
				&Symbol{rune('q'), 0.095, "001001001"},
				&Symbol{rune('r'), 5.987, "0101"},
				&Symbol{rune('s'), 6.327, "0111"},
				&Symbol{rune('t'), 9.056, "000"},
				&Symbol{rune('u'), 2.758, "01000"},
				&Symbol{rune('v'), 0.978, "001000"},
				&Symbol{rune('w'), 2.360, "00110"},
				&Symbol{rune('x'), 0.150, "001001010"},
				&Symbol{rune('y'), 1.974, "110010"},
				&Symbol{rune('z'), 0.074, "001001000"}},
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
			&Symbol{rune('a'), 8.17, ""},
			&Symbol{rune('b'), 1.49, ""},
			&Symbol{rune('c'), 2.78, ""},
			&Symbol{rune('d'), 4.253, ""},
			&Symbol{rune('e'), 12.702, ""},
			&Symbol{rune('f'), 2.228, ""},
			&Symbol{rune('g'), 2.015, ""},
			&Symbol{rune('h'), 6.094, ""},
			&Symbol{rune('i'), 6.966, ""},
			&Symbol{rune('j'), 0.153, ""},
			&Symbol{rune('k'), 0.772, ""},
			&Symbol{rune('l'), 4.025, ""},
			&Symbol{rune('m'), 2.406, ""},
			&Symbol{rune('n'), 6.749, ""},
			&Symbol{rune('o'), 7.507, ""},
			&Symbol{rune('p'), 1.929, ""},
			&Symbol{rune('q'), 0.095, ""},
			&Symbol{rune('r'), 5.987, ""},
			&Symbol{rune('s'), 6.327, ""},
			&Symbol{rune('t'), 9.056, ""},
			&Symbol{rune('u'), 2.758, ""},
			&Symbol{rune('v'), 0.978, ""},
			&Symbol{rune('w'), 2.360, ""},
			&Symbol{rune('x'), 0.150, ""},
			&Symbol{rune('y'), 1.974, ""},
			&Symbol{rune('z'), 0.074, ""}})
	}
}
