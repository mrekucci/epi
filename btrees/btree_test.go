// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package btrees

import (
	"reflect"
	"testing"
)

func TestHeight(t *testing.T) {
	for _, test := range []struct {
		tree *BTree
		want int
	}{
		{&BTree{"TREE_0", nil, nil}, 0},
		{&BTree{"TREE_1", &BTree{"A", nil, nil}, nil}, 1},
		{&BTree{"TREE_2", nil, &BTree{"A", nil, nil}}, 1},
		{&BTree{"TREE_3", &BTree{"A", nil, nil}, &BTree{"B", nil, nil}}, 1},
		{&BTree{"TREE_4", &BTree{"A", &BTree{"B", nil, nil}, nil}, nil}, 2},
	} {
		if got := Height(test.tree); got != test.want {
			t.Errorf("Height(%v) = %d; want %d", test.tree, got, test.want)
		}
	}
}

func TestWalk(t *testing.T) {
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
		fnName string
		fn     func(t *BTree, w *[]interface{})
		want   []interface{}
	}{
		{"Preorder", Preorder, []interface{}{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P"}},
		{"Inorder", Inorder, []interface{}{"D", "C", "E", "B", "F", "H", "G", "A", "J", "L", "M", "K", "N", "I", "O", "P"}},
		{"Postorder", Postorder, []interface{}{"D", "E", "C", "H", "G", "F", "B", "M", "L", "N", "K", "J", "P", "O", "I", "A"}},
	} {
		if got := Walk(tree, test.fn); !reflect.DeepEqual(got, test.want) {
			t.Errorf("tree.Walk %q got %v; want %v", test.fnName, got, test.want)
		}
	}
}
