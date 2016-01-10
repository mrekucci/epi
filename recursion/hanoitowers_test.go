// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package recursion

import "testing"

func TestHanoiSteps(t *testing.T) {
	for _, test := range []struct {
		in    int
		cnt   int
		steps string
		ok    bool
	}{
		{0, 0, "", true},
		{1, 1, "{p0->p1}", true},
		{2, 3, "{p0->p2}{p0->p1}{p2->p1}", true},
		{3, 7, "{p0->p1}{p0->p2}{p1->p2}{p0->p1}{p2->p0}{p2->p1}{p0->p1}", true},
		{4, 15, "{p0->p2}{p0->p1}{p2->p1}{p0->p2}{p1->p0}{p1->p2}{p0->p2}{p0->p1}{p2->p1}{p2->p0}{p1->p0}{p2->p1}{p0->p2}{p0->p1}{p2->p1}", true},

		// Test ok false.
		{intSize + 1, 0, "", false},
	} {
		cnt, steps, ok := HanoiSteps(test.in)
		if cnt != test.cnt || steps != test.steps || ok != test.ok {
			t.Errorf("HanoiSteps(%d) = %d, %q, %t; want %d, %q, %t", test.in, cnt, steps, ok, test.cnt, test.steps, test.ok)
		}
	}
}

func benchHanoiSteps(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		HanoiSteps(n)
	}
}

func BenchmarkHanoiStepsIntSizeDiv8(b *testing.B) { benchHanoiSteps(b, intSize/8) }
func BenchmarkHanoiStepsIntSizeDiv6(b *testing.B) { benchHanoiSteps(b, intSize/6) }
func BenchmarkHanoiStepsIntSizeDiv4(b *testing.B) { benchHanoiSteps(b, intSize/4) }
