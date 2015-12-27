// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package heaps

import "container/heap"

// entry holds value v of slice on index i in slice of slices.
type entry struct {
	v int
	i int
}

// minHeap is an implementation of min heap for entry values.
type minHeap []*entry

func (h *minHeap) Less(i, j int) bool { return (*h)[i].v < (*h)[j].v }
func (h *minHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *minHeap) Len() int           { return len(*h) }

func (h *minHeap) Pop() (v interface{}) {
	v = (*h)[h.Len()-1]
	(*h)[h.Len()-1] = nil
	*h = (*h)[:h.Len()-1]
	return v
}

func (h *minHeap) Push(v interface{}) { *h = append(*h, v.(*entry)) }

// MergeSorted merges given slice of already
// sorted slices into one sorted slice that slice.
// The time complexity is O(n*log(k)) where n is the total number
// of elements combining all slices together and k = len(ss).
// The O(k) additional space is needed (beyond the space needed
// to write the final result).
func MergeSorted(ss [][]int) (m []int) {
	h := new(minHeap)
	idx := make([]int, len(ss)) // Stores the next index of element from particular slice.
	for i, s := range ss {      // Add first entry from each slice to the heap.
		if len(s) > 0 {
			heap.Push(h, &entry{s[0], i})
			idx[i] = 1
		}
	}
	for h.Len() != 0 {
		e := heap.Pop(h).(*entry)
		m = append(m, e.v)
		if s, i := ss[e.i], idx[e.i]; i < len(s) {
			heap.Push(h, &entry{s[i], e.i}) // Add next entry from the slice that previously had the smallest element.
			idx[e.i]++
		}
	}
	return m
}
