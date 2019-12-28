// dynamic programming , DFS 
func numSquares(n int) int {
    if n >= 0 && n < 2 {
        return n 
    }
    dp := make([]int,n+1)
    perfects := []int{}
    // gather all squares smaller than n 
    for i:=1 ; i*i <= n;i++ {
        perfects = append(perfects,i*i)
    }

    for i:=1;i < n+1;i++ {
        dp[i] = i 
        for _,s :=range perfects {
            if i-s >= 0 {
                dp[i] = min(dp[i],dp[i-s]+1) // either itself or 12-9 = 3  , dp[3] + 1
            }
        }
    }
    return dp[n]
}

func min(x,y int) int {
    if x < y {
        return x
    } else { return y }
}