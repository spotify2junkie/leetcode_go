//nnn bit manipulation
func singleNumber(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result = result ^ nums[i] // bit manipulation
	}
	return result
}

