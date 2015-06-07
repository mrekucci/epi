package ch6

import (
	"reflect"
	"testing"
)

var nextTests = []struct {
	in   []int
	want []int
}{
	{[]int{1}, []int{}},
	{[]int{1, 2}, []int{2, 1}},
	{[]int{2, 1}, []int{}},
	{[]int{1, 2, 3}, []int{1, 3, 2}},
	{[]int{1, 3, 2}, []int{2, 1, 3}},
	{[]int{2, 1, 3}, []int{2, 3, 1}},
	{[]int{2, 3, 1}, []int{3, 1, 2}},
	{[]int{3, 1, 2}, []int{3, 2, 1}},
	{[]int{3, 2, 1}, []int{}},
}

func TestNext(t *testing.T) {
	for _, tt := range nextTests {
		got := Next(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Next(%d) = %d; want %d", tt.in, got, tt.want)
		}
	}
}

func BenchmarkNext(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Next([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})
	}
}
