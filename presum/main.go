package presum

type NumMatrix struct {
	sums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	rows := len(matrix)
	cols := len(matrix[0])

	sums := make([][]int, rows+1)
	sums[0] = make([]int, cols+1)

	for i := 1; i <= rows; i++ {
		sums[i] = make([]int, cols+1)
		for j := 1; j <= cols; j++ {
			sums[i][j] = sums[i-1][j] + sums[i][j-1] - sums[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	return NumMatrix{sums: sums}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	row1++
	col1++
	row2++
	col2++

	return this.sums[row2][col2] - this.sums[row1-1][col2] - this.sums[row2][col1-1] + this.sums[row1-1][col1-1]
}
