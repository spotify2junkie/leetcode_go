func quick(nums []int, start, end int) {
    if start < end {
        j, i := end, start + 1
        
        for i <= j {
            if nums[i] <= nums[start] {
                i++
            } else if nums[j] >= nums[start] {
                j--
            } else {
                nums[j], nums[i] = nums[i], nums[j]
                j--
                i++
            }
        }
        
        nums[j], nums[start] = nums[start], nums[j]
        
        quick(nums, start, j - 1)
        quick(nums, j + 1, end)
    }
}

func sortArray(nums []int) []int {
    quick(nums, 0, len(nums) - 1)
    return nums
} 