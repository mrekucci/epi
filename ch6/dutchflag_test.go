package ch6

import (
	"reflect"
	"testing"
)

var rearrangeTests = []struct {
	an   [5]int
	i    int
	want [5]int
}{
	{[...]int{5, 4, 3, 2, 1}, 0, [...]int{4, 3, 2, 1, 5}},
	{[...]int{5, 4, 3, 2, 1}, 1, [...]int{1, 3, 2, 4, 5}},
	{[...]int{5, 4, 3, 2, 1}, 2, [...]int{1, 2, 3, 4, 5}},
	{[...]int{5, 4, 3, 2, 1}, 3, [...]int{1, 2, 3, 4, 5}},
	{[...]int{5, 4, 3, 2, 1}, 4, [...]int{1, 3, 2, 4, 5}},
	{[...]int{4, 3, 3, 5, 5}, 0, [...]int{3, 3, 4, 5, 5}},
}

func TestRearrange(t *testing.T) {
	for _, tt := range rearrangeTests {
		got := tt.an
		rearrange(got[:], tt.i)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("rearrange(%d, %d) got %d; want %d", tt.an, tt.i, got, tt.want)
		}
	}
}

func BenchmarkRearrange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rearrange([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}, 7)
	}
}
