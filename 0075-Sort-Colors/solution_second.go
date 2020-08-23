func sortColors(nums []int)  {
    if len(nums) == 0 {
        return 
    }
    r,w,b := 0,0,0 // mark the end of color 
    for _, num := range nums {
        if num == 0 {
            nums[b] = 2
            b++
            nums[w] = 1
            w++
            nums[r] = 0
            r++
        } else if num == 1 {
            nums[b] = 2 
            b++
            nums[w] = 1
            w++
        } else if num == 2 {
            nums[b] = 2
            b++ 
        }
    }
}
