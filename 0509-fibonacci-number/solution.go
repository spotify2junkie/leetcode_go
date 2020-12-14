// recursion
// memos 备忘录

func fib(N int) int {
	memos := make(map[int]int, N+1)
	var helper func(int) int
	helper = func(N int) int {
		if N == 0 {
			return 0
		}
		if N == 1 {
			return 1
		}
		if _, ok := memos[N]; ok {
			return memos[N]
		}
		memos[N] = helper(N-1) + helper(N-2)
		return memos[N]
	}
	return helper(N)
}

