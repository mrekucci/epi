// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package lists

import (
	"errors"
	"math/big"
)

// ErrSort indicates that a list isn't sorted.
var ErrSort = errors.New("MedianOfSortedCircular: list isn't sorted")

// ErrNode indicates that a list doesn't contain the node.
var ErrNode = errors.New("MedianOfSortedCircular: list doesn't contain node")

// ErrNode indicates that a list doesn't contain the node.
var ErrCycle = errors.New("MedianOfSortedCircular: list doesn't contain cycle")

// ErrType indicates that a value has different type than expected.
var ErrType = errors.New("MedianOfSortedCircular: value is not of the type int")

// MedianOfSortedCircular returns a median of sorted circular linked list of
// integer values. An error is returned when the list isn't sorted, isn't
// circular, doesn't contain the node, or contains a value other than int type.
func MedianOfSortedCircular(l *List, csn *Node) (*big.Rat, error) {
	if l.head == nil {
		return nil, nil
	}

	// Find the list length.
	len := 0
	cnt := 0
	for p, n := l.head, l.head; n != nil; p, n = n, n.next {
		if n == csn {
			if cnt++; cnt == 2 { // Cycle detected.
				break
			}
		}
		len++
		vp, okp := p.Data.(int)
		vn, okn := n.Data.(int)
		if !okp || !okn { // Check whether the type assertion is ok.
			return nil, ErrType
		}
		if vp > vn { // Check whether data are sorted.
			return nil, ErrSort
		}
	}
	switch cnt {
	case 0:
		return nil, ErrNode
	case 1:
		return nil, ErrCycle
	}

	// Find the middle node(s).
	p, n := l.head, l.head
	for i := 0; n != nil && i < len/2; i++ {
		p, n = n, n.next
	}

	// Evaluate the median (we've already checked the type assertion for all values).
	if len%2 == 0 {
		return big.NewRat(int64(n.Data.(int)+p.Data.(int)), 2), nil
	}
	return big.NewRat(int64(n.Data.(int)), 1), nil
}
