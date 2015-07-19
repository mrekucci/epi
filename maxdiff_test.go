// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package epi

import (
	"math/rand"
	"testing"
)

var minBatteryCapTests = []struct {
	in   []int
	want int
	ok   bool
}{
	{[]int{}, 0, true},
	{[]int{0}, 0, true},
	{[]int{1}, 0, true},
	{[]int{-1}, 0, true},
	{[]int{0, 0}, 0, true},
	{[]int{0, 1}, 1, true},
	{[]int{1, 0}, 0, true},
	{[]int{3, 3, 10}, 7, true},
	{[]int{-3, 3, 10}, 13, true},
	{[]int{3, -3, 10}, 13, true},
	{[]int{3, 3, -10}, 0, true},
	{[]int{1, maxInt}, maxInt - 1, true},
	{[]int{-1, maxInt}, 0, false},
	{[]int{1, -maxInt}, 0, true},
	{[]int{-1, -maxInt}, 0, true},
	{[]int{1, minInt}, 0, false},
	{[]int{-1, minInt}, 0, true},
	{[]int{minInt, maxInt}, 0, false},
	{[]int{maxInt, minInt}, 0, false},
}

func TestMinBatteryCap(t *testing.T) {
	for _, tt := range minBatteryCapTests {
		got, ok := MinBatteryCap(tt.in)
		if ok != tt.ok || got != tt.want {
			t.Errorf("MinBatteryCap(%d) = %d, %t; want %d, %t", tt.in, got, ok, tt.want, tt.ok)
		}
	}
}

func benchMinBatteryCap(b *testing.B, size int) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := rand.New(rand.NewSource(1235)).Perm(size)
		b.StartTimer()
		MinBatteryCap(data)
		b.StopTimer()
	}
}

func BenchmarkMinBatteryCap1e2(b *testing.B) { benchMinBatteryCap(b, 1e2) }
func BenchmarkMinBatteryCap1e4(b *testing.B) { benchMinBatteryCap(b, 1e4) }
func BenchmarkMinBatteryCap1e6(b *testing.B) { benchMinBatteryCap(b, 1e6) }
