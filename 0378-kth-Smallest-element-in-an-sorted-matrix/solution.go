//nnn array slice 
func kthSmallest(matrix [][]int, k int) int {
	if len(matrix) <=0 {
		return 0
	}

	m, n := len(matrix), len(matrix[0])
	l, r := matrix[0][0], matrix[m-1][n-1]

	for l <= r {

		mid := l + (r-l)>>1  
		nt := 0 

		for i := 0; i < len(matrix); i++ {
			for j := 0; j < len(matrix[0]) && matrix[i][j] <= mid; j++ {
				nt++
			}
		}

		if nt < k {
			l = mid + 1
		} else {
			r = mid - 1
		}

	}

	return l
}

