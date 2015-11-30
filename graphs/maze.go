// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package graphs

// isFeasible returns true if c is inside the maze and is white.
func isFeasible(maze [][]bool, c Coordinate) bool {
	return c.x >= 0 && c.x < len(maze) && c.y >= 0 && c.y < len(maze[c.x]) && maze[c.x][c.y] == white
}

// pathExists returns true if path from c to e exists in maze.
// Function is recursive implementation of DFS.
func pathExists(maze [][]bool, c, e Coordinate, path *[]Coordinate) bool {
	if c == e { // Base case.
		return true
	}

	for _, s := range []Coordinate{right, left, down, up} {
		n := Coordinate{c.x + s.x, c.y + s.y}
		if isFeasible(maze, n) { // Check if we can move this way.
			*path = append(*path, n)
			maze[n.x][n.y] = black // Mark as visited.
			if pathExists(maze, n, e, path) {
				return true
			}
			*path = (*path)[:len(*path)-1] // e wasn't found, take step back and try another one.
		}
	}
	return false
}

// SearchMaze searches for a path from s to e in the maze and return this path if exists.
// The time complexity is O(v+e) where v is the number of vertices and e is the
// number of edges in maze. Vertex is represented by boolean value true and edge
// connects two vertices. O(v) additional space is needed.
func SearchMaze(maze [][]bool, s, e Coordinate) (path []Coordinate) {
	if !isFeasible(maze, s) {
		return path
	}
	path = append(path, s)
	maze[s.x][s.y] = black // Mark as visited.
	if !pathExists(maze, s, e, &path) {
		return []Coordinate(nil)
	}
	return path
}
