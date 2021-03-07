
func subsets(nums []int) [][]int {
	res := [][]int{}
	backtrack(nums, 0, []int{}, &res)
	return res
}

func backtrack(nums []int, start int, track []int, res *[][]int) {
	tmp := make([]int, len(track))
	copy(tmp, track) // why always need copy ？
	*res = append(*res, tmp)
	for i := start; i < len(nums); i++ {
		track = append(track, nums[i])
		backtrack(nums, i+1, track, res)
		track = track[:len(track)-1]
	}
}