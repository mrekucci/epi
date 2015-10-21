// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package btrees

import "testing"

func TestIsBalance(t *testing.T) {
	for _, test := range []struct {
		tree *BTree
		want bool
	}{
		// Balanced Binary Trees.
		{&BTree{"TREE_0", nil, nil}, true},
		{&BTree{"TREE_1", &BTree{"A", nil, nil}, nil}, true},
		{&BTree{"TREE_2",
			&BTree{"A",
				&BTree{"B", nil, nil},
				&BTree{"C", nil, nil}},
			&BTree{"D", nil, nil}}, true},

		// Not balanced Binary Trees.
		{&BTree{"TREE_3", &BTree{"A", &BTree{"B", nil, nil}, nil}, nil}, false},
		{&BTree{"TREE_4",
			&BTree{"A",
				&BTree{"B",
					&BTree{"C", nil, nil},
					&BTree{"D", nil, nil}},
				&BTree{"E", nil, nil}},
			&BTree{"F", nil, nil}}, false},
		{&BTree{"TREE_5",
			&BTree{"A",
				&BTree{"B",
					&BTree{"C",
						&BTree{"D", nil, nil},
						&BTree{"E", nil, nil}},
					&BTree{"F", nil, nil}},
				&BTree{"G", nil, nil}},
			&BTree{"H", nil, nil}}, false},
		{&BTree{"TREE_6",
			&BTree{"A", nil, nil},
			&BTree{"B",
				&BTree{"C", nil, nil},
				&BTree{"D",
					&BTree{"E", nil, nil},
					&BTree{"F", nil, nil}}}}, false},
		{&BTree{"TREE_7",
			&BTree{"A", nil, nil},
			&BTree{"B",
				&BTree{"C", nil, nil},
				&BTree{"D",
					&BTree{"E", nil, nil},
					&BTree{"F",
						&BTree{"G", nil, nil},
						&BTree{"H", nil, nil}}}}}, false},
	} {
		if got := IsBalanced(test.tree); got != test.want {
			t.Errorf("IsBalanced(%v) = %t; want %t", test.tree, got, test.want)
		}
	}
}

func BenchmarkIsBalanced(b *testing.B) {
	tree := &BTree{"A",
		&BTree{"B",
			&BTree{"C",
				&BTree{"D",
					&BTree{"E", nil, nil},
					&BTree{"F", nil, nil}},
				&BTree{"G", nil, nil}},
			&BTree{"H",
				&BTree{"I", nil, nil},
				&BTree{"J", nil, nil}}},
		&BTree{"K",
			&BTree{"L",
				&BTree{"M", nil, nil},
				&BTree{"N", nil, nil}},
			&BTree{"O", nil, nil}}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsBalanced(tree)
	}
}
