// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package graphs

import (
	"math/rand"
	"testing"
)

func TestFindCelebrity(t *testing.T) {
	for _, test := range []struct {
		in   [][]bool
		want int
	}{
		// There is no celebrity on the party.
		{[][]bool{
			//  A     B
			{false, true}, // A
			{true, false}, // B
		}, -1},
		// A doesn't know the celebrity B.
		{[][]bool{
			//  A     B     C
			{false, false, true},  // A
			{false, false, false}, // B
			{true, true, false},   // C
		}, -1},
		{[][]bool{
			//  A     B     C
			{false, true, true},   // A
			{false, false, false}, // B
			{true, true, false},   // C
		}, 1},
		{[][]bool{
			//  A     B     C     D     E
			{false, true, true, true, true},     // A
			{true, false, false, true, true},    // B
			{true, true, false, true, true},     // C
			{false, false, false, false, false}, // D
			{true, false, false, true, false},   // E
		}, 3},
	} {
		if got := FindCelebrity(test.in); got != test.want {
			t.Errorf("FindCelebrity(%v) = %d; want %d", test.in, got, test.want)
		}
	}
}

func benchFindCelebrity(b *testing.B, size int) {
	g := rand.New(rand.NewSource(int64(size)))
	var f [][]bool
	for r := 0; r < size; r++ {
		f = append(f, make([]bool, size))
		for c := 0; c < size; c++ {
			f[r][c] = g.Int()%size != 0
		}
		f[r][r] = false
	}
	celebrity := size - 1
	for i := 0; i < size; i++ {
		f[i][celebrity] = true
	}
	for i := 0; i < size; i++ {
		f[celebrity][i] = false
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindCelebrity(f)
	}
}

func BenchmarkFindCelebrity1e1(b *testing.B) { benchFindCelebrity(b, 1e1) }
func BenchmarkFindCelebrity1e2(b *testing.B) { benchFindCelebrity(b, 1e2) }
func BenchmarkFindCelebrity1e3(b *testing.B) { benchFindCelebrity(b, 1e3) }
