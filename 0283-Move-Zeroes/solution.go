// two pointers 
func moveZeroes(nums []int) {
    l := len(nums)
    i,j := 0,0 
    for i < l {  // 先把非0位置填上 
        if nums[i] != 0 {
            nums[j] = nums[i]
            j++
        }
        i++
    }

    for j < i {
        nums[j] = 0
        j++
    }
}
