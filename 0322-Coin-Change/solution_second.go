// 备忘录
func coinChange(coins []int, amount int) int {
	memos := make(map[int]int, amount+1)
	var dp func(int) int
	dp = func(n int) int {
		if _, ok := memos[n]; ok {
			return memos[n]
		}
		if n == 0 {
			return 0
		}
		if n < 0 {
			return -1
		}
		res := 1<<31 - 1
		for _, element := range coins {
			sub := dp(n - element)
			if sub == -1 {
				continue
			}
			res = min(res, sub+1)
		}
		if res != 1<<31-1 {
			memos[n] = res
		} else {
			memos[n] = -1
		}
		return memos[n]

	}
	return dp(amount)
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}