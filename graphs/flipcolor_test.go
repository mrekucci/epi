// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package graphs

import (
	"math/rand"
	"reflect"
	"testing"
)

type flipColorFn func(a [][]bool, x, y int)

func testFlipColorFn(t *testing.T, fn flipColorFn, fnName string) {
	for _, test := range []struct {
		a    [][]bool
		x, y int
		want [][]bool
	}{
		{ // 1x1 black
			[][]bool{{black}},
			0, 0,
			[][]bool{{white}},
		},
		{ // 1x1 white
			[][]bool{{white}},
			0, 0,
			[][]bool{{black}},
		},
		{ // 2x2
			[][]bool{
				{black, black},
				{black, black}},
			0, 0,
			[][]bool{
				{white, white},
				{white, white}},
		},
		{ // 3x3
			[][]bool{
				{black, black, white},
				{black, white, black},
				{white, white, black}},
			0, 0,
			[][]bool{
				{white, white, white},
				{white, white, black},
				{white, white, black}},
		},
		{ // 10x10
			[][]bool{
				{black, white, black, white, white, white, black, black, black, black},
				{white, white, black, white, white, black, white, white, black, black},
				{black, black, black, white, white, black, black, white, black, black},
				{white, black, white, black, black, black, black, white, black, white},
				{black, white, black, white, white, white, white, black, white, white},
				{black, white, black, white, white, black, white, black, black, black},
				{white, white, white, white, black, white, black, white, white, black},
				{black, white, black, white, black, white, black, white, white, white},
				{black, white, black, black, white, white, white, black, black, black},
				{white, white, white, white, white, white, white, black, black, white}},
			5, 4,
			[][]bool{
				{black, white, black, white, white, white, black, black, black, black},
				{white, white, black, white, white, black, white, white, black, black},
				{black, black, black, white, white, black, black, white, black, black},
				{white, black, white, black, black, black, black, white, black, white},
				{black, black, black, black, black, black, black, black, white, white},
				{black, black, black, black, black, black, black, black, black, black},
				{black, black, black, black, black, black, black, white, white, black},
				{black, black, black, black, black, black, black, white, white, white},
				{black, black, black, black, black, black, black, black, black, black},
				{black, black, black, black, black, black, black, black, black, white}},
		},
	} {
		got := append([][]bool(nil), test.a...) // Make a copy to avoid modification of the original.
		for i := range got {
			got[i] = append([]bool(nil), test.a[i]...)
		}
		if fn(got, test.x, test.y); !reflect.DeepEqual(got, test.want) {
			t.Errorf("%s(%v, %d, %d):", fnName, test.a, test.x, test.y)
			t.Error(" got:", got)
			t.Error("want:", test.want)
		}
	}
}
func TestFlipColorBFS(t *testing.T) { testFlipColorFn(t, FlipColorBFS, "FlipColorBFS") }
func TestFlipColorDFS(t *testing.T) { testFlipColorFn(t, FlipColorDFS, "FlipColorDFS") }

func benchFlipColorFn(b *testing.B, size int, fn flipColorFn) {
	var a [][]bool
	g := rand.New(rand.NewSource(int64(size)))
	for x := 0; x < size; x++ {
		a = append(a, make([]bool, size))
		for y := 0; y < size; y++ {
			a[x][y] = g.Intn(3)%2 == 0
		}
	}
	x, y := g.Intn(size), g.Intn(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FlipColorDFS(a, x, y)
	}
}

func BenchmarkFlipColorBSF1e1(b *testing.B) { benchFlipColorFn(b, 1e1, FlipColorBFS) }
func BenchmarkFlipColorDSF1e1(b *testing.B) { benchFlipColorFn(b, 1e1, FlipColorDFS) }
func BenchmarkFlipColorBSF1e2(b *testing.B) { benchFlipColorFn(b, 1e2, FlipColorBFS) }
func BenchmarkFlipColorDSF1e2(b *testing.B) { benchFlipColorFn(b, 1e2, FlipColorDFS) }
func BenchmarkFlipColorBSF1e3(b *testing.B) { benchFlipColorFn(b, 1e3, FlipColorBFS) }
func BenchmarkFlipColorDSF1e3(b *testing.B) { benchFlipColorFn(b, 1e3, FlipColorDFS) }
