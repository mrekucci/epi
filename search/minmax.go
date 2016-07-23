// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package search

// MinMax returns the min and max values from xs.
// The time complexity is O(n). The O(1) additional space is needed.
func MinMax(xs []int) (min, max int) {
	if len(xs) == 0 {
		return 0, 0
	}

	min, max = xs[0], xs[0]
	for _, v := range xs[1:] {
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}
	return min, max
}

// compare returns a, b if a < b, otherwise b, a.
func compare(a, b int) (min, max int) {
	if a < b {
		return a, b
	}
	return b, a
}

// MinMaxAlt is an alternative function to MinMax for searching min and max in xs.
// The time complexity is O(n). The O(1) additional space is needed.
// The time complexity hidden constant is more significant here than in the MaxMin.
func MinMaxAlt(xs []int) (min, max int) {
	switch len(xs) {
	case 0:
		return 0, 0
	case 1:
		return xs[0], xs[0]
	}

	min, max = compare(xs[0], xs[1])
	for i := 3; i < len(xs); i += 2 { // Compare 3*n/2 - 2 times.
		nMin, nMax := compare(xs[i-1], xs[i])
		min, _ = compare(nMin, min)
		_, max = compare(nMax, max)
	}

	if len(xs)%2 != 0 { // We need to compare the last element if length of xs is odd.
		v := xs[len(xs)-1]
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}

	return min, max
}
