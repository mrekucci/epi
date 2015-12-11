// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package graphs

// FlipColorDFS flips color on all x, y adjacent coordinates that has the same color.
// The time complexity is O(n*m) where n, m are dimensions of a.
// The space complexity is O(m+n).
// This version use graph DFS to do this job.
func FlipColorDFS(a [][]bool, x, y int) {
	color := a[x][y]
	a[x][y] = !color
	for _, shift := range []Coordinate{right, left, down, up} {
		next := Coordinate{x + shift.x, y + shift.y}
		if next.x >= 0 && next.x < len(a) && next.y >= 0 && next.y < len(a[next.x]) && a[next.x][next.y] == color {
			FlipColorDFS(a, next.x, next.y)
		}
	}
}

// FlipColorBFS flips color on all x, y adjacent coordinates that has the same color.
// The time complexity is O(n*m) where n, m are dimensions of a.
// The space complexity is O(m+n).
// This version use graph BFS to do this job.
func FlipColorBFS(a [][]bool, x, y int) {
	var q []Coordinate
	color := a[x][y]
	a[x][y] = !color
	q = append(q, Coordinate{x, y})
	for len(q) > 0 {
		curr := q[0]
		for _, shift := range []Coordinate{right, left, down, up} {
			next := Coordinate{curr.x + shift.x, curr.y + shift.y}
			if next.x >= 0 && next.x < len(a) && next.y >= 0 && next.y < len(a[next.x]) && a[next.x][next.y] == color {
				a[next.x][next.y] = !a[next.x][next.y]
				q = append(q, next)
			}
		}
		q = q[1:]
	}
}
