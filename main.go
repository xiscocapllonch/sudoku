package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

type options struct {
	rows    map[int][]int
	cols    map[int][]int
	squares map[int][]int
}

type sk struct {
	boxes   [81]box
	options options
}

func main() {
	skChan := make(chan string)
	sudoku := "000000000035070840097302510003904100060000090009503700051608920026090450000000000"
	// sudoku := "400070002080040050003209800009000500860000013005000200006804300030060020700020009"

	go solve(sudoku, skChan)

	for s := range skChan {
		go func(s string) {
			solve(s, skChan)
		}(s)
	}
}

func solve(input string, c chan string) {
	s := createSk(input)
	candidates := s.solveTrivial()
	if len(candidates) == 1 {
		s.print()
		os.Exit(1)
		return
	} else if len(candidates) > 1 {
		for _, candidate := range candidates {
			c <- candidate
		}
	}
}

func createSk(input string) sk {
	var boxesInput [81]int

	for idx, value := range strings.Split(input, "") {
		intV, _ := strconv.Atoi(value)
		boxesInput[idx] = intV
	}

	newSk := sk{}

	newSk.initBoxes(boxesInput)
	newSk.initOptions()

	return newSk
}

func (s *sk) print() {
	for _, box := range s.boxes {
		if box.colIdx%3 == 0 {
			fmt.Print("   |")
		}

		if box.value == 0 {
			fmt.Print(" |")
		} else {
			fmt.Print(box.value, "|")
		}

		if box.colIdx == 8 {
			fmt.Print("\n")
			if box.rowIdx%3 == 2 {
				fmt.Println("")
			}
		}
	}
}

func (s *sk) initBoxes(input [81]int) {
	var boxes [81]box

	for idx, value := range input {
		newBox := box{value: value}
		newBox.initIndexes(idx)
		boxes[idx] = newBox
	}

	(*s).boxes = boxes
}

func (s *sk) initOptions() {
	s.options.setDefaultValues()
	for _, box := range s.boxes {
		if box.value != 0 {
			s.options.removeOptions(box.value, box.rowIdx, box.colIdx, box.squareIdx)
		}
	}
}

func (s *sk) solveTrivial() []string {
	for idx, box := range s.boxes {
		if s.boxes[idx].value == 0 {
			options := box.getOptions(s.options)
			if len(options) == 1 {
				s.boxes[idx].value = options[0]
				box := s.boxes[idx]
				s.options.removeOptions(options[0], box.rowIdx, box.colIdx, box.squareIdx)
				s.solveTrivial()
			} else {
				s.boxes[idx].options = options
			}
		}
	}

	for _, box := range s.boxes {
		if box.value == 0 {
			return s.getCandidates()
		}
	}

	solvedSk := ""
	for _, box := range s.boxes {
		solvedSk = solvedSk + strconv.Itoa(box.value)
	}
	return []string{solvedSk}
}

func (s *sk) getCandidates() (candidates []string) {
	opts, idx := func() ([]int, int) {
		for i := 2; i < 9; i++ {
			for idx, box := range s.boxes {
				if box.value == 0 {
					if len(box.options) == i {
						return box.options, idx
					}
				}
			}
		}
		return []int{}, -1
	}()

	for _, op := range opts {
		output := ""
		for i, box := range s.boxes {
			var value int
			if i == idx {
				value = op
			} else {
				value = box.value
			}
			output = output + strconv.Itoa(value)
		}
		candidates = append(candidates, output)
	}

	return candidates
}

func (b *box) initIndexes(idx int) {
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

	b.rowIdx = (idx%9 - idx) / -9
	b.colIdx  = idx % 9
	b.squareIdx = ((b.colIdx%3 - b.colIdx) / -3) + (((b.rowIdx%3 - b.rowIdx) / -3) * 3)
	b.rowNeighborIdx1, b.rowNeighborIdx2 = getNeighborsIdx(b.rowIdx)
	b.colNeighborIdx1, b.colNeighborIdx2 = getNeighborsIdx(b.colIdx)
}

func (b *box) getOptions(o options) []int {
	var pointOpts []int
	var neighborOpts []int

	for i := 1; i < 10; i++ {
		if contains(o.rows[b.rowIdx], i) &&
			contains(o.cols[b.colIdx], i) &&
			contains(o.squares[b.squareIdx], i) {
			pointOpts = append(pointOpts, i)
		}

		if contains(o.rows[b.rowNeighborIdx1], i) ||
			contains(o.rows[b.rowNeighborIdx2], i) ||
			contains(o.cols[b.colNeighborIdx1], i) ||
			contains(o.cols[b.colNeighborIdx2], i) {
			neighborOpts = append(neighborOpts, i)
		}
	}

	if len(pointOpts) == 1 {
		return pointOpts
	}

	var options []int

	for _, v := range pointOpts {
		if !contains(neighborOpts, v) {
			options = append(options, v)
		}
	}

	if len(options) > 0 {
		return options
	} else {
		return pointOpts
	}
}

func (o *options) setDefaultValues() {
	rowsOp := make(map[int][]int)
	colsOp := make(map[int][]int)
	squaresOp := make(map[int][]int)

	for i := 0; i < 9; i++ {
		rowsOp[i] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		colsOp[i] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		squaresOp[i] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	}

	(*o).rows = rowsOp
	(*o).cols = colsOp
	(*o).squares = squaresOp
}

func (o *options) removeOptions(value int, rowIdx int, colIdx int, squareIdx int) {
	rm := func(s []int, rV int) []int {
		for i, v := range s {
			if v == rV {
				s = append(s[:i], s[i+1:]...)
				break
			}
		}

		return s
	}
	o.rows[rowIdx] = rm(o.rows[rowIdx], value)
	o.cols[colIdx] = rm((*o).cols[colIdx], value)
	o.squares[squareIdx] = rm(o.squares[squareIdx], value)
}

func contains(s []int, cV int) bool {
	for _, v := range s {
		if v == cV {
			return true
		}
	}
	return false
}
