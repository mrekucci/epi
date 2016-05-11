// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package recursion

import (
	"reflect"
	"testing"
)

func TestAllSubsets(t *testing.T) {
	for _, test := range []struct {
		n, k int
		want [][]int
	}{
		{0, 1, nil},
		{1, 0, nil},
		{1, 2, nil},
		{2, 2, [][]int{{1, 2}}},
		{3, 2, [][]int{{1, 2}, {1, 3}, {2, 3}}},
		{4, 2, [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
		{5, 1, [][]int{{1}, {2}, {3}, {4}, {5}}},
		{5, 2, [][]int{{1, 2}, {1, 3}, {1, 4}, {1, 5}, {2, 3}, {2, 4}, {2, 5}, {3, 4}, {3, 5}, {4, 5}}},
		{5, 3, [][]int{{1, 2, 3}, {1, 2, 4}, {1, 2, 5}, {1, 3, 4}, {1, 3, 5}, {1, 4, 5}, {2, 3, 4}, {2, 3, 5}, {2, 4, 5}, {3, 4, 5}}},
	} {
		if got := AllSubsets(test.n, test.k); !reflect.DeepEqual(got, test.want) {
			t.Errorf("AllSubsets(%d, %d)\n got %v\nwant %v", test.n, test.k, got, test.want)
		}
	}
}

// TODO: add benchmark!
