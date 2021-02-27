
func minFallingPathSum(matrix [][]int) int {
	res := 99999
	memos := make([][]int, len(matrix))
	for i := 0; i < len(memos); i++ {
		memos[i] = make([]int, len(matrix)) // 真的好麻烦 这个create dp
		for j := 0; j < len(memos[i]); j++ {
			memos[i][j] = -999999
		}
	}

	var helper func([][]int, int, int) int
	helper = func(matrix [][]int, i, j int) int {
		if i < 0 || j < 0 || j >= len(matrix[0]) || i >= len(matrix[0]) {
			return 99999
		}
		if memos[i][j] != -999999 {
			return memos[i][j]
		}
		if i == 0 {
			memos[i][j] = matrix[i][j] // base case
			return memos[i][j]
		}
		memos[i][j] = matrix[i][j] + minThree(helper(matrix, i-1, j-1),
			helper(matrix, i-1, j),
			helper(matrix, i-1, j+1))
		return memos[i][j]
	}

	for j := 0; j < len(matrix); j++ {
		res = min(res, helper(matrix, len(matrix)-1, j))
	}
	return res
}

func minThree(a, b, c int) int {
	return min(min(a, b), c)
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}