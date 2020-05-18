package sudoku

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestNewOptions(t *testing.T) {
	op := newOptions()

	expL := 9
	colsL := len(op.cols)
	rowsL := len(op.rows)
	squaresL := len(op.squares)

	if colsL != expL || rowsL != expL || squaresL != expL {
		t.Errorf(
			"Expected len options: %v, got cols: %v, rows: %v, squares: %v",
			expL,
			colsL,
			rowsL,
			squaresL,
		)
	}

	if op.cols[5][6] != 7 {
		t.Errorf(
			"Expected option value: %v, got: %v",
			7,
			op.cols[5][6],
		)
	}
}

func TestRemoveOption(t *testing.T) {
	op := newOptions()

	op.removeOption(5, 2, 0, 0)

	expL := 8
	colL := len(op.rows[2])
	rowL := len(op.cols[0])
	squareL := len(op.squares[0])

	if colL != expL || rowL != expL || squareL != expL {
		t.Errorf(
			"Expected len options: %v, got col: %v, row: %v, square: %v",
			expL,
			colL,
			rowL,
			squareL,
		)
	}

	if op.cols[0][6] != 8 {
		t.Errorf(
			"Expected option value: %v, got: %v",
			8,
			op.cols[0][6],
		)
	}
}

func TestGetPointOptions(t *testing.T) {
	op := newOptions()

	op.removeOption(5, 2, 0, 0)
	op.removeOption(1, 2, 0, 0)
	op.removeOption(2, 2, 0, 0)
	op.removeOption(6, 1, 2, 0)
	op.removeOption(8, 2, 4, 2)
	op.removeOption(9, 4, 0, 3)

	if diff := cmp.Diff([]int{3, 4, 7}, op.getPointOptions(2, 0, 0)); diff != "" {
		t.Error("Resulting point options not match the expected slice", diff)
	}
}

func TestGetNeighborOptions(t *testing.T) {
	op := newOptions()

	op.removeOption(5, 0, 1, 0)
	op.removeOption(5, 1, 2, 0)

	diff := cmp.Diff([]int{5}, op.getNeighborOptions(0, 1, 1, 2, []int{1, 2, 3, 5, 9}))
	if diff != "" {
		t.Error("Resulting Neighbor options not match the expected slice", diff)
	}
}
