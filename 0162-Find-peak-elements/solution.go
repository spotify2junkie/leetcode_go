//nnn maximum lies in one direction
func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func findPeakElement(nums []int) int {
	n := len(nums)
	l, r := 0, n
	for l < r { // 这个地方不能是n 因为会溢出
		mid := (l + r) / 2
		if nums[mid] < nums[mid+1] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}