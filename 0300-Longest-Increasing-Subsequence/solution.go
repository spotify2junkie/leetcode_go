func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	// for _,i := range dp {
	//     dp[i] = 1
	// }
	// base case：dp 数组全都初始化为 1
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j])
			}
		}
		dp[i] = dp[i] + 1
	}

	res := 0
	for i := 0; i < len(dp); i++ {
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
} 