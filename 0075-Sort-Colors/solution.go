// two pointers (actually three) 
// using two "swap" to help the third (middle color) be ranged 
// amazing
func sortColors(nums []int) {
	if len(nums) == 0 {
		return
	}

    redR,cur,blueL := 0,0,len(nums)-1
    for cur <= blueL {
        switch nums[cur] {
        case 0:
            nums[redR],nums[cur] = nums[cur],nums[redR]
            redR++
            cur++
        case 1:
            cur++
        case 2:
            nums[blueL],nums[cur] = nums[cur],nums[blueL]
            blueL--
        }
    }
}
