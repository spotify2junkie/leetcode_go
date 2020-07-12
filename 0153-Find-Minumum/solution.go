//nnn dont quite get 


func findMin(nums []int) int {
    start, end := 0, len(nums)-1
    for start + 1 < end {
        mid := (start+end) >> 1
        if nums[mid] <= nums[end] {
            end = mid
        } else {
            start = mid
        }
    }
    if nums[start] > nums[end] {
        return nums[end]
    }
    return nums[start]
}
