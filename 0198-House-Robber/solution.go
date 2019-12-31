// dynamic programming 
// 暴力！
func rob(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        return nums[0]
    }
    if len(nums) == 2 {
        return max(nums[0],nums[1])
    }
    dp := make([]int,len(nums))
    dp[0] = nums[0]
    dp[1] = nums[1]
    maxv := max(dp[0],dp[1])
    for i:=2; i < len(nums); i++ {
        if i >= 3 {
            dp[i] = nums[i] + max(dp[i-2],dp[i-3])
        } else {
            dp[i] = nums[i] + dp[i-2]
        }
        maxv = max(dp[i],maxv)
    }
    return maxv
}

func max (i, j int) int {
    if i > j {
        return i 
    } else {
        return j
    }
}

// another solution 
// func rob(nums []int) int {
// 	// a 对于偶数位上的最大值的记录
// 	// b 对于奇数位上的最大值的记录
// 	var a, b int
// 	for i, v := range nums {
// 		if i%2 == 0 {
// 			a = max(a+v, b)
// 		} else {
// 			b = max(a, b+v)
// 		}
// 	}

// 	return max(a, b)
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
