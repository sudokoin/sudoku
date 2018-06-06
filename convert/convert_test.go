package convert_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/sudokoin/sudoku/convert"
)

func Example() {
	board := [9][9]int{
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

	bytes, err := convert.ToBytes(board)

	if err != nil {
		// board is not solved correctly
	}

	board2, err := convert.FromBytes(bytes)

	if err != nil {
		// bytes are corrupted
	}

	fmt.Println(reflect.DeepEqual(board, board2))
	// Output: true
}

var toBytesTests = []struct {
	id          string
	in          [9][9]int
	out         []byte
	errExpected bool
}{
	{
		id:  "working",
		in:  working,
		out: workingBytes,
	}, {
		id:  "working with 9 last",
		in:  working9last,
		out: workingBytes9last,
	}, {
		id:  "working with two subgrids starting with 9",
		in:  working9firstOf2Grids,
		out: workingBytes9firstOf2Grids,
	}, {
		id:  "working with ideal 9s",
		in:  workingIdeal9s,
		out: workingBytesIdeal9s,
	}, {
		id:          "encoding empty",
		in:          emptyBoard,
		out:         nil,
		errExpected: true,
	}, {
		id:          "row with two 9s",
		in:          rowWithTwo9s,
		out:         nil,
		errExpected: true,
	}, {
		id:          "with 0",
		in:          with0,
		out:         nil,
		errExpected: true,
	}, {
		id:          "with 10",
		in:          with10,
		out:         nil,
		errExpected: true,
	}, {
		id:          "with -1",
		in:          withMinus1,
		out:         nil,
		errExpected: true,
	}, {
		id:          "wrong cols",
		in:          wrongCols,
		out:         nil,
		errExpected: true,
	}, {
		id:          "wrong grids",
		in:          wrongGrids,
		out:         nil,
		errExpected: true,
	},
}

var fromBytesTests = []struct {
	id          string
	in          []byte
	out         [9][9]int
	errExpected bool
}{
	{
		id:  "working bytes",
		in:  workingBytes,
		out: working,
	}, {
		id:  "working bytes with 9 last",
		in:  workingBytes9last,
		out: working9last,
	}, {
		id:  "working bytes with two subgrids starting with 9",
		in:  workingBytes9firstOf2Grids,
		out: working9firstOf2Grids,
	}, {
		id:  "working bytes with ideal 9s",
		in:  workingBytesIdeal9s,
		out: workingIdeal9s,
	}, {
		id:          "empty bytes",
		in:          []byte{},
		out:         emptyBoard,
		errExpected: true,
	}, {
		id:          "short bytes",
		in:          []byte{1, 2, 3, 4, 5, 6, 7, 8},
		out:         emptyBoard,
		errExpected: true,
	}, {
		id:          "wrong bytes",
		in:          []byte{129, 153, 241, 95, 163, 70, 198, 136, 232, 143, 172, 174, 17, 156, 33, 114, 23, 185, 204, 239, 8, 35, 51},
		out:         emptyBoard,
		errExpected: true,
	}, {
		id:          "bytes leading to incorrect board",
		in:          []byte{129, 154, 241, 95, 172, 104, 216, 209, 29, 17, 245, 158, 231, 8, 206, 16, 185, 11, 220, 230, 119, 132, 17, 153, 208},
		out:         emptyBoard,
		errExpected: true,
	},
}

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
	workingBytes = []byte{113, 153, 241, 95, 163, 70, 198, 136, 232, 143, 172, 174, 17, 156, 33, 114, 23, 185, 204, 239, 9, 222, 17, 152}
	working9last = [9][9]int{
		{9, 8, 7, 6, 5, 4, 3, 2, 1},
		{6, 5, 4, 3, 2, 1, 9, 8, 7},
		{3, 2, 1, 9, 8, 7, 6, 5, 4},
		{8, 9, 6, 7, 4, 5, 2, 1, 3},
		{7, 4, 5, 2, 1, 3, 8, 9, 6},
		{2, 1, 3, 8, 9, 6, 7, 4, 5},
		{5, 7, 9, 4, 6, 8, 1, 3, 2},
		{1, 3, 2, 5, 7, 9, 4, 6, 8},
		{4, 6, 8, 1, 3, 2, 5, 7, 9},
	}
	workingBytes9last     = []byte{129, 153, 241, 95, 163, 70, 198, 136, 232, 143, 172, 174, 17, 156, 33, 114, 23, 185, 204, 239, 8, 35, 51, 160}
	working9firstOf2Grids = [9][9]int{
		{9, 8, 7, 6, 5, 4, 3, 2, 1},
		{6, 5, 4, 3, 2, 1, 9, 8, 7},
		{3, 2, 1, 8, 9, 7, 6, 5, 4},
		{2, 1, 3, 9, 8, 6, 7, 4, 5},
		{8, 9, 6, 7, 4, 5, 2, 1, 3},
		{7, 4, 5, 2, 1, 3, 8, 9, 6},
		{5, 7, 9, 4, 6, 8, 1, 3, 2},
		{1, 3, 2, 5, 7, 9, 4, 6, 8},
		{4, 6, 8, 1, 3, 2, 5, 7, 9},
	}
	workingBytes9firstOf2Grids = []byte{129, 163, 61, 95, 163, 70, 198, 136, 232, 143, 172, 11, 220, 253, 206, 17, 156, 33, 121, 157, 225, 4, 102, 116}
	workingIdeal9s             = [9][9]int{
		{6, 5, 4, 3, 2, 1, 9, 8, 7},
		{9, 8, 7, 6, 5, 4, 3, 2, 1},
		{3, 2, 1, 8, 9, 7, 6, 5, 4},
		{8, 9, 6, 7, 4, 5, 2, 1, 3},
		{2, 1, 3, 9, 8, 6, 7, 4, 5},
		{7, 4, 5, 2, 1, 3, 8, 9, 6},
		{5, 7, 9, 4, 6, 8, 1, 3, 2},
		{1, 3, 2, 5, 7, 9, 4, 6, 8},
		{4, 6, 8, 1, 3, 2, 5, 7, 9},
	}
	workingBytesIdeal9s = []byte{140, 33, 125, 88, 200, 255, 88, 209, 68, 125, 101, 112, 130, 23, 185, 231, 8, 94, 103, 120, 65, 25, 157}
	rowWithTwo9s        = [9][9]int{
		{9, 9, 7, 6, 5, 4, 3, 2, 1},
		{6, 5, 4, 3, 2, 1, 9, 8, 7},
		{3, 2, 1, 9, 8, 7, 6, 5, 4},
		{8, 9, 6, 7, 4, 5, 2, 1, 3},
		{7, 4, 5, 2, 1, 3, 8, 9, 6},
		{2, 1, 3, 8, 9, 6, 7, 4, 5},
		{5, 7, 9, 4, 6, 8, 1, 3, 2},
		{4, 6, 8, 1, 3, 2, 5, 7, 9},
		{1, 3, 2, 5, 7, 9, 4, 6, 8},
	}
	with0 = [9][9]int{
		{9, 8, 7, 6, 5, 4, 3, 2, 1},
		{6, 5, 4, 3, 2, 1, 9, 8, 7},
		{3, 2, 1, 9, 8, 7, 6, 5, 4},
		{8, 9, 6, 0, 4, 5, 2, 1, 3},
		{7, 4, 5, 2, 1, 3, 8, 9, 6},
		{2, 1, 3, 8, 9, 6, 7, 4, 5},
		{5, 7, 9, 4, 6, 8, 1, 3, 2},
		{4, 6, 8, 1, 3, 2, 5, 7, 9},
		{1, 3, 2, 5, 7, 9, 4, 6, 8},
	}
	with10 = [9][9]int{
		{9, 8, 7, 6, 5, 4, 3, 2, 1},
		{6, 5, 4, 3, 2, 1, 9, 8, 7},
		{3, 2, 1, 9, 8, 7, 6, 5, 4},
		{8, 9, 6, 7, 4, 5, 2, 1, 3},
		{7, 4, 5, 2, 10, 3, 8, 9, 6},
		{2, 1, 3, 8, 9, 6, 7, 4, 5},
		{5, 7, 9, 4, 6, 8, 1, 3, 2},
		{4, 6, 8, 1, 3, 2, 5, 7, 9},
		{1, 3, 2, 5, 7, 9, 4, 6, 8},
	}
	withMinus1 = [9][9]int{
		{9, 8, 7, 6, 5, 4, 3, 2, 1},
		{6, 5, 4, 3, 2, 1, 9, 8, 7},
		{3, 2, 1, 9, 8, 7, 6, 5, 4},
		{8, 9, 6, 7, 4, 5, 2, 1, 3},
		{7, 4, 5, 2, 1, 3, 8, 9, 6},
		{2, 1, 3, 8, 9, 6, 7, 4, 5},
		{5, -1, 9, 4, 6, 8, 1, 3, 2},
		{4, 6, 8, 1, 3, 2, 5, 7, 9},
		{1, 3, 2, 5, 7, 9, 4, 6, 8},
	}
	wrongCols = [9][9]int{
		{9, 8, 7, 6, 5, 4, 3, 2, 1},
		{6, 5, 4, 3, 2, 1, 9, 8, 7},
		{3, 2, 1, 9, 8, 7, 6, 5, 4},
		{9, 8, 6, 7, 5, 4, 3, 2, 1},
		{7, 4, 5, 2, 1, 3, 8, 9, 6},
		{2, 1, 3, 8, 9, 6, 7, 4, 5},
		{5, 7, 9, 4, 6, 8, 1, 3, 2},
		{4, 6, 8, 1, 3, 2, 5, 7, 9},
		{1, 3, 2, 5, 7, 9, 4, 6, 8},
	}
	wrongGrids = [9][9]int{
		{9, 8, 7, 6, 5, 4, 3, 2, 1},
		{6, 5, 4, 3, 2, 1, 9, 8, 7},
		{8, 9, 6, 7, 4, 5, 2, 1, 3},
		{3, 2, 1, 9, 8, 7, 6, 5, 4},
		{7, 4, 5, 2, 1, 3, 8, 9, 6},
		{2, 1, 3, 8, 9, 6, 7, 4, 5},
		{5, 7, 9, 4, 6, 8, 1, 3, 2},
		{4, 6, 8, 1, 3, 2, 5, 7, 9},
		{1, 3, 2, 5, 7, 9, 4, 6, 8},
	}
)

