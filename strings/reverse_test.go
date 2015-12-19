// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package strings

import (
	"github.com/mrekucci/epi/internal/epiutil"
	"math/rand"
	"testing"
)

type reverseFn func(string) string

func testReverseFn(t *testing.T, fn reverseFn, fnName string) {
	for _, test := range []struct {
		in   string
		want string
	}{
		{"-1", "1-"},
		{"12345", "54321"},
		{"Hello world", "dlrow olleH"},
		{"你好世界", "界世好你"},
		{"γειά σου κόσμος", "ςομσόκ υοσ άιεγ"},
		{"こんにちは世界", "界世はちにんこ"},
		{"여보세요 세계", "계세 요세보여"},
		{"Здравствулте мир", "рим етлувтсвардЗ"},
	} {
		if got := fn(test.in); got != test.want {
			t.Errorf("%s(%q) = %q; want %q", fnName, test.in, got, test.want)
		}
	}
}

func TestReverseItr(t *testing.T)     { testReverseFn(t, ReverseItr, "ReverseItr") }
func TestReverseRecAux(t *testing.T)  { testReverseFn(t, ReverseRecAux, "ReverseRecAux") }
func TestReverseRecPure(t *testing.T) { testReverseFn(t, ReverseRecPure, "ReverseRecPure") }

func benchReverseFn(b *testing.B, size int, fn reverseFn) {
	s := epiutil.RandStr(size, "☺ abcdefghijklm nopqrstuvwxyz 世界", rand.NewSource(int64(size)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fn(s)
	}
}

func BenchmarkReverseItr1e1(b *testing.B)     { benchReverseFn(b, 1e1, ReverseItr) }
func BenchmarkReverseRecAux1e1(b *testing.B)  { benchReverseFn(b, 1e1, ReverseRecAux) }
func BenchmarkReverseRecPure1e1(b *testing.B) { benchReverseFn(b, 1e1, ReverseRecPure) }
func BenchmarkReverseItr1e3(b *testing.B)     { benchReverseFn(b, 1e3, ReverseItr) }
func BenchmarkReverseRecAux1e3(b *testing.B)  { benchReverseFn(b, 1e3, ReverseRecAux) }
func BenchmarkReverseRecPure1e3(b *testing.B) { benchReverseFn(b, 1e3, ReverseRecPure) }
func BenchmarkReverseItr1e6(b *testing.B)     { benchReverseFn(b, 1e6, ReverseItr) }
func BenchmarkReverseRecAux1e6(b *testing.B)  { benchReverseFn(b, 1e6, ReverseRecAux) }
func BenchmarkReverseRecPure1e6(b *testing.B) { benchReverseFn(b, 1e6, ReverseRecPure) }
