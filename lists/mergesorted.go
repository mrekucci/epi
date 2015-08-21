// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package lists

// MergeSorted merges int nodes from l and f sorted lists into the ordered list m.
// Note: when l or f contains different type from int then false is returned and
// merged list will contains some value(s) merged from l or f up to the different
// type.
func MergeSorted(l, f *List) (m *List, ok bool) {
	m = new(List)
	for l.Len() > 0 || f.Len() > 0 {
		vl, nl, okl := PopInt(l)
		vf, nf, okf := PopInt(f)
		if !okl || !okf {
			return m, false
		}

		ll, n := l, nl // The assumption is: vl <= vf.
		switch {
		case l.Len() == 0:
			ll, n = f, nf
		case f.Len() == 0:
			ll, n = l, nl
		case vl > vf:
			ll, n = f, nf
		}

		m.Insert(ll.Remove(n))
	}
	return m, true
}
