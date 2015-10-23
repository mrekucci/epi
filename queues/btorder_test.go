// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package queues

import (
	"reflect"
	"testing"
)

func TestDepthOrder(t *testing.T) {
	for _, test := range []struct {
		tree *IntBTree
		want [][]int
	}{
		{&IntBTree{1e1, nil, nil}, [][]int{[]int{1e1}}},
		{&IntBTree{1e2, &IntBTree{2e2, nil, nil}, nil}, [][]int{[]int{1e2}, []int{2e2}}},
		{&IntBTree{1e3, &IntBTree{2e3, &IntBTree{3e3, nil, nil}, nil}, nil}, [][]int{[]int{1e3}, []int{2e3}, []int{3e3}}},
		{&IntBTree{1e4,
			&IntBTree{2e4,
				&IntBTree{3e4,
					&IntBTree{4e4, nil, nil},
					&IntBTree{5e4, nil, nil}},
				&IntBTree{6e4, nil, nil}},
			&IntBTree{7e4, nil, nil}}, [][]int{[]int{1e4}, []int{2e4, 7e4}, []int{3e4, 6e4}, []int{4e4, 5e4}}},
		{&IntBTree{1e5,
			&IntBTree{2e5,
				&IntBTree{3e5,
					&IntBTree{4e5,
						&IntBTree{5e5, nil, nil},
						&IntBTree{6e5, nil, nil}},
					&IntBTree{7e5, nil, nil}},
				&IntBTree{8e5,
					&IntBTree{9e5, nil, nil},
					nil}},
			nil}, [][]int{[]int{1e5}, []int{2e5}, []int{3e5, 8e5}, []int{4e5, 7e5, 9e5}, []int{5e5, 6e5}}},
	} {
		if got := DepthOrder(test.tree); !reflect.DeepEqual(got, test.want) {
			t.Errorf("DepthOrder(%v) = %v; want %v", test.tree, got, test.want)
		}
	}
}

func BenchmarkDepthOrder(b *testing.B) {
	tree := &IntBTree{1e5,
		&IntBTree{2e5,
			&IntBTree{3e5,
				&IntBTree{4e5,
					&IntBTree{5e5, nil, nil},
					&IntBTree{6e5, nil, nil}},
				&IntBTree{7e5, nil, nil}},
			&IntBTree{8e5,
				&IntBTree{9e5, nil, nil},
				nil}},
		nil}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DepthOrder(tree)
	}
}
