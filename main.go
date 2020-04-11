package main

import (
	"fmt"
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

type sk struct {
	boxes   [81]box
	options struct {
		rows    map[int][]int
		cols    map[int][]int
		squares map[int][]int
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
		boxes[idx] = box{
			value: value,
			rowIdx: func() int {
				return (idx%9 - idx) / -9
			}(),
			colIdx: func() int {
				return idx % 9
			}(),
			squareIdx: func() int {
				col := idx % 9
				row := (idx%9 - idx) / -9
				return ((col%3 - col) / -3) + (((row%3 - row) / -3) * 3)
			}(),
		}

		boxes[idx].rowNeighborIdx1, boxes[idx].rowNeighborIdx2 = getNeighborsIdx(boxes[idx].rowIdx)
		boxes[idx].colNeighborIdx1, boxes[idx].colNeighborIdx2 = getNeighborsIdx(boxes[idx].colIdx)

	}

	(*s).boxes = boxes
}

func getNeighborsIdx(idx int) (idx1 int, idx2 int) {
	if idx%3 == 0 {
		return idx + 1, idx + 2
	} else if idx%3 == 1 {
		return idx - 1, idx + 1
	} else {
		return idx - 2, idx - 1
	}
}

func (s *sk) initOptions() {
	rowsOp := make(map[int][]int)
	colsOp := make(map[int][]int)
	squaresOp := make(map[int][]int)

	for i := 0; i < 9; i++ {
		rowsOp[i] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		colsOp[i] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		squaresOp[i] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	}

	for _, box := range s.boxes {
		if box.value != 0 {
			rowsOp[box.rowIdx] = remove(rowsOp[box.rowIdx], box.value)
			colsOp[box.colIdx] = remove(colsOp[box.colIdx], box.value)
			squaresOp[box.squareIdx] = remove(squaresOp[box.squareIdx], box.value)
		}
	}

	(*s).options.rows = rowsOp
	(*s).options.cols = colsOp
	(*s).options.squares = squaresOp
}

func (s *sk) solveTrivial() (bool, []int) {
	for idx, box := range s.boxes {
		if s.boxes[idx].value == 0 {
			var pointOpts []int
			var neighborOpts []int

			for i := 1; i < 10; i++ {
				if contains(s.options.rows[box.rowIdx], i) &&
					contains(s.options.cols[box.colIdx], i) &&
					contains(s.options.squares[box.squareIdx], i) {
					pointOpts = append(pointOpts, i)
				}

				if contains(s.options.rows[box.rowNeighborIdx1], i) ||
					contains(s.options.rows[box.rowNeighborIdx2], i) ||
					contains(s.options.cols[box.colNeighborIdx1], i) ||
					contains(s.options.cols[box.colNeighborIdx2], i) {
					neighborOpts = append(neighborOpts, i)
				}
			}

			if len(pointOpts) == 1 {
				(*s).setNewBoxValue(pointOpts[0], idx)
				(*s).solveTrivial()
			}

			var options []int

			for _, v := range pointOpts {
				if !contains(neighborOpts, v) {
					options = append(options, v)
				}
			}

			if len(options) == 1 {
				(*s).setNewBoxValue(options[0], idx)
				(*s).solveTrivial()
			} else if len(options) != 0 {
				(*s).boxes[idx].options = options
			} else {
				(*s).boxes[idx].options = pointOpts
			}
		}
	}

	for _, box := range s.boxes {
		if box.value == 0 {
			return false, func() []int {
				for i := 1; i < 9; i++ {
					for _, box := range s.boxes {
						if box.value == 0 {
							if len(box.options) == i {
								return box.options
							}
						}
					}
				}
				return []int{}
			}()
		}
	}

	return true, []int{}
}

func (s *sk) setNewBoxValue(value int, idx int) {
	(*s).boxes[idx].value = value

	box := (*s).boxes[idx]

	(*s).options.rows[box.rowIdx] = remove((*s).options.rows[box.rowIdx], value)
	(*s).options.cols[box.colIdx] = remove((*s).options.cols[box.colIdx], value)
	(*s).options.squares[box.squareIdx] = remove((*s).options.squares[box.squareIdx], value)
}

func remove(s []int, rV int) []int {
	for i, v := range s {
		if v == rV {
			s = append(s[:i], s[i+1:]...)
			break
		}
	}

	return s
}

func contains(s []int, cV int) bool {
	for _, v := range s {
		if v == cV {
			return true
		}
	}
	return false
}
