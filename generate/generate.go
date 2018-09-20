// Package generate contains helpers to generate unsolved 9x9 sudokus.
package generate

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/sudokoin/sudoku/solve"
)

// Random generates a random solved sudoku.
func Random() [9][9]int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	board := [9][9]int{}
	copy(board[0][:], r.Perm(9))
	for colIdx, _ := range board[0] {
		board[0][colIdx]++
	}
	solved, solutions := solve.Backtrack(board, 10)
	l := len(solutions)
	if !solved || l < 1 {
		panic(fmt.Sprintf("unanticipated problem with solving board: %v\n", board))
	}
	return solutions[r.Intn(l)]
}

// SingleCandidate derives a sudoku that can be solved with single candidate strategy
// from provided solved board.
func SingleCandidate(board [9][9]int, minFields int) [9][9]int {
	if minFields < 0 || minFields > 80 {
		minFields = 10 // minimum found sudoku is 17 right now, 10 is for safety.
	}
	fields := randomFields(board)
	unsolved := fillTillMinimum(fields, minFields)
	unsolved = fillTillSolvableSingleCandidate(unsolved, fields[minFields:])
	return unsolved
}

func fillTillSolvableSingleCandidate(board [9][9]int, fields [][3]int) [9][9]int {
	_, solved := solve.SolveSingleCandidate(board)
	fIdx := 0
	for !solved {
		f := fields[fIdx]
		board[f[0]][f[1]] = f[2]
		fIdx++
		_, solved = solve.SolveSingleCandidate(board)
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