func TestToBytes(t *testing.T) {
	for _, test := range toBytesTests {
		out, err := convert.ToBytes(test.in)
		if test.errExpected != (err != nil) {
			t.Errorf("unexpected error for %s:\n%v\n", test.id, err)
		}
		if !reflect.DeepEqual(test.out, out) {
			t.Errorf("unexpected output for %s:\n%v\n%v\n", test.id, test.out, out)
		}
	}
}

func TestFromBytes(t *testing.T) {
	for _, test := range fromBytesTests {
		out, err := convert.FromBytes(test.in)
		if test.errExpected != (err != nil) {
			t.Errorf("unexpected error for %s:\n%v\n", test.id, err)
		}
		if !reflect.DeepEqual(test.out, out) {
			t.Errorf("unexpected output for %s:\n%v\n%v\n", test.id, test.out, out)
		}
	}
}

func TestShort(t *testing.T) {
	expected := "a19a28a37a46a55a64a73a82a91b16b25b34b43b52b61b79b88b97c13c22c31c49c58c67c76c85c94d18d29d36d47d54d65d72d81d93e17e24e35e42e51e63e78e89e96f12f21f33f48f59f66f77f84f95g15g27g39g44g56g68g71g83g92h14h26h38h41h53h62h75h87h99i11i23i32i45i57i69i74i86i98"
	actual := convert.ToShort(working)
	if expected != actual {
		t.Errorf("Expected short notations to match:\n%s\n%s", expected, actual)
	}

	short := convert.ToShort(with0)
	parsed := convert.FromShort(short)
	if !reflect.DeepEqual(with0, parsed) {
		t.Errorf("Expected original to equal parsed board:\n%+v\n%+v", with0, parsed)
	}
}

func TestUltraShort(t *testing.T) {
	expected := "9876543265432198321987658967452174521389213896745794681346813257"
	actual, err := convert.ToUltraShort(working)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if expected != actual {
		t.Errorf("Expected short notations to match:\n%s\n%s", expected, actual)
	}
}
