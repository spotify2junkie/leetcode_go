// mergeSort

func sortArray(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	mid := n / 2
	l := sortArray(nums[:mid])
	r := sortArray(nums[mid:])
	res := merge(l, r)
	return res
}

func merge(l, r []int) []int {
	ret := make([]int, 0, len(r)+len(l))
	for len(r) > 0 || len(l) > 0 {
		if len(r) == 0 {
			return append(ret, l...)
		}
		if len(l) == 0 {
			return append(ret, r...)
		}
		if l[0] <= r[0] {
			ret = append(ret, l[0])
			l = l[1:]
		} else {
			ret = append(ret, r[0])
			r = r[1:]
		}
	}
	return ret
}