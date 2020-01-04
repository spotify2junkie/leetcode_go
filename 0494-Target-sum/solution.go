// dynamic programming , backpack 
func findTargetSumWays(nums []int, S int) int {
    // define sum
    sum := 0 
    for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
    if sum < S {
        return 0 
    }
    // sum and S 必须同为奇数或者偶数
    if (sum+S)%2==1 {
        return S 
    }

    return backpack(nums,(sum+S)/2)
}

func backpack(nums []int, target int) int {
    // dp[i]表示nums[i]和为target的个数
    dp := make([]int, target+1)
    dp[0] = 1
    for i:=0; i < len(nums);i++{
        for j:=target; j >= nums[i];j-- { // minus 
            dp[j] += dp[j-nums[i]] // this 
        }
    }
    return dp[target]
}