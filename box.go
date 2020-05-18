package sudoku

type box struct {
	value           int
	rowIdx          int
	rowNeighborIdx1 int
	rowNeighborIdx2 int
	colIdx          int
	colNeighborIdx1 int
	colNeighborIdx2 int
	squareIdx       int
	options         []int
}

func newBox(idx, value int) *box {
	getNeighborsIdx := func(idx int) (idx1 int, idx2 int) {
		switch idx % 3 {
		case 0:
			return idx + 1, idx + 2
		case 1:
			return idx - 1, idx + 1
		default:
			return idx - 2, idx - 1
		}
	}

	rowIdx := (idx%9 - idx) / -9
	rowNeighborIdx1, rowNeighborIdx2 := getNeighborsIdx(rowIdx)
	colIdx := idx % 9
	colNeighborIdx1, colNeighborIdx2 := getNeighborsIdx(colIdx)

	return &box{
		value:           value,
		rowIdx:          rowIdx,
		rowNeighborIdx1: rowNeighborIdx1,
		rowNeighborIdx2: rowNeighborIdx2,
		colIdx:          colIdx,
		colNeighborIdx1: colNeighborIdx1,
		colNeighborIdx2: colNeighborIdx2,
		squareIdx:       ((colIdx%3 - colIdx) / -3) + (((rowIdx%3 - rowIdx) / -3) * 3),
	}
}
