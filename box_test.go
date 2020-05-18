package sudoku

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestNewBox(t *testing.T) {
	cmpOptions := cmp.AllowUnexported(box{})

	expectedBox1 := box{2, 0, 1, 2, 7, 6, 8, 2, nil}

	b1 := *newBox(7, 2)

	if diff := cmp.Diff(expectedBox1, b1, cmpOptions); diff != "" {
		t.Error("Resulting box1 does not match the expected box", diff)
	}

	expectedBox2 := box{0, 7, 6, 8, 8, 6, 7, 8, nil}

	b2 := *newBox(71, 0)

	if diff := cmp.Diff(expectedBox2, b2, cmpOptions); diff != "" {
		t.Error("Resulting box2 does not match the expected box", diff)
	}

	expectedBox3 := box{9, 8, 6, 7, 0, 1, 2, 6, nil}

	b3 := *newBox(72, 9)

	if diff := cmp.Diff(expectedBox3, b3, cmpOptions); diff != "" {
		t.Error("Resulting box3 does not match the expected box", diff)
	}

}
