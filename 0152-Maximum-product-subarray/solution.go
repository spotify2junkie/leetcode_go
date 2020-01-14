// array
package main

import "fmt"

func maxProduct(nums []int) int {
	//nnn using DP
	min, max, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			max, min = min, max
		}
		max = maxF(nums[i], max*nums[i])
		min = minF(nums[i], min*nums[i])
		res = maxF(res, max)
	}
	return res
}

func maxF(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minF(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxProduct([]int{2, 3, 0, -2, -3, -1}))
}
