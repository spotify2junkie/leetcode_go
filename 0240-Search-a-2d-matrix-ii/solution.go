// divide and conquer , matrix

func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}
    for _,slice := range matrix {
        if slice[len(slice)-1] >= target {
            if judge_slice(slice, target) {
                return true 
            }
        }
    }
    return false
}



func judge_slice(nums []int, target int) bool {
    i:=0
    for i < len(nums) {
        if nums[i] == target {
            return true 
        } else if nums[i] < target {
            i++
        } else {
            return false 
        }
    }
    return false 
}

