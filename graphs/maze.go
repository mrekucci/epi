// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package graphs

// isFeasible returns true if c is inside the maze and is white.
func isFeasible(maze [][]bool, c Coordinate) bool {
	return c.x >= 0 && c.x < len(maze) && c.y >= 0 && c.y < len(maze[c.x]) && maze[c.x][c.y] == white
}

// pathExists returns true if path from curr to end exists in maze.
// Function is recursive implementation of DFS.
func pathExists(maze [][]bool, curr, end Coordinate, path *[]Coordinate) bool {
	if curr == end { // Base case.
		return true
	}

	for _, shift := range []Coordinate{right, left, down, up} {
		next := Coordinate{curr.x + shift.x, curr.y + shift.y}
		if isFeasible(maze, next) { // Check if we can move this way.
			*path = append(*path, next)
			maze[next.x][next.y] = black // Mark as visited.
			if pathExists(maze, next, end, path) {
				return true
			}
			*path = (*path)[:len(*path)-1] // e wasn't found, take step back and try another one.
		}
	}
	return false
}

// SearchMaze returns true if a path exists from start to end in the maze.
// The time complexity is O(v+e) where v is the number of vertices and e is the
// number of edges in maze. Vertex is represented by boolean value true and edge
// connects two vertices. O(v) additional space is needed.
func SearchMaze(maze [][]bool, start, end Coordinate) (path []Coordinate) {
	if !isFeasible(maze, start) {
		return path
	}
	path = append(path, start)
	maze[start.x][start.y] = black // Mark as visited.
	if !pathExists(maze, start, end, &path) {
		return []Coordinate(nil)
	}
	return path
}
