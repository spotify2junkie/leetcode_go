// back tracking
package main

import "fmt"

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	res := make([][]int, 0)
	backtrack(nums, []int{}, &res)
	return res
}

func backtrack(nums []int, subres []int, res *[][]int) {
	// subres keep track of prev result
	//nnn input explanation
	/*
	   - nums: array can be used
	   - subrescopy: array to be append on
	   - res: resuult to be updated
	*/
	if len(nums) == 0 {
		*res = append(*res, subres)
		return
	}

	for i := 0; i < len(nums); i++ {
		numsCopy := append([]int{}, nums...)
		subresCopy := append([]int{}, subres...)
		backtrack(append(numsCopy[:i], numsCopy[i+1:]...), append(subresCopy, nums[i]), res) //local reference
	}
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
