package main

import (
	"fmt"
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
	boxes [81]box
}

func main() {
	sk := createSk("000000000035070840097302510003904100060000090009503700051608920026090450000000000")
	fmt.Println("sk:", sk.boxes[10])

}

func createSk(input string) sk {
	var boxes [81]box

	for idx, value := range strings.Split(input, "") {
		v, _ := strconv.Atoi(value)
		boxes[idx] = box{
			value: v,
			rowIdx: func () int {
				return (idx%9-idx)/-9
			}(),
			colIdx: func () int {
				return idx%9
			}(),
			squareIdx: func () int {
				col := idx%9
				row := (idx%9-idx)/-9
				return ((col%3-col)/-3) + (((row%3-row)/-3)*3)
			}(),
		}
	}

	return sk{
		boxes: boxes,
	}
}
