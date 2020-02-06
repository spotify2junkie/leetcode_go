//nnn dp 
func coinChange(coins []int, amount int) int {
    dp := make([]int,amount+1)
    dp[0] = 0 
    for i:=1;i < len(dp);i++ {
        dp[i] = amount+100
    }
    for i := 1; i <= amount; i++ {
        for _,coin := range coins {
            if coin <= i {
                dp[i] = min(dp[i],dp[i-coin]+1)  // 自己本身（初始的自己（被设定）或者迭代的自己最小）或者 coin的最小值
            }
        }
    }
    if dp[amount] == amount+100 {
        return -1
    } else {
        return dp[amount]
    }
}

func min (x,y int) int {
    if x > y {
        return y
    } else { return x}
}
