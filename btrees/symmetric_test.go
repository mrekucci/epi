// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package btrees

import "testing"

func TestIsSymmetric(t *testing.T) {
	for _, test := range []struct {
		tree *BTree
		want bool
	}{
		// Symmetric binary trees.
		{(*BTree)(nil), true},
		{&BTree{"TREE_0", nil, nil}, true},
		{&BTree{"TREE_1",
			&BTree{"A",
				&BTree{"B",
					nil,
					&BTree{"C", nil, nil}},
				nil},
			&BTree{"A",
				nil,
				&BTree{"B",
					&BTree{"C", nil, nil},
					nil}}}, true},

		// Non-symmetric binary trees.
		{&BTree{"TREE_2", &BTree{"A", nil, nil}, nil}, false},
		{&BTree{"TREE_3", nil, &BTree{"A", nil, nil}}, false},
		{&BTree{"TREE_4", &BTree{"A", nil, nil}, &BTree{"B", nil, nil}}, false},
		{&BTree{"TREE_5",
			&BTree{"A",
				&BTree{"B", nil, nil},
				nil},
			&BTree{"C", nil, nil}}, false},
	} {
		if got := IsSymmetric(test.tree); got != test.want {
			t.Errorf("IsSymmetric(%v) = %t; want %t", test.tree, got, test.want)
		}
	}
}

func BenchmarkIsSymmetric(b *testing.B) {
	tree := &BTree{"ROOT",
		&BTree{"A",
			&BTree{"B",
				nil,
				&BTree{"C", nil, nil}},
			nil},
		&BTree{"A",
			nil,
			&BTree{"B",
				&BTree{"C", nil, nil},
				nil}}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsSymmetric(tree)
	}
}
