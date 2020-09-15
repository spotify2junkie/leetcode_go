// binary search

func findDuplicate(nums []int) int {
    left, right := 0, len(nums)-1
    for left <= right { // look for bounday
        mid,count := (left+right)>>1, 0
        for _, v := range nums {
            if v <= mid {
                count++
            }
        }
        if count > mid {
            right = mid-1
        } else {
            left = mid + 1 
        }
    }
    return left
}
