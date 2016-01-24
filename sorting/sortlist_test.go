// Copyright (c) 2016, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package sorting

import (
	"sort"
	"testing"
)

// newFromSlice returns linked list filled with given data.
func newFromSlice(data []int) (list *Node) {
	if len(data) == 0 {
		return nil
	}

	list = new(Node)
	l := list
	l.Data = data[0]
	for _, v := range data[1:] {
		l.next = &Node{Data: v}
		l = l.next
	}
	return list
}

// toSlice returns data of this linked list as a slice.
func toSlice(list *Node) (data []int) {
	for n := list; n != nil; n = n.next {
		data = append(data, n.Data)
	}
	return data
}

func TestSortList(t *testing.T) {
	for _, data := range [][]int{
		nil,
		[]int{0},
		[]int{1, 0},
		[]int{1, 3, 2, 5, 4, 6, 9, 8, 7, 0},
		[]int{1, -3, 2, -5, 4, -6, 9, -8, 7, 0},
	} {
		l := newFromSlice(data)
		if got := toSlice(SortList(l)); !sort.IsSorted(sort.IntSlice(got)) {
			t.Errorf("SortList(%v)", data)
			t.Errorf(" got %v", got)
			sort.Ints(data)
			t.Errorf("want %v", data)
		}
	}
}

// TODO: benchmark!
