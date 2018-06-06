// Package generate contains helpers to generate unsolved 9x9 sudokus.
package generate

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/sudokoin/sudoku/validate"
)

// GenerateSingleCandidate creates a board that can be solved with single candidate strategy.
func GenerateSingleCandidate(minFields int) [9][9]int {
	if minFields < 0 || minFields > 80 {
		minFields = 10 // minimum found sudoku is 17 right now, 10 is for safety.
	}
	rb := randomBoard()
	fields := randomFields(rb)
	board := fillTillMinimum(fields, minFields)
	board = fillTillSolvableSingleCandidate(board, fields[minFields:])
	return board
}

func fillTillSolvableSingleCandidate(board [9][9]int, fields [][3]int) [9][9]int {
	_, solved := solveSingleCandidate(board)
	fIdx := 0
	for !solved {
		f := fields[fIdx]
		board[f[0]][f[1]] = f[2]
		fIdx++
		_, solved = solveSingleCandidate(board)
	}
	return board
}

func fillTillMinimum(fields [][3]int, minFields int) [9][9]int {
	board := [9][9]int{}
	for fIdx, f := range fields {
		if fIdx < minFields {
			board[f[0]][f[1]] = f[2]
		}
	}
	return board
}

func randomFields(board [9][9]int) [][3]int {
	fields := [][3]int{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rf := r.Perm(81)
	for _, rIdx := range rf {
		rowIdx := rIdx / 9
		colIdx := rIdx % 9
		fields = append(fields, [3]int{rowIdx, colIdx, board[rowIdx][colIdx]})
	}
	return fields
}

func randomBoard() [9][9]int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	board := [9][9]int{}
	copy(board[0][:], r.Perm(9))
	for colIdx, _ := range board[0] {
		board[0][colIdx]++
	}
	solved, solutions := backtrackEntry(board, 10)
	l := len(solutions)
	if !solved || l < 1 {
		panic(fmt.Sprintf("unanticipated problem with solving board: %v\n", board))
	}
	return solutions[r.Intn(l)]
}

// solveSingleCandidate tries to solve a board with single candidate strategy.
// The returned bool indicates whether it was successful.
func solveSingleCandidate(board [9][9]int) ([9][9]int, bool) {
	next := annotateSingleCandidate(board).toBoard()
	for board != next {
		board = next
		next = annotateSingleCandidate(board).toBoard()
	}
	return board, validate.Solved(board)
}

const (
	all uint = 1022 // bits 1-9 are set (1111111110)
)

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

// backtrack implements a simple backtracking solver. It is not performant but guaranteed to finish.
func backtrackEntry(board [9][9]int, maxSolutions int) (solved bool, solutions [][9][9]int) {
	return backtrack(board, maxSolutions, &solutions), solutions
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
