// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package epi

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
		got := Parity(tt.in)
		if got != tt.want {
			t.Errorf("Parity(%d) = %d; want %d", tt.in, got, tt.want)
		}
	}
}

func BenchmarkParity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Parity(uint16(i))
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
		got := ParityLookup(tt.in)
		if got != tt.want {
			t.Errorf("ParityLookup(%d) = %d; want %d", tt.in, got, tt.want)
		}
	}
}

func BenchmarkParityLookup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParityLookup(uint64(i))
	}
}
