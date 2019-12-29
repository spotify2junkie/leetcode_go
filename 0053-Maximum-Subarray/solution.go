func maxSubArray(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    // 这里动态规划只要确保 dp > 0 
    dp := 0
    re := nums[0]
    for _,num := range nums {
        if dp > 0 {
            dp += num
        } else {
            dp = num
        }

        if dp > re {
            re = dp // keep track of current max
        }
    }
    return re     
}