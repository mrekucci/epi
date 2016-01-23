// Copyright (c) 2016, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package sorting

import (
	"reflect"
	"testing"
)

func TestUnionOfIntervals(t *testing.T) {
	for _, test := range []struct {
		in   []Interval
		want []Interval
	}{
		{nil, nil},
		{
			[]Interval{
				{EndPoint{true, 0}, EndPoint{true, 4}},
				{EndPoint{true, 0}, EndPoint{true, 5}}},
			[]Interval{{EndPoint{true, 0}, EndPoint{true, 5}}},
		},
		{
			[]Interval{
				{EndPoint{true, 1}, EndPoint{true, 5}},
				{EndPoint{true, 0}, EndPoint{true, 4}}},
			[]Interval{{EndPoint{true, 0}, EndPoint{true, 5}}},
		},
		{
			[]Interval{
				{EndPoint{true, 0}, EndPoint{true, 5}},
				{EndPoint{true, 0}, EndPoint{true, 4}}},
			[]Interval{{EndPoint{true, 0}, EndPoint{true, 5}}},
		},
		{
			[]Interval{
				{EndPoint{true, 1}, EndPoint{true, 3}},
				{EndPoint{true, 0}, EndPoint{true, 4}}},
			[]Interval{{EndPoint{true, 0}, EndPoint{true, 4}}},
		},
		{
			[]Interval{
				{EndPoint{true, 1}, EndPoint{true, 5}},
				{EndPoint{true, 0}, EndPoint{true, 4}}},
			[]Interval{{EndPoint{true, 0}, EndPoint{true, 5}}},
		},
		{
			[]Interval{
				{EndPoint{true, 2}, EndPoint{true, 4}},
				{EndPoint{true, 8}, EndPoint{false, 11}},
				{EndPoint{false, 13}, EndPoint{false, 15}},
				{EndPoint{false, 16}, EndPoint{false, 17}},
				{EndPoint{true, 1}, EndPoint{true, 1}},
				{EndPoint{true, 3}, EndPoint{false, 4}},
				{EndPoint{true, 7}, EndPoint{false, 8}},
				{EndPoint{false, 12}, EndPoint{true, 16}},
				{EndPoint{false, 0}, EndPoint{false, 3}},
				{EndPoint{true, 5}, EndPoint{false, 7}},
				{EndPoint{false, 9}, EndPoint{true, 11}},
				{EndPoint{true, 12}, EndPoint{true, 14}}},
			[]Interval{
				{EndPoint{false, 0}, EndPoint{true, 4}},
				{EndPoint{true, 5}, EndPoint{true, 11}},
				{EndPoint{true, 12}, EndPoint{false, 17}}},
		},
	} {
		if got := UnionOfIntervals(test.in); !reflect.DeepEqual(got, test.want) {
			t.Errorf("UnionOfIntervals(%v) = %v; want %v", test.in, got, test.want)
		}
	}
}

// TODO: benchmark!
