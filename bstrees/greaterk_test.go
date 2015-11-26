// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package bstrees

import "testing"

func TestFindFirstGreaterK(t *testing.T) {
	tree := &BSTree{19, // A
		&BSTree{7, // B
			&BSTree{3, // C
				&BSTree{2, nil, nil},  // D
				&BSTree{5, nil, nil}}, // E
			&BSTree{11, // F
				nil,
				&BSTree{17, // G
					&BSTree{13, nil, nil}, // H
					nil}}},
		&BSTree{43, // I
			&BSTree{23, // J
				nil,
				&BSTree{37, // K
					&BSTree{29, // L
						nil,
						&BSTree{31, nil, nil}}, // M
					&BSTree{41, nil, nil}}}, // N
			&BSTree{47, // O
				nil,
				&BSTree{53, nil, nil}}}} // P

	for _, test := range []struct {
		k    interface{}
		want *BSTree
	}{
		{-1, tree.left.left.left},        // _ -> D
		{3, tree.left.left.right},        // C -> E
		{23, tree.right.left.right.left}, // J -> L
		{31, tree.right.left.right},      // M -> K
		{47, tree.right.right.right},     // O -> P
		{53, nil},                        // P -> nil
	} {
		if got := FindFirstGreaterK(tree, test.k.(int)); got != test.want {
			t.Errorf("FindFirstGreaterK(%v, %d) = %v; want %v", tree, test.k, got, test.want)
		}
	}
}

func BenchmarkFindFirstGreaterK(b *testing.B) {
	tree := &BSTree{19,
		&BSTree{7,
			&BSTree{3,
				&BSTree{2, nil, nil},
				&BSTree{5, nil, nil}},
			&BSTree{11,
				nil,
				&BSTree{17,
					&BSTree{13, nil, nil},
					nil}}},
		&BSTree{43,
			&BSTree{23,
				nil,
				&BSTree{37,
					&BSTree{29,
						nil,
						&BSTree{31, nil, nil}},
					&BSTree{41, nil, nil}}},
			&BSTree{47,
				nil,
				&BSTree{53, nil, nil}}}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindFirstGreaterK(tree, 31)
	}
}
