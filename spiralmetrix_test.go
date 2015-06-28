// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package epi

import (
	"math/rand"
	"reflect"
	"testing"
)

var clockwiseTests = []struct {
	in   [][]int
	want []int
}{
	{ // 1x1
		[][]int{{1}},
		[]int{1}},
	{ // 2x2
		[][]int{
			{1, 2},
			{3, 4}},
		[]int{1, 2, 4, 3}},
	{ // 3x3
		[][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9}},
		[]int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
	{ // 4x4
		[][]int{
			{0, 1, 2, 3},
			{4, 5, 6, 7},
			{8, 9, 10, 11},
			{12, 13, 14, 15}},
		[]int{0, 1, 2, 3, 7, 11, 15, 14, 13, 12, 8, 4, 5, 6, 10, 9}},
}

func TestClockwise(t *testing.T) {
	for _, tt := range clockwiseTests {
		got := Clockwise(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Clockwise(%d)", tt.in)
			t.Errorf(" got %d", got)
			t.Errorf("want %d", tt.want)
		}
	}
}

func benchClockwise(b *testing.B, size int) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		ints := rand.Perm(size)
		data := make([][]int, len(ints))
		for j := range data {
			data[j] = append([]int(nil), ints...)
		}
		b.StartTimer()
		Clockwise(data)
		b.StopTimer()
	}
}

func BenchmarkClockwise1e1(b *testing.B) { benchClockwise(b, 1e1) }
func BenchmarkClockwise1e2(b *testing.B) { benchClockwise(b, 1e2) }
func BenchmarkClockwise1e3(b *testing.B) { benchClockwise(b, 1e3) }
