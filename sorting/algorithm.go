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
	// Loop invariant: data[0:i-1] consist of sorted elements
	// that are smaller (or equal) then data[i:len(data)].
	for i := 1; !isSorted && i < data.Len(); i++ {
		isSorted = true
		// Loop invariant: data[j:len(data)] will
		// have the smallest element on position j.
		for j := data.Len() - 1; j >= i; j-- {
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
	// Loop invariant: data[0:i] holds the number of i
	// smallest elements in sorted order from data[:].
	for i := 0; i < data.Len()-1; i++ {
		k := i
		// Loop invariant: data[k] is the smallest element in data[i:j].
		for j := i + 1; j < data.Len(); j++ {
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
	// Loop invariant: at the start of each iteration, the data[0:i] consist
	// of the elements originally in data[0:i], but in sorted order.
	for i := 1; i < data.Len(); i++ {
		// Loop invariant: data[k:i] will have the smallest element on position k.
		for k := i; k > 0 && data.Less(k, k-1); k-- {
			data.Swap(k, k-1)
		}
	}
}

// HeapSort sorts given data and has next properties:
//
// - Not stable
// - O(1) extra space
// - O(n*lg(n)) time
// - Not really adaptive
//
func HeapSort(data sort.Interface) {
	// down restores order of heap.
	down := func(root, n int) {
		for {
			lch := 2*root + 1
			if lch >= n || lch < 0 { // child < 0 when int overflow.
				break
			}
			if rch := lch + 1; rch < n && data.Less(lch, rch) { // lch+1 == 2*root + 2 // Right child.
				lch = rch
			}
			if !data.Less(root, lch) { // Heap is ordered.
				return
			}
			data.Swap(root, lch)
			root = lch
		}
	}

	// Heapify (build a max heap).
	for i := (data.Len() - 1) / 2; i >= 0; i-- {
		down(i, data.Len())
	}

	// Pop elements, largest first, into end of data.
	// Loop invariant: data[i:] contains the data.Len()-1-i largest elements
	// of maxHeap and the maxHeap contains i+1 smallest elements.
	for i := data.Len() - 1; i >= 0; i-- {
		data.Swap(0, i)
		down(0, i)
	}
}

// MergeSort (in-place version) sorts given data and has next properties:
//
// - Not stable
// - O(1) extra space
// - O(n*n*lg(n)) time worst case (when array is: reverse sorted/few unique/random)
// - O(n*lg(n)) time best case (when array is: sorted/nearly sorted)
// - Adaptive
//
// Note: For regular Merge Sort (not in-place version) the worst
// case is O(n*lg(n)) and O(n) additional space is needed.
//
func MergeSort(data sort.Interface) {
	// move moves the largest element in data[s:e+1] into the position e.
	// It has O(n) time complexity (n=e-s) and O(1) additional space is needed.
	move := func(s, e int) {
		for i := s; i < e; i++ {
			if data.Less(i+1, i) {
				data.Swap(i+1, i)
			}
		}
	}

	// merge combines sorted data[p:q+1] and data[q+1:r+1] together.
	// The time complexity is O((q-p)*(r-q+1)). No additional space is needed.
	merge := func(p, q, r int) {
		// Loop invariant: data[0:i], data[i:q+1] and data[q+1:r+1] are sorted, and every
		// elements in data[0:i] are less or equal to the data[q+1].
		for i := p; i <= q; i++ {
			if data.Less(q+1, i) {
				data.Swap(q+1, i)
				move(q+1, r)
			}
		}
	}

	// mergeSortFn is a recursive function that sorts data[l:r+1].
	var mergeSortFn func(l, r int)
	mergeSortFn = func(l, r int) {
		switch cnt := r - l; {
		case cnt == 1: // Two elements left.
			if data.Less(r, l) {
				data.Swap(l, r)
			}
		case cnt > 1:
			q := l + (r-l)/2
			mergeSortFn(l, q)   // Divide.
			mergeSortFn(q+1, r) // Divide.
			merge(l, q, r)      // Combine and Conquer.
		}
	}

	mergeSortFn(0, data.Len()-1)
}

// QuickSort sorts given data and has next properties:
//
// - Not stable
// - O(n) extra space in worst case; O(lg(n)) typically
// - O(n*n) time when few unique keys or is sorted, but typically O(n*lg(n)) time
// - Not adaptive
//
func QuickSort(data sort.Interface) {
	// median moves the median of data[a], data[b], data[c] into data[b].
	// Selecting a median helps avoid to cause worst-case behavior on already
	// sorted (or reverse-sorted) data, and gives a better estimate of the optimal
	// pivot than selecting any single element, when no information about the
	// ordering of the input is known.
	median := func(a, b, c int) {
		if data.Less(b, a) {
			data.Swap(b, a)
		}
		if data.Less(c, b) {
			data.Swap(c, b)
			if data.Less(b, a) {
				data.Swap(b, a)
			}
		}
	}

	// quickSortFn is a recursive function that sorts data[l:r+1].
	var quickSortFn func(l, r int)
	quickSortFn = func(l, r int) {
		if l < r {
			// Partition (divide). The time complexity is O(n). The O(1) additional space is needed.
			q := l
			pivot := l + (r-l)/2
			median(l, pivot, r) // Is much faster then rand.Intn(r-p+1)+p.
			data.Swap(pivot, r)       // Move the pivot value into the position r.
			// Loop invariant:
			// each element in data[p:q] is less than or equal to the pivot;
			// each element in data[q:u] is greater than the pivot;
			// each element in data[u:r] is unsorted;
			// the element data[r] holds the pivot.
			for u := q; u < r; u++ {
				if data.Less(u, r) {
					data.Swap(u, q)
					q++
				}
			}
			data.Swap(r, q)

			quickSortFn(l, q-1) // Conquer.
			quickSortFn(q+1, r) // Conquer.
			// Because the sub-arrays are already sorted, no work is needed to combine them.
		}
	}

	quickSortFn(0, data.Len()-1)
}
