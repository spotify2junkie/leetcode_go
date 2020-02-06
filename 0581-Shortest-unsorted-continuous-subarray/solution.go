// slice , two pointers 
func findUnsortedSubarray(nums []int) int {
    n := len(nums)
    left,right := 0,-1
    min,max := nums[n-1],nums[0] //nnn reverse
    for i:=1;i<n;i++ {
        if max <= nums[i] {
            max = nums[i]
        } else {
            right = i
        }

        j := n-i-1 //nnn one for but two pointers 
        if min >= nums[j] {
            min = nums[j]
        } else {
            left = j 
        }
    }
    return right-left+1 //nnn length way 
}
