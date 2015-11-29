// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package graphs

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestSearchMaze(t *testing.T) {
	for _, test := range []struct {
		maze [][]bool
		s, e Coordinate
		want []Coordinate
	}{
		// 1x1 black
		{[][]bool{{black}},
			Coordinate{0, 0}, Coordinate{0, 0},
			[]Coordinate(nil)},
		// 1x1 white
		{[][]bool{{white}},
			Coordinate{0, 0}, Coordinate{0, 0},
			[]Coordinate{{0, 0}}},
		// 2x2
		{[][]bool{
			{black, white},
			{white, black}},
			Coordinate{1, 0}, Coordinate{0, 1},
			[]Coordinate(nil)},
		// 3x3
		{[][]bool{
			{black, white, white},
			{white, white, black},
			{white, black, white}},
			Coordinate{2, 0}, Coordinate{0, 2},
			[]Coordinate{{2, 0}, {1, 0}, {1, 1}, {0, 1}, {0, 2}}},
		// 10x10
		{[][]bool{
			{black, white, white, white, white, white, black, black, white, white},
			{white, white, black, white, white, white, white, white, white, white},
			{black, white, black, white, white, black, black, white, black, black},
			{white, white, white, black, black, black, white, white, black, white},
			{white, black, black, white, white, white, white, white, white, white},
			{white, black, black, white, white, black, white, black, black, white},
			{white, white, white, white, black, white, white, white, white, white},
			{black, white, black, white, black, white, black, white, white, white},
			{black, white, black, black, white, white, white, black, black, black},
			{white, white, white, white, white, white, white, black, black, white}},
			Coordinate{9, 0}, Coordinate{0, 9},
			[]Coordinate{
				{9, 0}, {9, 1}, {9, 2}, {9, 3}, {9, 4}, {9, 5}, {9, 6}, {8, 6}, {8, 5},
				{7, 5}, {6, 5}, {6, 6}, {6, 7}, {6, 8}, {6, 9}, {5, 9}, {4, 9}, {4, 8},
				{4, 7}, {4, 6}, {4, 5}, {4, 4}, {4, 3}, {5, 3}, {6, 3}, {6, 2}, {6, 1},
				{6, 0}, {5, 0}, {4, 0}, {3, 0}, {3, 1}, {2, 1}, {1, 1}, {0, 1}, {0, 2},
				{0, 3}, {0, 4}, {0, 5}, {1, 5}, {1, 6}, {1, 7}, {1, 8}, {1, 9}, {0, 9}}},
	} {
		if got := SearchMaze(test.maze, test.s, test.e); !reflect.DeepEqual(got, test.want) {
			t.Errorf("SearchMaze(%v, %v, %v)", test.maze, test.s, test.e)
			t.Error(" got:", got)
			t.Error("want:", test.want)
		}
	}
}

func benchSearchMaze(b *testing.B, size int) {
	var maze [][]bool
	var white []Coordinate
	g := rand.New(rand.NewSource(int64(size)))
	for x := 0; x < size; x++ {
		maze = append(maze, make([]bool, size))
		for y := 0; y < size; y++ {
			maze[x][y] = g.Intn(3)%2 == 0
			if maze[x][y] {
				white = append(white, Coordinate{x, y})
			}
		}
	}
	s, e := Coordinate{size - 1, 0}, Coordinate{0, size - 1}
	if len(white) > 0 {
		s, e = white[0], white[len(white)-1]
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SearchMaze(maze, s, e)
	}
}

func BenchmarkSearchMaze1e1(b *testing.B) { benchSearchMaze(b, 1e1) }
func BenchmarkSearchMaze1e2(b *testing.B) { benchSearchMaze(b, 1e2) }
func BenchmarkSearchMaze1e3(b *testing.B) { benchSearchMaze(b, 1e3) }
