func searchRange(nums []int, target int) []int {
    if len(nums) == 0 {
        return []int{-1, -1}
    }
    result := make([]int, 2)
    start, end := 0, len(nums)-1
    for start+ 1 < end {
        mid := (start+end) >> 1
        if nums[mid] >= target {
            end = mid
        } else {
            start = mid
        }
    }
    if nums[start] == target {
        result[0] = start
    } else if nums[end] == target {
        result[0] = end
    } else {
        result[0], result[1] = -1, -1
        return result
    }
    start, end = 0, len(nums)-1
    for start+ 1 < end {
        mid := (start+end) >> 1
        if nums[mid] <= target {
            start = mid
        } else {
            end = mid 
        }
    }
    if nums[end] == target {
        result[1] = end
    } else if nums[start] == target {
        result[1] = start
    } else {
        result[0], result[1] = -1, -1
        return result
    }
    return result 
}


