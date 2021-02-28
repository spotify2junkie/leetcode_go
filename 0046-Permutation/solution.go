// backtracking
func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	res := [][]int{}
	used := make([]bool, len(nums))
	p := []int{}
	genPermute(nums, p, &used, &res)
	return res
}

func genPermute(nums []int, p []int, used *[]bool, res *[][]int) {
	if len(p) == len(nums) {
		temp := make([]int, len(p))
		copy(temp, p)
		*res = append(*res, temp)
		return // 这个避免了return 其他东西
	}

	for i := 0; i < len(nums); i++ {
		if !(*used)[i] {
			(*used)[i] = true // 因为排列可以乱序，就需要记录是否使用了这个元素
			p = append(p, nums[i])
			genPermute(nums, p, used, res)
			(*used)[i] = false
			p = p[:len(p)-1]
		}
	}
	return
}