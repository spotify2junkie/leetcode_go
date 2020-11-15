func maxSlidingWindow(nums []int, k int) []int {
	// left , right partition
	size := len(nums)
	if k <= 1 {
		return nums
	}
	p := k - 1
	left := make([]int, size)
	right := make([]int, size)
	for i := 0; i < size; i++ {
		if i%p == 0 {
			left[i] = nums[i]
		} else {
			left[i] = max(nums[i], left[i-1])
		}
	}
	right[size-1] = nums[size-1]
	for i := size - 2; i >= 0; i-- {
		if (i+1)%p == 0 {
			right[i] = nums[i]
		} else {
			right[i] = max(nums[i], right[i+1])
		}
	}
	max_size := make([]int, size-k+1)
	for i := 0; i <= size-k; i++ {
		max_size[i] = max(right[i], left[i+p])
	}
	return max_size
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
