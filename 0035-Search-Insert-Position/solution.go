//nnn binart search
//nnn need neighbours 
func searchInsert(nums []int, target int) int {
    ln := len(nums)
    if ln == 0 {
        return 0
    }
    start, end := 0, len(nums)-1
    for start+ 1 < end {
        mid := (start+end) >> 1
        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            start = mid
        } else {
            end = mid
        }
    }
    if nums[start] >= target {
        return start
    } else if nums[end] >= target {
        return end
    } else if nums[end] < target { // 目标值比所有值都大
        return end + 1
    }
    return 0
}
