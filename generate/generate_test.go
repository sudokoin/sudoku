package generate

import (
	"testing"

	"github.com/sudokoin/sudoku/solve"
)

func TestGenerate(t *testing.T) {
	for minFields := 10; minFields < 81; minFields++ {
		board := SingleCandidate(Random(), minFields)
		symbolCount := 0
		for _, row := range board {
			for _, val := range row {
				if val != 0 {
					symbolCount++
				}
			}
		}
		if symbolCount < minFields {
			t.Errorf("expected at least %d fields: \n %v", minFields, board)
		}
		_, solved := solve.SolveSingleCandidate(board)
		if !solved {
			t.Errorf("expected board to be solvable: \n %v", board)
		}
	}
}
