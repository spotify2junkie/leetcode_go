// 
func subarraySum(nums []int, k int) int {
	rec := make(map[int]int, len(nums)) // 什么时候给index
	res, sum := 0, 0
	rec[0] = 1
	for i := range nums {
		sum += nums[i] // 前i个数字的和
		res += rec[sum-k] //就说明i之前存在j，使得a[j:i]元素和为k
		rec[sum]++
	}
	return res
}

