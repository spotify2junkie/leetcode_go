// backtrack 
func combinationSum(candidates []int, target int) [][]int {
    //nnn backtrack 可以实现
    result := make([][]int,0)
    backtrack([]int{},candidates,target,0,&result)
    return result
}

func backtrack(subset, candidates []int, target,sum int, res *[][]int) {
    if sum == target {
        *res = append(*res,subset)
        return 
    } else if sum > target {
        return  //nnn you only break look so here return 
    }
    for i := range candidates { //nnn  this 结构 
        subsetCopy := make([]int, 0, len(subset)+1)
        subsetCopy = append(subsetCopy,subset...)
        backtrack(append(subsetCopy,candidates[i]),candidates[i:],target,sum+candidates[i],res) 
    }
}


