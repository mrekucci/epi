// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package heaps

import "container/heap"

// entry holds value v of slice on index i in slice of slices.
type entry struct { i, v int }

// minEntryHeap is an implementation of min heap for entry values.
type minEntryHeap []*entry

func (h *minEntryHeap) Less(i, j int) bool { return (*h)[i].v < (*h)[j].v }
func (h *minEntryHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *minEntryHeap) Len() int           { return len(*h) }

func (h *minEntryHeap) Pop() (v interface{}) {
	v = (*h)[h.Len()-1]
	(*h)[h.Len()-1] = nil
	*h = (*h)[:h.Len()-1]
	return v
}

func (h *minEntryHeap) Push(v interface{}) { *h = append(*h, v.(*entry)) }

// MergeSorted merges given slice of already sorted slices into one sorted slice.
// The time complexity is O(n*log(k)) where n is the total number
// of elements combining all slices together and k = len(ss).
// The O(k) additional space is needed (beyond the space needed
// to write the final result).
func MergeSorted(ss [][]int) (m []int) {
	h := new(minEntryHeap)
	next := make([]int, len(ss)) // Stores the next index of element from particular slice.
	for i, s := range ss {      // Add first entry from each slice to the heap.
		if len(s) > 0 {
			heap.Push(h, &entry{i, s[0]})
			next[i] = 1
		}
	}
	for h.Len() != 0 {
		e := heap.Pop(h).(*entry)
		m = append(m, e.v)
		if s, i := ss[e.i], next[e.i]; i < len(s) {
			heap.Push(h, &entry{e.i, s[i]}) // Add next entry from the slice that previously had the smallest element.
			next[e.i]++
		}
	}
	return m
}
