//nnn dynamic programming 

func isMatch(s string, p string) bool {
    sSize, pSize := len(s), len(p)
    dp := make([][]bool, sSize+1)
    for i := range dp {
        dp[i] = make([]bool, pSize+1)
    }
    dp[0][0] = true 
    //nnn s[:1] can match with p[:1]
    for j:=1; j <= pSize; j++ {
        if p[j-1] == '*' {
            dp[0][j] = dp[0][j-1] //nnn depends on p[j-1]
        }
    }

    for i:=1;i <= sSize;i++ {
        for j:=1;j<= pSize; j++ {
            if p[j-1] != '*' {
                dp[i][j] = dp[i-1][j-1] &&
					(p[j-1] == s[i-1] || p[j-1] == '?') //nnn ?dont know why p[j-1]
            } else {
                dp[i][j] = dp[i-1][j] || dp[i][j-1] 
            }
        }
    }
    return dp[sSize][pSize]
}
