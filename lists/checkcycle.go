// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package lists

// HasCycle returns *Node that represents the start of
// the cycle in l or nil when there is no cycle in l.
func HasCycle(l *List) *Node {
	ns := make(map[*Node]bool)
	for n := l.head; n != nil; n = n.next {
		if ns[n] {
			return n
		}
		ns[n] = true
	}
	return nil
}
