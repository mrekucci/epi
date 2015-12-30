// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package arrays

import (
	"reflect"
	"testing"
)

type genPrimeFn func(int) ([]int, bool)

func testGenPrimesFn(t *testing.T, fn genPrimeFn, fnName string) {
	for _, test := range []struct {
		in   int
		want []int
		ok   bool
	}{
		{-1, nil, true},
		{1, nil, true},
		{2, []int{2}, true},
		{3, []int{2, 3}, true},
		{7, []int{2, 3, 5, 7}, true},
		{15, []int{2, 3, 5, 7, 11, 13}, true},
		{25, []int{2, 3, 5, 7, 11, 13, 17, 19, 23}, true},
		{100, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}, true},
	} {
		if got, ok := fn(test.in); !reflect.DeepEqual(got, test.want) || ok != test.ok {
			t.Errorf("%s(%d) = %v, %t; want %v, %t", fnName, test.in, got, ok, test.want, test.ok)
		}
	}
}

func TestGenPrimesTrialDiv(t *testing.T) { testGenPrimesFn(t, GenPrimesTrialDiv, "GenPrimesTrialDiv") }
func TestGenPrimesSieve(t *testing.T)    { testGenPrimesFn(t, GenPrimesSieve, "GenPrimesSieve") }

func benchGenPrimesFn(b *testing.B, size int, fn genPrimeFn) {
	for i := 0; i < b.N; i++ {
		fn(size)
	}
}

func BenchmarkGenPrimesTrialDiv1e2(b *testing.B) { benchGenPrimesFn(b, 1e2, GenPrimesTrialDiv) }
func BenchmarkGenPrimesSieve1e2(b *testing.B)    { benchGenPrimesFn(b, 1e2, GenPrimesSieve) }
func BenchmarkGenPrimesTrialDiv1e4(b *testing.B) { benchGenPrimesFn(b, 1e4, GenPrimesTrialDiv) }
func BenchmarkGenPrimesSieve1e4(b *testing.B)    { benchGenPrimesFn(b, 1e4, GenPrimesSieve) }
func BenchmarkGenPrimesTrialDiv1e6(b *testing.B) { benchGenPrimesFn(b, 1e6, GenPrimesTrialDiv) }
func BenchmarkGenPrimesSieve1e6(b *testing.B)    { benchGenPrimesFn(b, 1e6, GenPrimesSieve) }
