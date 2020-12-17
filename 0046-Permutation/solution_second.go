func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	used, p, res := make([]bool, len(nums)), []int{}, [][]int{}
	generatePermutation(nums, 0, p, &res, &used)
	return res
}

func generatePermutation(nums []int, index int, p []int, res *[][]int, used *[]bool) {
	if index == len(nums) { // 记录结果
		temp := make([]int, len(p))
		copy(temp, p)
		fmt.Println(temp)
		*res = append(*res, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !(*used)[i] {
			(*used)[i] = true // 进入选择
			p = append(p, nums[i])
			generatePermutation(nums, index+1, p, res, used) // 进入了 // return 了之后要改掉used
			//p = p[:len(p)-1] // 撤销了nums
			(*used)[i] = false
		}
	}
	return
}