// binary search 
// 先找左边界
// 再找右边界
func searchRange(nums []int, target int) []int {
    l := searchfirst(nums,target)
    r := searchLastEqualElement(nums, target)
    range1 := []int{l,r}
    return range1
}

func searchfirst(nums []int, target int) int {
    low, high := 0, len(nums)-1
    for low <= high {
        mid := low + ((high-low)>>1)
        if nums[mid] > target {
            high = mid-1
        } else if nums[mid] < target {
            low = mid+1
        } else {  // 这里开始 不同
            if (mid == 0) || (nums[mid-1] != target) {
                return mid
            }
            high = mid-1 // 反向移动到前面
        }
    }
    return -1
}

// 二分查找最后一个与 target 相等的元素，时间复杂度 O(logn)
func searchLastEqualElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			if (mid == len(nums)-1) || (nums[mid+1] != target) { // 找到最后一个与 target 相等的元素
				return mid
			}
			low = mid + 1 //// 正向移动到后面
		}
	}
	return -1
}