// dp  dynamic programming
func uniquePaths(m int, n int) int {
    //nnn robot move steps 
    dp := make([][]int,n) //nnn n for indexing 
    for i:=0;i<n;i++{
        dp[i] = make([]int,m)
    }
    for i:=0;i<m;i++{
        dp[0][i] = 1
    }
    for i:=0;i<n;i++{
        dp[i][0] = 1
    }
    for i:=1;i<n;i++{
        for j:=1;j<m;j++{
            dp[i][j]=dp[i][j-1]+dp[i-1][j]
        }
    }
    return dp[m-1][n-1] //nnn 理解为在地板上而不是角落
}
