// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package sorting

import "sort"

// BubbleSort sorts given data and has next properties:
//
// - Stable
// - O(1) extra space needed
// - O(n*n) comparisons and swaps
// - Adaptive: O(n) when nearly sorted
//
func BubbleSort(data sort.Interface) {
	var isSorted bool
	for i := 1; !isSorted && i < data.Len(); i++ { // Invariant: data[0:i-1] consist of sorted elements that are smaller (or equal) then data[i:len(data)].
		isSorted = true
		for j := data.Len() - 1; j >= i; j-- { // Invariant: data[j:len(data)] will have the smallest element on position j.
			if data.Less(j, j-1) {
				data.Swap(j, j-1)
				isSorted = false
			}
		}
	}
}
