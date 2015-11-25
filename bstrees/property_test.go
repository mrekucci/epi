// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package bstrees

import "testing"

func TestIsBinaryTreeBST(t *testing.T) {
	for _, test := range []struct {
		tree *BSTree
		want bool
	}{
		{&BSTree{minInt, &BSTree{maxInt, nil, nil}, nil}, false},
		{&BSTree{minInt, nil, &BSTree{maxInt, nil, nil}}, true},
		{&BSTree{maxInt, &BSTree{minInt, nil, nil}, nil}, true},
		{&BSTree{maxInt, nil, &BSTree{minInt, nil, nil}}, false},
		{&BSTree{0, &BSTree{minInt, nil, nil}, &BSTree{maxInt, nil, nil}}, true},
		{&BSTree{0, &BSTree{1, nil, nil}, &BSTree{2, nil, nil}}, false},
		{&BSTree{0, &BSTree{2, nil, nil}, &BSTree{1, nil, nil}}, false},
		{&BSTree{1, &BSTree{0, nil, nil}, &BSTree{2, nil, nil}}, true},
		{&BSTree{1, &BSTree{2, nil, nil}, &BSTree{0, nil, nil}}, false},
		{&BSTree{2, &BSTree{0, nil, nil}, &BSTree{1, nil, nil}}, false},
		{&BSTree{2, &BSTree{1, nil, nil}, &BSTree{0, nil, nil}}, false},
		{&BSTree{3,
			&BSTree{2,
				&BSTree{1, nil, nil},
				&BSTree{4, nil, nil}},
			&BSTree{5, nil, nil}}, false},
		{&BSTree{19,
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
					&BSTree{53, nil, nil}}}}, true},
	} {
		if got := IsBinaryTreeBST(test.tree); got != test.want {
			t.Errorf("IsBinaryTreeBST(%v) = %t; want %t", test.tree, got, test.want)
		}
	}
}

func BenchmarkIsBinaryTreeBST(b *testing.B) {
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
		IsBinaryTreeBST(tree)
	}
}
