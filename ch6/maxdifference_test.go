package ch6

import "testing"

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

func BenchmarkMinBatteryCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinBatteryCap([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}) // TODO: improve benchmark.
	}
}
