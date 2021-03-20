// merge sort

func sortArray(nums []int) []int {
	size := len(nums)
	if size <= 1 {
		return nums
	}
	mid := size / 2
	left := sortArray(nums[:mid])
	right := sortArray(nums[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	tmp := []int{}
	for len(left) > 0 || len(right) > 0 {
		if len(right) == 0 {
			return append(tmp, left...)
		}
		if len(left) == 0 {
			return append(tmp, right...)
		}
		if left[0] > right[0] {
			tmp = append(tmp, right[0])
			right = right[1:]
		} else { // 为啥不判断 直接用else
			tmp = append(tmp, left[0])
			left = left[1:]
		}
	}
	return tmp
}