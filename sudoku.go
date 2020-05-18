package sudoku

import (
	"strconv"
	"strings"
	"sync"
)

type sudoku struct {
	parent  string
	boxes   [81]box
	options options
}

type TryResult struct {
	parent     string
	solution   string
	candidates []string
}

func SolveSudoku(input string) string {
	const NumReceivers = 10

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	solve := func (input string, c chan string) TryResult {
		result := TryToSolve(input)
		if result.solution != "" {
			close(c)
		} else {
			for _, candidate := range result.candidates {
				c <- candidate
			}
		}

		return result
	}

	skChan := make(chan string)

	var solution string

	go func(s string) {
		result := solve(s, skChan)
		if result.solution != "" {
			solution = result.solution
		}
	}(input)

	for i := 0; i < NumReceivers; i++ {
		go func() {
			defer wgReceivers.Done()
			for s := range skChan {
				go func(s string) {
					result := solve(s, skChan)
					if result.solution != "" {
						solution = result.solution
					}
				}(s)
			}
		}()
	}

	wgReceivers.Wait()

	return solution
}

func FormatPrintSudoku(input string) string {
	result := ""

	for idx, value := range strings.Split(input, "") {
		if (idx % 9)%3 == 0 {
			result += "   |"
		}

		if value == "0" {
			result += " |"
		} else {
			result += value + "|"
		}

		if idx % 9 == 8 {
			result += "\n"
			if ((idx%9 - idx) / -9)%3 == 2 {
				result += "\n"
			}
		}
	}

	return result
}

func TryToSolve(input string) TryResult {
	sudoku := newSudoku(input)
	return sudoku.getCandidates()
}

func newSudoku(input string) *sudoku {
	var boxes [81]box
	options := *newOptions()

	for idx, value := range strings.Split(input, "") {
		intV, _ := strconv.Atoi(value)
		b := *newBox(idx, intV)
		boxes[idx] = b
		options.removeOption(b.value, b.rowIdx, b.colIdx, b.squareIdx)
	}

	return &sudoku{
		parent:  input,
		boxes:   boxes,
		options: options,
	}
}

func (s *sudoku) getCandidates() TryResult {
	for idx, box := range s.boxes {
		if s.boxes[idx].value == 0 {
			var options []int
			pointOpts := s.options.getPointOptions(box.rowIdx, box.colIdx, box.squareIdx)

			if len(pointOpts) == 1 {
				options = pointOpts
			}

			neighborOpts := s.options.getNeighborOptions(
				box.rowNeighborIdx1,
				box.rowNeighborIdx2,
				box.colNeighborIdx1,
				box.colNeighborIdx2,
				pointOpts,
			)

			if len(neighborOpts) > 0 {
				options = neighborOpts
			} else {
				options = pointOpts
			}

			if len(options) == 1 {
				s.boxes[idx].value = options[0]
				box := s.boxes[idx]
				s.options.removeOption(options[0], box.rowIdx, box.colIdx, box.squareIdx)
				s.getCandidates()
			} else {
				s.boxes[idx].options = options
			}
		}
	}

	for _, box := range s.boxes {
		if box.value == 0 {
			return TryResult{
				parent:     s.parent,
				candidates: s.findCandidates(),
				solution:   "",
			}
		}
	}

	solvedSk := ""
	for _, box := range s.boxes {
		solvedSk = solvedSk + strconv.Itoa(box.value)
	}
	return TryResult{
		parent:   s.parent,
		solution: solvedSk,
	}
}

func (s *sudoku) findCandidates() (candidates []string) {
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
