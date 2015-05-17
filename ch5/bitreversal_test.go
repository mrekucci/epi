package ch5

import "testing"

var reverseBitsTests = []struct {
	in   uint16
	want uint16
}{
	{1, 1 << 15},
	{2, 1 << 14},
	{4, 1 << 13},
	{8, 1 << 12},
	{0x000f, 0xf000},
	{0x00f0, 0x0f00},
	{0x00ff, 0xff00},
	{0xffff, 0xffff},
}

func TestReverseBits(t *testing.T) {
	for _, tt := range reverseBitsTests {
		got := ReverseBits(tt.in)
		if got != tt.want {
			t.Errorf("ReverseBits(%#x) = %#x; want %#x", tt.in, got, tt.want)
		}
	}
}

func BenchmarkReverseBits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReverseBits(uint16(i))
	}
}

var reverseBitsLookupTests = []struct {
	in   uint64
	want uint64
}{
	{1, 1 << 63},
	{2, 1 << 62},
	{4, 1 << 61},
	{8, 1 << 60},
	{0x00000000ffffffff, 0xffffffff00000000},
	{0x00000000000000ff, 0xff00000000000000},
	{0x00000000ff000000, 0x000000ff00000000},
	{0xffffffffffffffff, 0xffffffffffffffff},
}

func TestReverseBitsLookup(t *testing.T) {
	for _, tt := range reverseBitsLookupTests {
		got := ReverseBitsLookup(tt.in)
		if got != tt.want {
			t.Errorf("ReverseBitsLookup(%#x) = % #x; want %#x", tt.in, got, tt.want)
		}
	}
}

func BenchmarkReverseBitsLookup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReverseBitsLookup(uint64(i))
	}
}
