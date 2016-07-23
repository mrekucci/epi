// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package recursion

// isSafe returns true if the last added queen into the
// placement won't conflict any earlier queens placed before.
func isSafe(placement []int) bool {
	row := len(placement) - 1 // placement indices represents rows indices.
	for i := 0; i < row; i++ {
		dir := placement[i] - placement[row]
		// on the same column || on the left-down diagonal || on the right-down diagonal
		if dir == 0 || -dir == row-i || dir == row-i {
			return false
		}
	}
	return true
}

// NQueens returns 2D slice that represents all distinct non-attacking
// placements of n queens onto n*n board. In every sub-slice the indices
// represents row index and the value represents column index.
// The time complexity is lower bounded by the number of non-attacking
// placements. No exact form is known for this quantity as a function of
// n, but it's conjectured to tend to n!/(c**n), where c is approximately
// 2.52, which is super-exponential. Beyond the returned n*n board O(1)
// additional space is needed.
func NQueens(n int) (positions [][]int) {
	if n == 0 {
		return nil
	}

	// solve use backtracking technique to place
	// queens on board to the non-attacking positions.
	var solve func(row int, queens []int)
	solve = func(row int, queens []int) {
		if row == n { // Base case, all queens are placed on non-attacking positions.
			positions = append(positions, append([]int(nil), queens...))
			return
		}
		for col := 0; col < n; col++ {
			queens = append(queens, col) // Place queen.
			if isSafe(queens) {
				solve(row+1, queens)
			}
			queens = queens[:len(queens)-1] // Step back, remove queen and go try another position.
		}
	}

	solve(0, nil)
	return positions
}
