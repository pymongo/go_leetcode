package go_leetcode

func transpose_solution_1(A [][]int) [][]int {
	m, n := len(A), len(A[0])
	res := make([][]int, 0, n)
	for j := 0; j < n; j++ {
		row := make([]int, 0, m)
		for i := 0; i < m; i++ {
			row = append(row, A[i][j])
		}
		res = append(res, row)
	}
	return res
}

func transpose_solution_2(A [][]int) [][]int {
	m, n := len(A), len(A[0])
	res := make([][]int, len(A[0]))
	for i := range res {
		res[i] = make([]int, m)
	}
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			res[col][row] = A[row][col]
		}
	}
	return res
}
