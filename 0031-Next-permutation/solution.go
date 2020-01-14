//nnn array
func nextPermutation(nums []int)  {
    n := len(nums)
    for i := n - 1; i > 0; i-- {
        if nums[i] <= nums[i-1] {
			continue
		}
        j:=i  //nnn find 
        for ; j < n-1;j++ {
            if nums[j] > nums[i-1] && nums[j+1] <= nums[i-1] { //nnn good stuff 
                break
            }
        } 
        nums[i-1], nums[j] = nums[j], nums[i-1]
        reverse(nums[i:])
        return 
    }
        //nnn else condtion 
        // in case an arrangement is impossible
	reverse(nums)
}




func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}