//nnn binary search solution 
func twoSum(nums []int, target int) []int {
    start, end := 0 , len(nums) -1
    res := []int{}
    for start < end {
        if nums[start] + nums[end] == target {
            res = append(res, start+1)
            res = append(res, end + 1)
            return res //nnn miss return before 
        } else if nums[start] + nums[end] > target {
            end--
        } else {
            start++
        }
    }
    return res 
}
