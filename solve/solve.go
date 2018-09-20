package solve

import "github.com/sudokoin/sudoku/validate"

const (
	all uint = 1022 // bits 1-9 are set (1111111110)
)

// Backtrack implements a simple backtracking solver. It is not performant but guaranteed to finish.
func Backtrack(board [9][9]int, maxSolutions int) (bool, [][9][9]int) {
	solutions := [][9][9]int{}
	return backtrack(board, maxSolutions, &solutions), solutions
}

// SolveSingleCandidate tries to solve a board with single candidate strategy.
// The returned bool indicates whether it was successful.
func SolveSingleCandidate(board [9][9]int) ([9][9]int, bool) {
	next := annotateSingleCandidate(board).toBoard()
	for board != next {
		board = next
		next = annotateSingleCandidate(board).toBoard()
	}
	return board, validate.Solved(board)
}

func backtrack(board [9][9]int, maxSolutions int, solutions *[][9][9]int) bool {
	rowIdx, colIdx, found := firstEmpty(board)
	if !found {
		*solutions = append(*solutions, board)
		return len(*solutions) >= maxSolutions
	}
	an := annotateSingleCandidate(board)
	for _, v := range allSymbols(an.fields[rowIdx][colIdx]) {
		board[rowIdx][colIdx] = v
		if backtrack(board, maxSolutions, solutions) {
			return true
		}
	}
	return false
}

func allSymbols(bits uint) []int {
	syms := []int{}
	for idx, msk := range [9]uint{2, 4, 8, 16, 32, 64, 128, 256, 512} {
		if msk&bits == msk {
			syms = append(syms, idx+1)
		}
	}
	return syms
}

func firstEmpty(board [9][9]int) (int, int, bool) {
	for rowIdx, row := range board {
		for colIdx, val := range row {
			if val == 0 {
				return rowIdx, colIdx, true
			}
		}
	}
	return 0, 0, false
}

type annotated struct {
	blocks [3][3]uint
	cols   [9]uint
	rows   [9]uint
	fields [9][9]uint
}

func (an annotated) toBoard() [9][9]int {
	board := [9][9]int{}
	for rowIdx, row := range an.fields {
		for colIdx, val := range row {
			board[rowIdx][colIdx] = firstSymbol(val)
		}
	}
	return board
}

func firstSymbol(bits uint) int {
	for idx, msk := range [10]uint{0, 2, 4, 8, 16, 32, 64, 128, 256, 512} {
		if msk == bits {
			return idx
		}
	}
	return 0
}

func annotateSingleCandidate(board [9][9]int) annotated {
	an := annotated{
		blocks: [3][3]uint{{all, all, all}, {all, all, all}, {all, all, all}},
		cols:   [9]uint{all, all, all, all, all, all, all, all, all},
		rows:   [9]uint{all, all, all, all, all, all, all, all, all},
	}
	for rowIdx, row := range board {
		for colIdx, val := range row {
			var bit uint = toBit(val)
			an.blocks[rowIdx/3][colIdx/3] = an.blocks[rowIdx/3][colIdx/3] &^ bit
			an.cols[colIdx] = an.cols[colIdx] &^ bit
			an.rows[rowIdx] = an.rows[rowIdx] &^ bit
		}
	}
	for rowIdx, row := range board {
		for colIdx, val := range row {
			if val != 0 {
				an.fields[rowIdx][colIdx] = toBit(val)
			} else {
				an.fields[rowIdx][colIdx] = an.rows[rowIdx] & an.cols[colIdx] & an.blocks[rowIdx/3][colIdx/3]
			}
		}
	}
	return an
}

func toBit(i int) uint {
	return 1 << uint(i)
}
