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
	for _, s := range []Coordinate{right, left, down, up} {
		nx, ny := x+s.x, y+s.y
		if nx >= 0 && nx < len(a) && ny >= 0 && ny < len(a[nx]) && a[nx][ny] == color {
			FlipColorDFS(a, nx, ny)
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
		c := q[0]
		for _, s := range []Coordinate{right, left, down, up} {
			n := Coordinate{c.x + s.x, c.y + s.y}
			if n.x >= 0 && n.x < len(a) && n.y >= 0 && n.y < len(a[n.x]) && a[n.x][n.y] == color {
				a[n.x][n.y] = !a[n.x][n.y]
				q = append(q, n)
			}
		}
		q = q[1:]
	}
}