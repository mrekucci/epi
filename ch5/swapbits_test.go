package ch5

import "testing"

var swapBitsTests = []struct {
	x    uint64
	i, j uint64
	want uint64
}{
	{1, 0, 1, 1 << 1},
	{5, 2, 4, 17},
	{20, 2, 4, 20},
	{1, 0, 63, 1 << 63},
}

func TestSwapBits(t *testing.T) {
	for _, tt := range swapBitsTests {
		got := swapBits(tt.x, tt.i, tt.j)
		if got != tt.want {
			t.Errorf("swapBits(%d, %d, %d) = %d; want %d", tt.x, tt.i, tt.j, got, tt.want)
		}
	}
}

func BenchmarkSwapBits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		swapBits(uint64(i), 0, 63)
	}
}
