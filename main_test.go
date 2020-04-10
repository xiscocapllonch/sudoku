package main

import (
	"testing"
)

func TestCreateSkInitBoxes(t *testing.T) {
	sk := createSk("000000000035070840097302510003904100060000090009503700051608920026090450000000000")

	boxesIdx := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 80}

	for idx, boxIdx := range boxesIdx {

		expValue := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 5, 0, 7, 0, 8, 4, 0, 0, 9, 7, 0}

		if sk.boxes[boxIdx].value != expValue[idx] {
			t.Errorf("Expected box value=%v, got %v", expValue[idx], sk.boxes[boxIdx].value)
		}

		expRowIdx := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 8}

		if sk.boxes[boxIdx].rowIdx != expRowIdx[idx] {
			t.Errorf("Expected box rowIdx=%v, got %v", expRowIdx[idx], sk.boxes[boxIdx].rowIdx)
		}

		expRowNeighborIdx1 := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 6}

		if sk.boxes[boxIdx].rowNeighborIdx1 != expRowNeighborIdx1[idx] {
			t.Errorf("Expected box rowNeighborIdx1=%v, got %v", expRowNeighborIdx1[idx], sk.boxes[boxIdx].rowNeighborIdx1)
		}

		expRowNeighborIdx2 := []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 1, 1, 7}

		if sk.boxes[boxIdx].rowNeighborIdx2 != expRowNeighborIdx2[idx] {
			t.Errorf("Expected box rowNeighborIdx2=%v, got %v", expRowNeighborIdx2[idx], sk.boxes[boxIdx].rowNeighborIdx2)
		}

		expColIdx := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 0, 1, 2, 3, 4, 5, 6, 7, 8, 0, 1, 2, 8}

		if sk.boxes[boxIdx].colIdx != expColIdx[idx] {
			t.Errorf("Expected box colIdx=%v, got %v", expColIdx[idx], sk.boxes[boxIdx].colIdx)
		}

		expColNeighborIdx1 := []int{1, 0, 0, 4, 3, 3, 7, 6, 6, 1, 0, 0, 4, 3, 3, 7, 6, 6, 1, 0, 0, 6}

		if sk.boxes[boxIdx].colNeighborIdx1 != expColNeighborIdx1[idx] {
			t.Errorf("Expected box colNeighborIdx1=%v, got %v", expColNeighborIdx1[idx], sk.boxes[boxIdx].colNeighborIdx1)
		}
		expColNeighborIdx2 := []int{2, 2, 1, 5, 5, 4, 8, 8, 7, 2, 2, 1, 5, 5, 4, 8, 8, 7, 2, 2, 1, 7}

		if sk.boxes[boxIdx].colNeighborIdx2 != expColNeighborIdx2[idx] {
			t.Errorf("Expected box colNeighborIdx2=%v, got %v", expColNeighborIdx2[idx], sk.boxes[boxIdx].colNeighborIdx2)
		}

		expSquareIdx := []int{0, 0, 0, 1, 1, 1, 2, 2, 2, 0, 0, 0, 1, 1, 1, 2, 2, 2, 0, 0, 0, 8}

		if sk.boxes[boxIdx].squareIdx != expSquareIdx[idx] {
			t.Errorf("Expected  box squareIdx=%v, got %v", expSquareIdx[idx], sk.boxes[boxIdx].squareIdx)
		}
	}
}

func TestCreateSkInitOptions(t *testing.T) {
	sk := createSk("000000000035070840097302510003904100060000090009503700051608920026090450000000000")

	if !contains(sk.options.rows[6], 7) ||
		!contains(sk.options.rows[6], 4) ||
		!contains(sk.options.rows[6], 3) {
		t.Errorf(
			"Expected row options contains %v, got false. Options are: %v",
			"7 and 4 and 3",
			sk.options.rows[7],
		)
	}

	if len(sk.options.rows[6]) != 3 {
		t.Errorf(
			"Expected row options length equal to %v, got %v",
			3,
			len(sk.options.rows[7]),
		)
	}

	if !contains(sk.options.cols[1], 8) ||
		!contains(sk.options.cols[1], 7) ||
		!contains(sk.options.cols[1], 4) ||
		!contains(sk.options.cols[1], 1) {
		t.Errorf(
			"Expected columns options contains %v, got false. Options are: %v",
			"7 and 4 and 1",
			sk.options.cols[1],
		)
	}

	if len(sk.options.cols[1]) != 4 {
		t.Errorf(
			"Expected columns options length equal to %v, got %v",
			4,
			len(sk.options.cols[1]),
		)
	}

	if !contains(sk.options.squares[6], 3) ||
		!contains(sk.options.squares[6], 4) ||
		!contains(sk.options.squares[6], 7) ||
		!contains(sk.options.squares[6], 8) ||
		!contains(sk.options.squares[6], 9) {
		t.Errorf(
			"Expected squares options contains %v, got false. Options are: %v",
			"3, 4, 7, 8 and 9",
			sk.options.squares[6],
		)
	}

	if len(sk.options.squares[6]) != 5 {
		t.Errorf(
			"Expected squares options length equal to %v, got %v",
			5,
			len(sk.options.squares[6]),
		)
	}
}

func TestGetBoxValues(t *testing.T) {
	sk := createSk("000000000035070840097302510003904100060000090009503700051608920026090450000000000")

	for hasNewValue := true; hasNewValue == true; {
		hasNewValue = sk.getBoxValues()
	}

	index := []int{12, 26, 66, 68, 72, 77, 4, 41, 42, 43}
	expectValue := []int{1, 6, 7, 1, 9, 5, 5, 7, 0, 9}

	for idx, boxIdx := range index {
		if sk.boxes[boxIdx].value != expectValue[idx] {
			t.Errorf("Expected box value=%v, got %v", expectValue[idx], sk.boxes[boxIdx].value)
		}
	}

	if sk.boxes[78].value != 6 {
		t.Errorf("Expected box value=%v, got %v", 6, sk.boxes[78].value)
	}
}
