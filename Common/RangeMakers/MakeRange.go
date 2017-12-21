package RangeMakers

func MakeRange(min, max int) []int {
	a := make([]int, max-min)
	for i :=range a{
		a[i] = min + i
	}
	return a
}

func MakeMatrix(rowSize int, columnSize int) [][]int {
	a := make([][]int, rowSize)
	for iCol := range a {
		a[iCol] = MakeRange(0, columnSize)
	}
	return a
}
