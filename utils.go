package sudoku

func contains(s []int, i int) bool {
	for _, v := range s {
		if v == i {
			return true
		}
	}
	return false
}

func appendUnique(s []int, i int) []int {
	for _, v := range s {
		if v == i {
			return s
		}
	}
	return append(s, i)
}
