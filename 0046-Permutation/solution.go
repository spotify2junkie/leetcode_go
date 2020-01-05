// back tracking 

func permute(nums []int) [][]int {
    if len(nums) == 0 {
        return nil 
    }
    res := make([][]int,0)
    backtrack(nums,[]int{},&res)
    return res 
}

func backtrack(nums []int, subres []int, res *[][]int) { // subres keep track of prev result
    if len(nums) == 0 {
        *res = append(*res,subres)
        return 
    }

    for i:=0;i < len(nums);i++ {
        numsCopy := append([]int{}, nums...)
		subresCopy := append([]int{}, subres...)
        backtrack(append(numsCopy[:i],numsCopy[i+1:]...),append(subresCopy,nums[i]),res)
    }
}