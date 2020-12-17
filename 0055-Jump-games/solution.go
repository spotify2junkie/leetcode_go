func canJump(nums []int) bool {
	for i := len(nums) - 2; i >= 0; i-- { //nnn start from 倒数第二个
		// find 0
		if nums[i] != 0 {
			continue
		}
		j := i - 1 //nnn j 是等于0的前面一个index
		for ; j >= 0; j-- {
			if i-j < nums[j] {
				// 在 j 号位置上，可以跨过 0 元素
				i = j
				break
			}
		}

		if j == -1 {
			// 在 0 元素之前，没有位置可以跨过 0
			return false
		}
	}

	return true
}