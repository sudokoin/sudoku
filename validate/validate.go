// Package validate contains helpers to validate 9x9 sudokus.
package validate

import "sort"

// Symbols returns true iff all ints are in the range of 0-9.
func Symbols(board [9][9]int) bool {
	for _, row := range board {
		for _, val := range row {
			if val < 0 || val > 9 {
				return false
			}
		}
	}
	return true
}

// Solved returns true iff board is solved correctly.
func Solved(board [9][9]int) bool {
	for _, row := range board {
		if !validateGroup(row) {
			return false
		}
	}
	for colIdx := 0; colIdx < 9; colIdx++ {
		if !validateGroup(extractCol(board, colIdx)) {
			return false
		}
	}
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if !validateGroup(extractGrid(board, x, y)) {
				return false
			}
		}
	}
	return true
}

func validateGroup(group [9]int) bool {
	sorted := group[:]
	sort.Ints(sorted)
	for idx, val := range sorted {
		if val != idx+1 {
			return false
		}
	}
	return true
}

func extractCol(board [9][9]int, idx int) [9]int {
	return [9]int{
		board[0][idx],
		board[1][idx],
		board[2][idx],
		board[3][idx],
		board[4][idx],
		board[5][idx],
		board[6][idx],
		board[7][idx],
		board[8][idx],
	}
}

func extractGrid(board [9][9]int, x int, y int) [9]int {
	var grid [9]int
	var gridIdx int
	for rowIdx := x * 3; rowIdx < (x+1)*3; rowIdx++ {
		for colIdx := y * 3; colIdx < (y+1)*3; colIdx++ {
			grid[gridIdx] = board[rowIdx][colIdx]
			gridIdx++
		}
	}
	return grid
}
