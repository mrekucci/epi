// Copyright (c) 2016, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package recursion

// SolveSudoku solves sudoku grid and returns true if grid was solved successfully.
// Since there is no notion of scaling with a size, the time and space complexity
// are irrelevant here, because the function allows only grids of size 9x9.
// In general, the problem of solving n*n sudoku grids is NP-complete, and has
// exponential complexity.
func SolveSudoku(grid [][]int) bool {
	if len(grid) != 9 || len(grid[0]) != 9 {
		return false
	}
	return solve(0, 0, grid)
}

func solve(i, j int, grid [][]int) bool {
	row, col, done := nextUnassigned(i, j, grid)
	if done {
		return true
	}

	for v := 1; v <= len(grid); v++ {
		if isValidPlacement(v, row, col, grid) {
			grid[row][col] = v // Assign.
			if solve(row, col, grid) {
				return true
			}
			grid[row][col] = 0 // Undo.
		}
	}
	return false
}

func nextUnassigned(row, col int, grid [][]int) (i, j int, done bool) {
	for row < len(grid) {
		for col < len(grid[row]) {
			if grid[row][col] == 0 {
				return row, col, false
			}
			col++
		}
		col = 0
		row++
	}
	return 0, 0, true
}

func isValidPlacement(v, i, j int, grid [][]int) bool {
	for row := 0; row < len(grid); row++ {
		if grid[row][j] == v {
			return false
		}
	}
	for col := 0; col < len(grid[i]); col++ {
		if grid[i][col] == v {
			return false
		}
	}
	// Check region of 3x3.
	i, j = i/3, j/3
	for row := 0; row*row < len(grid); row++ {
		for col := 0; col*col < len(grid[row]); col++ {
			if v == grid[3*i+row][3*j+col] {
				return false
			}
		}
	}
	return true
}
