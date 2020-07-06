func findDuplicate(nums []int) int {
    start, end := 0, len(nums) -1
    for start <= end {
        mid := (start + end) >> 1
        cnt := 0
        for _, val := range nums {
            if val <= mid {
                cnt++ 
            }
        }
        if cnt > mid {
            end = mid -1
        } else { start = mid + 1}
    }
    return start //nnn why start <= end  
}
