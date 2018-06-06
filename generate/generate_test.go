package generate

import (
	"testing"
)

var (
	emptyBoard = [9][9]int{}
	working    = [9][9]int{
		{9, 8, 7, 6, 5, 4, 3, 2, 1},
		{6, 5, 4, 3, 2, 1, 9, 8, 7},
		{3, 2, 1, 9, 8, 7, 6, 5, 4},
		{8, 9, 6, 7, 4, 5, 2, 1, 3},
		{7, 4, 5, 2, 1, 3, 8, 9, 6},
		{2, 1, 3, 8, 9, 6, 7, 4, 5},
		{5, 7, 9, 4, 6, 8, 1, 3, 2},
		{4, 6, 8, 1, 3, 2, 5, 7, 9},
		{1, 3, 2, 5, 7, 9, 4, 6, 8},
	}
	solvable = [9][9]int{
		{9, 0, 0, 6, 5, 4, 3, 2, 1},
		{6, 5, 4, 3, 2, 1, 9, 8, 7},
		{3, 2, 1, 9, 8, 7, 6, 5, 4},
		{8, 9, 6, 7, 4, 5, 2, 1, 3},
		{7, 4, 5, 2, 1, 3, 8, 9, 6},
		{2, 1, 3, 8, 9, 6, 7, 4, 5},
		{5, 0, 9, 4, 6, 8, 1, 3, 2},
		{4, 6, 8, 1, 3, 2, 5, 7, 9},
		{1, 3, 2, 5, 7, 9, 4, 6, 8},
	}
	unsolvable = [9][9]int{
		{9, 0, 7, 0, 0, 4, 3, 0, 1},
		{6, 0, 4, 3, 0, 1, 9, 0, 7},
		{3, 0, 1, 9, 0, 7, 6, 0, 4},
		{0, 9, 6, 7, 4, 5, 2, 1, 3},
		{7, 0, 5, 2, 1, 3, 0, 9, 6},
		{2, 1, 3, 8, 9, 0, 7, 4, 5},
		{5, 7, 9, 4, 0, 0, 1, 0, 2},
		{4, 0, 0, 1, 3, 2, 5, 7, 9},
		{1, 0, 2, 5, 7, 9, 4, 0, 0},
	}
)

func TestGenerate(t *testing.T) {
	for minFields := 10; minFields < 81; minFields++ {
		board := GenerateSingleCandidate(minFields)
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
		_, solved := solveSingleCandidate(board)
		if !solved {
			t.Errorf("expected board to be solvable: \n %v", board)
		}
	}
}

var solveSingleCandidateTests = []struct {
	in          [9][9]int
	solution    [9][9]int
	success     bool
	errExpected bool
}{
	{
		in:       working,
		solution: working,
		success:  true,
	},
	{
		in: unsolvable,
		solution: [9][9]int{
			{9, 0, 7, 6, 0, 4, 3, 0, 1},
			{6, 0, 4, 3, 0, 1, 9, 0, 7},
			{3, 0, 1, 9, 0, 7, 6, 0, 4},
			{8, 9, 6, 7, 4, 5, 2, 1, 3},
			{7, 4, 5, 2, 1, 3, 8, 9, 6},
			{2, 1, 3, 8, 9, 6, 7, 4, 5},
			{5, 7, 9, 4, 6, 8, 1, 3, 2},
			{4, 6, 8, 1, 3, 2, 5, 7, 9},
			{1, 3, 2, 5, 7, 9, 4, 6, 8},
		},
		success: false,
	},
}

func TestSolveSingleCandidate(t *testing.T) {
	for _, test := range solveSingleCandidateTests {
		solution, success := solveSingleCandidate(test.in)
		if test.solution != solution {
			t.Errorf("unexpected solution:\n%d\n%d\n", test.solution, solution)
		}
		if test.success != success {
			t.Errorf("unexpected success bool:\n%v\n%v\n", test.success, success)
		}
	}
}

var annotateSingleCandidateTests = []struct {
	in  [9][9]int
	out annotated
}{
	{
		in: emptyBoard,
		out: annotated{
			blocks: [3][3]uint{{1022, 1022, 1022}, {1022, 1022, 1022}, {1022, 1022, 1022}},
			cols:   [9]uint{1022, 1022, 1022, 1022, 1022, 1022, 1022, 1022, 1022},
			rows:   [9]uint{1022, 1022, 1022, 1022, 1022, 1022, 1022, 1022, 1022},
			fields: [9][9]uint{
				{1022, 1022, 1022, 1022, 1022, 1022, 1022, 1022, 1022},
				{1022, 1022, 1022, 1022, 1022, 1022, 1022, 1022, 1022},
				{1022, 1022, 1022, 1022, 1022, 1022, 1022, 1022, 1022},
				{1022, 1022, 1022, 1022, 1022, 1022, 1022, 1022, 1022},
				{1022, 1022, 1022, 1022, 1022, 1022, 1022, 1022, 1022},
				{1022, 1022, 1022, 1022, 1022, 1022, 1022, 1022, 1022},
				{1022, 1022, 1022, 1022, 1022, 1022, 1022, 1022, 1022},
				{1022, 1022, 1022, 1022, 1022, 1022, 1022, 1022, 1022},
				{1022, 1022, 1022, 1022, 1022, 1022, 1022, 1022, 1022},
			},
		},
	},
	{
		in: working,
		out: annotated{
			blocks: [3][3]uint{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			cols:   [9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
			rows:   [9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
			fields: [9][9]uint{
				{512, 256, 128, 64, 32, 16, 8, 4, 2},
				{64, 32, 16, 8, 4, 2, 512, 256, 128},
				{8, 4, 2, 512, 256, 128, 64, 32, 16},
				{256, 512, 64, 128, 16, 32, 4, 2, 8},
				{128, 16, 32, 4, 2, 8, 256, 512, 64},
				{4, 2, 8, 256, 512, 64, 128, 16, 32},
				{32, 128, 512, 16, 64, 256, 2, 8, 4},
				{16, 64, 256, 2, 8, 4, 32, 128, 512},
				{2, 8, 4, 32, 128, 512, 16, 64, 256},
			},
		},
	},
	{
		in: unsolvable,
		out: annotated{
			blocks: [3][3]uint{{292, 356, 292}, {272, 64, 256}, {328, 320, 328}},
			cols:   [9]uint{256, 380, 256, 64, 356, 320, 256, 364, 256},
			rows:   [9]uint{356, 292, 292, 256, 272, 64, 328, 320, 328},
			fields: [9][9]uint{
				{512, 292, 128, 64, 356, 16, 8, 292, 2},
				{64, 292, 16, 8, 292, 2, 512, 292, 128},
				{8, 292, 2, 512, 292, 128, 64, 292, 16},
				{256, 512, 64, 128, 16, 32, 4, 2, 8},
				{128, 272, 32, 4, 2, 8, 256, 512, 64},
				{4, 2, 8, 256, 512, 64, 128, 16, 32},
				{32, 128, 512, 16, 320, 320, 2, 328, 4},
				{16, 320, 256, 2, 8, 4, 32, 128, 512},
				{2, 328, 4, 32, 128, 512, 16, 328, 256},
			},
		},
	},
}

func TestAnnotateSingleCandidate(t *testing.T) {
	for _, test := range annotateSingleCandidateTests {
		actual := annotateSingleCandidate(test.in)
		expected := test.out
		if expected != actual {
			t.Errorf("unexpected output:\n%d\n%d\n", expected, actual)
		}
	}
}