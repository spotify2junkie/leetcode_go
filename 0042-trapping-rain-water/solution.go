// two pointers

func trap(height []int) int {
	if len(height) < 1 {
		return 0
	}
	n := len(height)
	left, right := 0, n-1
	res := 0
	l_max, r_max := height[0], height[n-1]
	for left <= right {
		l_max = max(l_max, height[left])
		r_max = max(r_max, height[right])

		if l_max < r_max {
			res += l_max - height[left]
			left++
		} else {
			res += r_max - height[right]
			right--
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}