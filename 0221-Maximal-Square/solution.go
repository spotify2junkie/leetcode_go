//nnn dp
func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	if n == 0 {
		return 0
	}
	e := 0
	res := make([][]int, m) // why use m
	for i := range res {
		res[i] = make([]int, n)
		if matrix[i][0] == '1' {
			res[i][0] = 1
			e = 1
		}
	}
	for j := 0; j < n; j++ {
		if matrix[0][j] == '1' {
			res[0][j] = 1
			e = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '1' {
				res[i][j] = 1 + min(
					res[i-1][j-1],
					min(
						res[i-1][j],
						res[i][j-1], // 画一下
					))
				e = max(e, res[i][j])
			}
		}
	}
	return e * e
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
