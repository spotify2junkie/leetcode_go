// dp
func longestCommonSubsequence(s1 string, s2 string) int {
	sl1, sl2 := len(s1), len(s2)
	dp := make([][]int, sl1+1)
	for i := range dp {
		dp[i] = make([]int, sl2+1)
	}
	for i := 1; i <= sl1; i++ {
		for j := 1; j <= sl2; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[sl1][sl2]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}