// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package bstrees

import "testing"

func TestFindFirstK(t *testing.T) {
	tree := &BSTree{108, // A
		&BSTree{108, // B
			&BSTree{-10, // C
				&BSTree{-14, nil, nil}, // D
				&BSTree{2, nil, nil}},  // E
			&BSTree{108, nil, nil}}, // F
		&BSTree{285, // G
			&BSTree{243, nil, nil}, // H
			&BSTree{285, // I
				nil,
				&BSTree{401, nil, nil}}}} // J

	for _, test := range []struct {
		k    interface{}
		want *BSTree
	}{
		{-11, nil},        // _ -> nil
		{108, tree.left},  // A,B,F -> B
		{285, tree.right}, // G,I -> G
	} {
		if got := FindFirstK(tree, test.k.(int)); got != test.want {
			t.Errorf("FindFirstK(%v, %d) = %v; want %v", tree, test.k, got, test.want)
		}
	}
}

func BenchmarkFindFirstK(b *testing.B) {
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
		FindFirstK(tree, 31)
	}
}
