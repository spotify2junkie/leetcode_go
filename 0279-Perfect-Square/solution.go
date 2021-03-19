// dp
// redo
func numSquares(n int) int {
	squares := []int{}
	for i := 1; i*i <= n; i++ {
		squares = append(squares, i*i)
	}
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i
		for _, element := range squares {
			if element <= i {
				dp[i] = min(dp[i], dp[i-element]+1)
			}
		}
	}
	return dp[n]
}

func min(i, j int) int {
	if i > j {
		return j
	} else {
		return i
	}
}