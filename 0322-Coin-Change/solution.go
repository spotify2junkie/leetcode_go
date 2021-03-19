//nnn dp
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = 1<<31 - 1
	}
	for i := 1; i <= amount; i++ {
		for _, element := range coins {
			if element <= i {
				dp[i] = min(dp[i], dp[i-element]+1)
			}
		}
	}
	if dp[amount] == 1<<31-1 {
		return -1
	} else {
		return dp[amount]
	}
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}