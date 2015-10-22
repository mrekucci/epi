// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package btrees

import "testing"

func TestLCA(t *testing.T) {
	// A Binary Tree of height 5.
	tree := &BTree{"A",
		&BTree{"B",
			&BTree{"C",
				&BTree{"D", nil, nil},
				&BTree{"E", nil, nil}},
			&BTree{"F",
				nil,
				&BTree{"G",
					&BTree{"H", nil, nil},
					nil}}},
		&BTree{"I",
			&BTree{"J",
				nil,
				&BTree{"K",
					&BTree{"L",
						nil,
						&BTree{"M", nil, nil}},
					&BTree{"N", nil, nil}}},
			&BTree{"O",
				nil,
				&BTree{"P", nil, nil}}}}

	for _, test := range []struct {
		t, n0, n1 *BTree
		want      *BTree
	}{
		{&BTree{"TREE", nil, nil}, nil, nil, nil},
		{tree, tree.left, tree.right, tree},                                                          // n0 = B; n1 = I; want = A
		{tree, tree.right.left.right.left.right, tree.right.left.right.right, tree.right.left.right}, // n0 = M; n1 = N; want = K
	} {
		if got := LCA(test.t, test.n0, test.n1); got != test.want {
			t.Errorf("LCA(%v, %v, %v) = %v; want %v", test.t, test.n0, test.n1, got, test.want)
		}
	}
}
