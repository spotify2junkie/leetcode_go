func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return max(robRange(nums, 0, len(nums)-2),
		robRange(nums, 1, len(nums)-1))
}

func robRange(nums []int, s, e int) int {
	dpFirst, dpSecond := 0, 0
	dpCur := 0
	for i := e; i >= s; i-- {
		dpCur = max(dpFirst, dpSecond+nums[i])
		dpSecond = dpFirst
		dpFirst = dpCur
	}
	return dpCur
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
