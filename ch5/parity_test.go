package ch5

import (
	"math"
	"testing"
)

const (
	odd  = 1
	even = 0
)

var parityTests = []struct {
	in   uint16
	want uint16
}{
	{0, even},
	{1, odd},
	{2, odd},
	{3, even},
	{4, odd},
	{5, even},
	{6, even},
	{7, odd},
	{8, odd},
	{9, even},
	{math.MaxUint16, even},
}

func TestParity(t *testing.T) {
	for _, tt := range parityTests {
		got := parity(tt.in)
		if got != tt.want {
			t.Errorf("parity(%d) = %d; want %d", tt.in, got, tt.want)
		}
	}
}

func BenchmarkParity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parity(uint16(i))
	}
}

var parityLookupTests = []struct {
	in   uint64
	want uint16
}{
	{0, even},
	{1, odd},
	{2, odd},
	{3, even},
	{4, odd},
	{5, even},
	{6, even},
	{7, odd},
	{8, odd},
	{9, even},
	{math.MaxUint16, even},
}

func TestParityLookup(t *testing.T) {
	for _, tt := range parityLookupTests {
		got := parityLookup(tt.in)
		if got != tt.want {
			t.Errorf("parityLookup(%d) = %d; want %d", tt.in, got, tt.want)
		}
	}
}

func BenchmarkParityLookup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parityLookup(uint64(i))
	}
}
