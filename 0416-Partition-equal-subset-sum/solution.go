func canPartition(nums []int) bool {
    sum := 0 
    for _, n := range nums {
        sum += n 
    } 
    if sum&1 == 1 {
        return false
    }
    sum = sum >> 1
    n := len(nums)

    dp := make([][]bool,n+1) // n+1 for first []
    for i := range dp {
        dp[i] = make([]bool,sum+1) // sum+1 for 0 
    }

    /*
    define starting point
    dp[i][0]
    d[[j][0]
    */
    for i:=0;i < n+1;i++ {
        dp[i][0] = true
    }
    for j:=1; j < sum+1;j++ {  // do not do dp[0][0] 
        dp[0][j] = false
    }

    for i:=1; i < n+1;i++{
        for j:=1;j < sum+1;j++ {
            dp[i][j] = dp[i-1][j] // typo I did i-i
            if j >= nums[i-1] {
                dp[i][j] = dp[i][j] || dp[i-1][j-nums[i-1]]
            }
        }
    }
    return dp[n][sum]
}