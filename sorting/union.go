// Copyright (c) 2016, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package sorting

import "sort"

// EndPoint represents an Interval endpoint.
type EndPoint struct {
	IsClosed bool
	Val      int
}

// Interval defines a set of real numbers.
type Interval struct {
	Left, Right EndPoint
}

type byLeftClosed []Interval

func (it byLeftClosed) Less(i, j int) bool {
	if it[i].Left.Val-it[j].Left.Val != 0 {
		return it[i].Left.Val < it[j].Left.Val
	}
	return it[i].Left.IsClosed && !it[j].Left.IsClosed
}
func (it byLeftClosed) Swap(i, j int) { it[i], it[j] = it[j], it[i] }
func (it byLeftClosed) Len() int      { return len(it) }

// UnionOfIntervals returns union of intervals.
// The time complexity is O(n*log(n)), and O(1) additional space
// is needed, beyond the space needed to write the final result.
func UnionOfIntervals(intervals []Interval) (union []Interval) {
	if len(intervals) == 0 {
		return union
	}

	sort.Sort(byLeftClosed(intervals))
	prev := intervals[0]
	for _, next := range intervals[1:] {
		// intervals overlap || have the same point
		if prev.Right.Val > next.Left.Val || (prev.Right.Val == next.Left.Val && (prev.Left.IsClosed || next.Right.IsClosed)) {
			if prev.Right.Val < next.Right.Val || (prev.Right.Val == next.Right.Val && next.Right.IsClosed) {
				prev.Right = next.Right
			}
		} else {
			union = append(union, prev)
			prev = next
		}
	}
	union = append(union, prev)
	return union
}
