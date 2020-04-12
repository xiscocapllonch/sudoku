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

func (s *sk) solveTrivial() (candidates []string) {
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
			return func() []string {
				for i := 2; i < 9; i++ {
					for idx, box := range s.boxes {
						if box.value == 0 {
							if len(box.options) == i {
								for _, op := range box.options {
									output := ""
									for i2, box := range s.boxes {
										if i2 == idx {
											output = output + strconv.Itoa(op)
										} else {
											output = output + strconv.Itoa(box.value)
										}
									}
									candidates = append(candidates, output)
								}
								return candidates
							}
						}
					}
				}
				return []string{}
			}()
		}
	}

	output := ""
	for _, box := range s.boxes {
		output = output + strconv.Itoa(box.value)
	}

	candidates = append(candidates, output)

	return candidates
}

func (s *sk) setNewBoxValue(value int, idx int) {
	(*s).boxes[idx].value = value
	box := (*s).boxes[idx]
	(*s).options.removeOptions(value, box.rowIdx, box.colIdx, box.squareIdx)
}

func (b *box) initIndexes (idx int) {
	getNeighborsIdx := func(idx int) (idx1 int, idx2 int) {
		if idx%3 == 0 {
			return idx + 1, idx + 2
		} else if idx%3 == 1 {
			return idx - 1, idx + 1
		} else {
			return idx - 2, idx - 1
		}
	}

	rowIdx := func() int {
		return (idx%9 - idx) / -9
	}()

	colIdx := func() int {
		return idx % 9
	}()

	b.rowIdx = rowIdx
	b.colIdx = colIdx
	b.rowNeighborIdx1, b.rowNeighborIdx2 = getNeighborsIdx(rowIdx)
	b.colNeighborIdx1, b.colNeighborIdx2 = getNeighborsIdx(colIdx)
	b.squareIdx = ((colIdx%3 - colIdx) / -3) + (((rowIdx%3 - rowIdx) / -3) * 3)
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
