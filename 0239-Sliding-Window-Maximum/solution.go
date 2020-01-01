// sliding window 
func maxSlidingWindow(nums []int, k int) []int {
    // first create a function to 
    // get max of an array
    if len(nums) < k || k == 0 {
        return []int{}
    }
    
    max_array := []int{}
    for i:=0; i <= len(nums)-k; i++ {
        max_array = append(max_array,maxOfAnArray(nums[i:i+k]))
    }
    return max_array
}


func maxOfAnArray(slice []int) int {
    max := slice[0]
    for _,v := range slice {
        if v > max {
            max = v
        }
    }
    return max 
}