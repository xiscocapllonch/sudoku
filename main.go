package main

import (
	"strconv"
	"strings"
)

type box struct {
	value     int
	rowIdx    int
	colIdx    int
	squareIdx int
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
	}

	(*s).boxes = boxes
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
