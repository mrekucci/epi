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

// SelectionSort sorts given data and has next properties:
//
// - Not stable
// - O(1) extra space
// - Θ(n*n) comparisons
// - Θ(n) swaps
// - Not adaptive
//
func SelectionSort(data sort.Interface) {
	for i := 0; i < data.Len()-1; i++ { // Invariant: data[0:i] holds the number of i smallest elements in sorted order from data[:].
		k := i
		for j := i + 1; j < data.Len(); j++ { // Invariant: data[k] is the smallest element in data[i:j].
			if data.Less(j, k) {
				k = j
			}
		}
		data.Swap(i, k)
	}
}

// InsertionSort sorts given data and has next properties:
//
// - Stable
// - O(1) extra space
// - O(n*n) comparisons and swaps
// - Adaptive: O(n) time when nearly sorted
// - Very low overhead
//
func InsertionSort(data sort.Interface) {
	for i := 1; i < data.Len(); i++ { // Invariant: at the start of each iteration, the data[0:i] consist of the elements originally in data[0:i], but in sorted order.
		for k := i; k > 0 && data.Less(k, k-1); k-- { // Invariant: data[k:i] will have the smallest element on position k.
			data.Swap(k, k-1)
		}
	}
}
