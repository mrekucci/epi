// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package lists

// EvenOddMerge rearrange l in way that nodes with even
// number appear first followed by nodes with odd numbers.
// A false is returned when l doesn't start with 0.
// The time complexity is O(n), and O(1) additional space is needed.
func EvenOddMerge(l *List) (ok bool) {
	switch {
	case l.Len() < 2: // l.head == nil || l.head.next == nil
		return true
	case l.head.Data.(int) != 0:
		return false
	}

	var even, odd, curr *Node = l.head, nil, l.head.next
	for curr != nil {
		if curr.Data.(int)%2 == 0 {
			if odd != nil {
				odd.next = curr.next
				curr.next = even.next
				even.next = curr
				even = curr
				curr = odd
			} else {
				even = curr // Track last even.
			}
		} else {
			odd = curr // Track last odd.
		}
		curr = curr.next
	}
	return true
}
