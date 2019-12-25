func twoSum(nums []int, target int) []int {
	maps := make(map[int]int) 
	result := make([]int, 0)
	for index, num := range nums {
		_, ok := maps[target - num]
		if ok {
			result = append(result, maps[target - num]) // find the second matched-input 
			result = append(result, index) // get 
			break 
		}
		maps[num] = index
	}
	return result
}