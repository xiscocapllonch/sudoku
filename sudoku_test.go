package sudoku

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestSolveSudoku(t *testing.T) {
	// LEVEL HARD, SOLUTION WITH EXPLORATION
	solvedSK := SolveSudoku("000000000035070840097302510003904100060000090009503700051608920026090450000000000")
	expSolution := "618459237235176849497382516873924165562817394149563782751648923326791458984235671"
	if solvedSK != expSolution {
		t.Errorf(
			"Result solution should be: %v, but got: %v",
			expSolution,
			solvedSK,
		)
	}

	// LEVEL EASY, SOLUTION WITHOUT EXPLORATION
	solvedSK = SolveSudoku("030006040980204000060809257000700090000000801000053406490008600058107020203600009")
	expSolution = "732516948985274163164839257346781592579462831821953476497328615658197324213645789"

	if solvedSK != expSolution {
		t.Errorf(
			"Result solution should be: %v, but got: %v",
			expSolution,
			solvedSK,
		)
	}

}

func TestFormatPrintSudoku(t *testing.T) {
	sudokuToPrint := FormatPrintSudoku("618459237235176849497382516873924165562817394149563782751648923326791458984235671")
	expectedFormat := "" +
		"   |6|1|8|   |4|5|9|   |2|3|7|\n" +
		"   |2|3|5|   |1|7|6|   |8|4|9|\n" +
		"   |4|9|7|   |3|8|2|   |5|1|6|\n\n" +
		"   |8|7|3|   |9|2|4|   |1|6|5|\n" +
		"   |5|6|2|   |8|1|7|   |3|9|4|\n" +
		"   |1|4|9|   |5|6|3|   |7|8|2|\n\n" +
		"   |7|5|1|   |6|4|8|   |9|2|3|\n" +
		"   |3|2|6|   |7|9|1|   |4|5|8|\n" +
		"   |9|8|4|   |2|3|5|   |6|7|1|\n\n"

	if diff := cmp.Diff(expectedFormat, sudokuToPrint); diff != "" {
		t.Error("Resulting SudokuToPrint not match the expected format", diff)
	}
}

func TestTryToSolve(t *testing.T) {
	// LEVEL HARD, SOLUTION WITH EXPLORATION
	result := TryToSolve("000000000035070840097302510003904100060000090009503700051608920026090450000000000")

	if len(result.candidates) != 2 {
		t.Errorf("Sudoku length candidates should be 2 but got %v", len(result.candidates))
	}

	if string(result.candidates[0][3]) == string(result.candidates[1][3]) {
		t.Errorf("candidates should be diferent for this string idx")
	}

	if (string(result.candidates[0][3]) == "3" || string(result.candidates[0][3]) == "4") &&
		(string(result.candidates[1][3]) == "3" || string(result.candidates[1][3]) == "4") {
		t.Errorf("candidates should be 3 or 4 for this string idx")
	}
	result = TryToSolve(result.candidates[0])
	result = TryToSolve(result.candidates[0])
	result = TryToSolve(result.candidates[1])
	result = TryToSolve(result.candidates[0])

	if len(result.candidates) != 0 {
		t.Errorf("Sudoku length candidates should be 0 but got %v", len(result.candidates))
	}

	solution := "618459237235176849497382516873924165562817394149563782751648923326791458984235671"

	if result.solution != solution {
		t.Errorf(
			"Result solution shoul be: %v, but got: %v",
			solution,
			result.solution,
		)
	}

	// LEVEL EASY, SOLUTION WITHOUT EXPLORATION
	result = TryToSolve("030006040980204000060809257000700090000000801000053406490008600058107020203600009")

	solution = "732516948985274163164839257346781592579462831821953476497328615658197324213645789"

	if result.solution != solution {
		t.Errorf(
			"Result solution shoul be: %v, but got: %v",
			solution,
			result.solution,
		)
	}
}

func TestNewSudoku(t *testing.T) {
	input := "000000000035070840097302510003904100060000090009503700051608920026090450000000000"

	s := newSudoku(input)

	if diff := cmp.Diff([]int{3, 4, 7}, s.options.rows[6]); diff != "" {
		t.Error("Resulting op rows 6 not match the expected slice", diff)
	}

	if diff := cmp.Diff([]int{1, 4, 7, 8}, s.options.cols[1]); diff != "" {
		t.Error("Resulting op cols 1 not match the expected slice", diff)
	}

	if diff := cmp.Diff([]int{3, 4, 7, 8, 9}, s.options.squares[6]); diff != "" {
		t.Error("Resulting op squares 6 not match the expected slice", diff)
	}

	boxesIdx := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 80}

	for idx, boxIdx := range boxesIdx {

		expValue := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 5, 0, 7, 0, 8, 4, 0, 0, 9, 7, 0}

		if s.boxes[boxIdx].value != expValue[idx] {
			t.Errorf("Expected box value=%v, got %v", expValue[idx], s.boxes[boxIdx].value)
		}

		expRowIdx := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 8}

		if s.boxes[boxIdx].rowIdx != expRowIdx[idx] {
			t.Errorf("Expected box rowIdx=%v, got %v", expRowIdx[idx], s.boxes[boxIdx].rowIdx)
		}

		expRowNeighborIdx1 := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 6}

		if s.boxes[boxIdx].rowNeighborIdx1 != expRowNeighborIdx1[idx] {
			t.Errorf("Expected box rowNeighborIdx1=%v, got %v", expRowNeighborIdx1[idx], s.boxes[boxIdx].rowNeighborIdx1)
		}

		expRowNeighborIdx2 := []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 1, 1, 7}

		if s.boxes[boxIdx].rowNeighborIdx2 != expRowNeighborIdx2[idx] {
			t.Errorf("Expected box rowNeighborIdx2=%v, got %v", expRowNeighborIdx2[idx], s.boxes[boxIdx].rowNeighborIdx2)
		}

		expColIdx := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 0, 1, 2, 3, 4, 5, 6, 7, 8, 0, 1, 2, 8}

		if s.boxes[boxIdx].colIdx != expColIdx[idx] {
			t.Errorf("Expected box colIdx=%v, got %v", expColIdx[idx], s.boxes[boxIdx].colIdx)
		}

		expColNeighborIdx1 := []int{1, 0, 0, 4, 3, 3, 7, 6, 6, 1, 0, 0, 4, 3, 3, 7, 6, 6, 1, 0, 0, 6}

		if s.boxes[boxIdx].colNeighborIdx1 != expColNeighborIdx1[idx] {
			t.Errorf("Expected box colNeighborIdx1=%v, got %v", expColNeighborIdx1[idx], s.boxes[boxIdx].colNeighborIdx1)
		}
		expColNeighborIdx2 := []int{2, 2, 1, 5, 5, 4, 8, 8, 7, 2, 2, 1, 5, 5, 4, 8, 8, 7, 2, 2, 1, 7}

		if s.boxes[boxIdx].colNeighborIdx2 != expColNeighborIdx2[idx] {
			t.Errorf("Expected box colNeighborIdx2=%v, got %v", expColNeighborIdx2[idx], s.boxes[boxIdx].colNeighborIdx2)
		}

		expSquareIdx := []int{0, 0, 0, 1, 1, 1, 2, 2, 2, 0, 0, 0, 1, 1, 1, 2, 2, 2, 0, 0, 0, 8}

		if s.boxes[boxIdx].squareIdx != expSquareIdx[idx] {
			t.Errorf("Expected  box squareIdx=%v, got %v", expSquareIdx[idx], s.boxes[boxIdx].squareIdx)
		}
	}
}
