// Copyright (c) 2016, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package heaps

import (
	"container/heap"
	"math/big"
)

type minIntHeap []int

func (h *minIntHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *minIntHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *minIntHeap) Len() int           { return len((*h)) }

func (h *minIntHeap) Pop() (e interface{}) {
	e, *h = (*h)[h.Len()-1], (*h)[:h.Len()-1]
	return e
}

func (h *minIntHeap) Push(e interface{}) { *h = append(*h, e.(int)) }

type maxIntHeap []int

func (h *maxIntHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *maxIntHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *maxIntHeap) Len() int           { return len((*h)) }

func (h *maxIntHeap) Pop() (e interface{}) {
	e, *h = (*h)[h.Len()-1], (*h)[:h.Len()-1]
	return e
}

func (h *maxIntHeap) Push(e interface{}) { *h = append(*h, e.(int)) }

// MedianStream reads the stream of values from in channel, computes a median from
// first n read values, and write the result after every read entry to out channel.
// The time complexity per read entry is O(log(n)).
// The O(n) additional space is needed for the first n read entries.
func MedianStream(in <-chan int, out chan<- *big.Rat) {
	min, max := new(minIntHeap), new(maxIntHeap)
	for v := range in {
		if min.Len() == 0 || v >= (*min)[0] { // min.Len() == 0: to insert the very first element.
			heap.Push(min, v)
		} else {
			heap.Push(max, v)
		}

		// Balance the two heaps:
		// - if even number of elements was read from in, then
		// the min and max heap must have equal number of elements
		// - if odd number of elements was read from in, then
		// the the min heap must have one more element then the max heap
		if min.Len() > max.Len()+1 {
			heap.Push(max, heap.Pop(min))
		} else if max.Len() > min.Len() {
			heap.Push(min, heap.Pop(max))
		}

		if min.Len() == max.Len() { // Even number of elements was read.
			out <- big.NewRat(int64((*min)[0])+int64((*max)[0]), 2)
		} else { // Odd number of elements was read.
			out <- big.NewRat(int64((*min)[0]), 1)
		}
	}
}
