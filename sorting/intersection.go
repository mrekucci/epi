// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package sorting

// IntersectSorted returns intersection of x and y without duplicated elements.
// The time complexity is O(m+n). The O(1) additional space is needed.
func IntersectSorted(x, y []int) (ixy []int) {
	i, j := 0, 0
	for i < len(x) && j < len(y) {
		switch ex, ey := x[i], y[j]; {
		case ex == ey && (i == 0 || x[i-1] != ex):
			ixy = append(ixy, ex)
			i++
			j++
		case ex <= ey:
			i++
		case ex > ey:
			j++
		}
	}
	return ixy
}
