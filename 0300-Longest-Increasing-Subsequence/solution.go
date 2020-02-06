// dynamic programming 
func lengthOfLIS(nums []int) int {
    if len(nums) == 1 {
        return 1
    }
    if len(nums) == 0 {
        return 0 
    }
    if len(nums) == 2 {
        if max(nums[0],nums[1]) == nums[0] {
            return 1
        } else {
            return 2
        }
    }
    dp := make([]int,len(nums)+1)
    dp[0] = 0 
    res := 0 
    for i:=1;i <= len(nums); i++ {
        for j:=1; j < i;j++ {
            if nums[j-1] < nums[i-1] {
                dp[i] = max(dp[i],dp[j])
            }
        }
        dp[i] = dp[i]+1 
        res = max(res,dp[i]) // get longest among all , probably not the last one 
    }
    return res // 注意一下
}


func max(i, j int) int {
    if i > j {
        return i 
    } else {
        return j 
    }
}
