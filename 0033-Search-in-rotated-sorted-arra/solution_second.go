func search(nums []int, target int) int {
	rotated := indexOfMin(nums)
	size := len(nums)
	start, end := 0, len(nums)-1
	for start <= end {
		mid := (start + end) >> 1
		originalMid := (mid + rotated) % size
		if nums[originalMid] < target {
			start = mid
		} else if nums[originalMid] > target {
			end = mid
		} else {
			return originalMid
		}
	}
	return -1

}

func indexOfMin(nums []int) int { // index of min
	size := len(nums)
	left, right := 0, size-1
	for left < right {
		mid := (left + right) / 2
		if nums[right] < nums[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}


