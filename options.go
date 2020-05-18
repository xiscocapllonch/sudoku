package sudoku

import (
	"sort"
)

type options struct {
	rows    map[int][]int
	cols    map[int][]int
	squares map[int][]int
}

func newOptions() *options {
	rows := make(map[int][]int)
	cols := make(map[int][]int)
	squares := make(map[int][]int)

	for i := 0; i < 9; i++ {
		rows[i] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		cols[i] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		squares[i] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	}

	return &options{
		rows:    rows,
		cols:    cols,
		squares: squares,
	}
}

func (o *options) removeOption(value, rowIdx, colIdx, squareIdx int) {
	rm := func(s []int, rV int) []int {
		for i, v := range s {
			if v == rV {
				s = append(s[:i], s[i+1:]...)
				break
			}
		}

		return s
	}

	if value != 0 {
		o.rows[rowIdx] = rm(o.rows[rowIdx], value)
		o.cols[colIdx] = rm(o.cols[colIdx], value)
		o.squares[squareIdx] = rm(o.squares[squareIdx], value)
	}
}

func (o *options) getPointOptions(rowIdx, colIdx, squareIdx int) []int {

	m := make(map[int]uint8)

	for _, k := range o.rows[rowIdx] {
		m[k] |= 1 << 0
	}

	for _, k := range o.cols[colIdx] {
		m[k] |= 1 << 1
	}

	for _, k := range o.squares[squareIdx] {
		m[k] |= 1 << 2
	}

	var result []int

	for k, v := range m {
		if v&(1<<0) != 0 && v&(1<<1) != 0 && v&(1<<2) != 0 {
			result = append(result, k)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	return result
}

func (o *options) getNeighborOptions(rowNIdx1, rowNIdx2, colNIdx1, colNIdx2 int, pointOpts []int) []int {
	var neighborOpts []int

	for _, opts := range [][]int{
		o.rows[rowNIdx1],
		o.rows[rowNIdx2],
		o.cols[colNIdx1],
		o.cols[colNIdx2],
	} {
		for _, opt := range opts {
			neighborOpts = appendUnique(neighborOpts, opt)
		}
	}

	var options []int

	for _, v := range pointOpts {
		if !contains(neighborOpts, v) {
			options = append(options, v)
		}
	}

	return options
}
