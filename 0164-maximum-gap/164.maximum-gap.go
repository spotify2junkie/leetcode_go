/*
 * @lc app=leetcode id=164 lang=golang
 *
 * [164] Maximum Gap
 */

// @lc code=start
func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	var quickSort func(i, j int)
	quickSort = func(i, j int) {
		ci, cj := i, j
		p := （i+j) >>1
		for i < j {
			for nums[i] >= nums[p] && i < j { // 为啥必须需要这个顺序
				i++
			}
			for nums[j] <= nums[p] && i < j {
				j--
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
		nums[i], nums[p] = nums[p], nums[i] // 很关键, 交换次序
		p = i
		if ci < p {
			quickSort(ci, p)
		}
		if p+1 < cj {
			quickSort(p+1, cj) //location +1 就是右边
		}
	}
	quickSort(0, len(nums)-1)

	ret := 0
	for i := 1; i < len(nums); i++ {
		tmp := nums[i] - nums[i-1]
		if ret < tmp {
			ret = tmp
		}
	}

	return ret
}

// @lc code=end

