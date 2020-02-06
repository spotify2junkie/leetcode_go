// slice , two sweep
func productExceptSelf(nums []int) []int {
	//nnn how to finish in O(n)
	n := len(nums)
	res := make([]int,n)
	left := 1
	right := 1
	for i:=0;i<n;i++{
		res[i] = left
		left *= nums[i]
	}
	for i:=n-1;i>=0;i--{
		res[i] *= right //nnn amazing 
		right *= nums[i]
	}
	return res
 }
	
