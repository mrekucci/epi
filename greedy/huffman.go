// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package greedy

import "container/heap"

// Symbol represents a character with its frequency and optimal binary code.
type Symbol struct {
	c    rune
	prob float64
	code string
}

type bTree struct {
	prob        float64
	s           *Symbol
	left, right *bTree
}

// minHeap is an implementation of min heap for Symbol values.
type minHeap []*bTree

func (h *minHeap) Less(i, j int) bool { return (*h)[i].prob < (*h)[j].prob }
func (h *minHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *minHeap) Len() int           { return len(*h) }

func (h *minHeap) Pop() (v interface{}) {
	v = (*h)[h.Len()-1]
	(*h)[h.Len()-1] = nil
	*h = (*h)[:h.Len()-1]
	return v
}

func (h *minHeap) Push(v interface{}) { *h = append(*h, v.(*bTree)) }

// assignCode assigns huffman code to symbol base on its position in the tree t.
func assignCode(t *bTree, code []byte) {
	if t.s != nil { // Base case. Assign code only for children.
		t.s.code = string(code)
	} else {
		assignCode(t.left, append(code, '0'))
		assignCode(t.right, append(code, '1'))
	}
}

// HuffmanEncoding assigns huffman code to given symbols based on its given frequency.
// The time complexity is O(n*log(n))), and O(n) additional space is needed.
func HuffmanEncoding(symbols []*Symbol) {
	if len(symbols) == 0 {
		return
	}

	h := new(minHeap)
	for _, s := range symbols { // Initialize heap with disconnected nodes of symbols, which all represents children.
		heap.Push(h, &bTree{s.prob, s, nil, nil})
	}

	for h.Len() > 1 { // Create parents for the children and build a binary tree until one (root) node left.
		l, r := heap.Pop(h).(*bTree), heap.Pop(h).(*bTree)
		heap.Push(h, &bTree{l.prob + r.prob, nil, l, r})
	}

	assignCode(heap.Pop(h).(*bTree), []byte(nil)) // Traverse the binary tree and assign the codes.
}
