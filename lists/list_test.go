// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package lists

import (
	"fmt"
	"reflect"
	"testing"
)

// Check if list length match.
func checkListLen(l *List, len int) error {
	if l.Len() != len {
		return fmt.Errorf("l.Len() = %d; want %d", l.Len(), len)
	}
	return nil
}

// checkListPointers checks list pointers connections.
func checkListPointers(l *List, ns []*Node) error {
	if err := checkListLen(l, len(ns)); err != nil {
		return err
	}

	// Zero length lists must be the zero value.
	if len(ns) == 0 && l.head != nil {
		return fmt.Errorf("l.head = %p; want nil", l.head)
	}

	// Check pointer connections.
	for i, n := range ns {
		next := (*Node)(nil)
		if i < len(ns)-1 {
			next = ns[i+1]
		}
		if n.next != next {
			return fmt.Errorf("node[%d](%p).next = %p; want %p", i, n, n.next, next)
		}
	}

	return nil
}

// Check list for data equality.
func checkListData(l *List, ns []interface{}) error {
	if err := checkListLen(l, len(ns)); err != nil {
		return err
	}

	i := 0
	for n := l.head; n != nil; n = n.next {
		if n.Data != ns[i] {
			return fmt.Errorf("node[%d].Data = %v; want %v", i, n.Data, ns[i])
		}
		i++
	}

	return nil
}

func TestList(t *testing.T) {
	l := new(List)
	if err := checkListPointers(l, []*Node{}); err != nil {
		t.Errorf("new(List) got %s", err)
	}

	for _, ns := range [][]*Node{
		[]*Node{&Node{Data: 0}},
		[]*Node{&Node{Data: "A"}, &Node{Data: "B"}, &Node{Data: "C"}, &Node{Data: "D"}},
		[]*Node{&Node{Data: 0}, &Node{Data: 1}, &Node{Data: 2}, &Node{Data: 3}, &Node{Data: 4}},
	} {
		// Test insert.
		for _, n := range ns {
			l.Insert(n)
		}
		if err := checkListPointers(l, ns); err != nil {
			t.Errorf("l.Insert all from %v got %s", ns, err)
		}

		// Test to slice.
		s := l.ToSlice()
		if err := checkListData(l, s); err != nil {
			t.Fatalf("l.ToSlice() got %s", err)
		}

		// Test get first node data.
		vf := l.First()
		nf := ns[0]
		if vf != nf.Data {
			t.Errorf("l.First() = %v; want %v", vf, nf.Data)
		}
		if err := checkListPointers(l, ns); err != nil {
			t.Errorf("l.First() from %v got %s", l.ToSlice(), err)
		}

		// Test get last node data.
		vl := l.Last()
		nl := ns[len(ns)-1]
		if vl != nl.Data {
			t.Errorf("l.Last() = %v; want %v", vl, nl.Data)
		}
		if err := checkListPointers(l, ns); err != nil {
			t.Errorf("l.Last() from %v got %s", l.ToSlice(), err)
		}

		// Test remove.
		nl = l.Remove(nl) // Remove last node.
		l.Remove(nl)      // Remove already removed node.
		l.Remove(nil)     // Remove nil.
		if err := checkListPointers(l, ns[:len(ns)-1]); err != nil {
			t.Errorf("l.Remove(%v) from %v got %s", nl.Data, l.ToSlice(), err)
		}
		for n := l.head; n != nil; n = l.head {
			l.Remove(n) // Clean-up list by removing heads.
		}
		if err := checkListPointers(l, []*Node{}); err != nil {
			t.Errorf("l.Remove from %v got %s", l.ToSlice(), err)
		}

		l = new(List) // Set a new list for next test.
	}
}

func TestNewFromSlice(t *testing.T) {
	s := []interface{}{1, 2, 3, 4}
	l := NewFromSlice(s)
	if err := checkListData(l, s); err != nil {
		t.Errorf("NewFromSlice(%v) got %s", s, err)
	}
}

func TestPopInt(t *testing.T) {
	for _, test := range []struct {
		in *Node
		i  int
		n  *Node
		ok bool
	}{
		{nil, 0, nil, true},
		{&Node{Data: 1}, 1, &Node{Data: 1}, true},
		{&Node{Data: "a"}, 0, nil, false},
	} {
		l := new(List)
		l.Insert(test.in)
		if i, n, ok := PopInt(l); i != test.i || !reflect.DeepEqual(n, test.n) || ok != test.ok {
			t.Errorf("PopInt(%v) = %d, %v, %t; want %d, %v, %t", test.in, i, n, ok, test.i, test.n, test.ok)
		}
	}
}
