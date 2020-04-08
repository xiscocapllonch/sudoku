package main

import (
	"testing"
)

func TestCreateSk(t *testing.T) {
	sk := createSk("000000000035070840097302510003904100060000090009503700051608920026090450000000000")

	boxesIdx := []int{0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,80}


	for idx, boxIdx := range boxesIdx {

		expValue := []int{0,0,0,0,0,0,0,0,0,0,3,5,0,7,0,8,4,0,0,9,7,0}

		if  sk.boxes[boxIdx].value != expValue[idx] {
			t.Errorf("Expected box value=%v, got %v", expValue[idx], sk.boxes[boxIdx].value)
		}

		expRowIdx := []int{0,0,0,0,0,0,0,0,0,1,1,1,1,1,1,1,1,1,2,2,2,8}

		if  sk.boxes[boxIdx].rowIdx != expRowIdx[idx] {
			t.Errorf("Expected box rowIdx=%v, got %v", expRowIdx[idx], sk.boxes[boxIdx].rowIdx)
		}

		expColIdx := []int{0,1,2,3,4,5,6,7,8,0,1,2,3,4,5,6,7,8,0,1,2,8}

		if  sk.boxes[boxIdx].colIdx != expColIdx[idx] {
			t.Errorf("Expected box colIdx=%v, got %v", expColIdx[idx], sk.boxes[boxIdx].colIdx)
		}

		expSquareIdx := []int{0,0,0,1,1,1,2,2,2,0,0,0,1,1,1,2,2,2,0,0,0,8}

		if  sk.boxes[boxIdx].squareIdx != expSquareIdx[idx] {
			t.Errorf("Expected  box squareIdx=%v, got %v", expSquareIdx[idx], sk.boxes[boxIdx].squareIdx)
		}
	}
}
