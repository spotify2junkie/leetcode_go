func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s))
	for i := 0; i < len(s); i++ {
		for _, key := range wordDict {
			if i < len(key)-1 {
				continue
			}
			if string(s[i-len(key)+1:i+1]) == key && ((i == len(key)-1) || (dp[i-len(key)])) { ///多打括号
				dp[i] = true
			}
		}
	}
	return dp[len(s)-1]
}

