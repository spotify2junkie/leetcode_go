import (
    "sort"
)

func combinationSum2(candidates []int, target int) [][]int {
    if len(candidates) == 0 {
        return [][]int{}
    }
    path, res := []int{}, [][]int{}
    sort.Ints(candidates)
    backtrack(candidates, target, path, &res, 0)
    return res
}


func backtrack(candidates []int, target int, path []int, res *[][]int, idx int) {
    if target == 0 {
        tmp := make([]int, len(path))
        copy(tmp, path)
        *res = append(*res, tmp)
        return 
    }
    for i:= idx; i < len(candidates); i++ { //nnn idx 是外部结果
        if i > idx && candidates[i] == candidates[i-1] {
            continue
        }
        if target >= candidates[i] {
            path = append(path, candidates[i])
            backtrack(candidates, target-candidates[i], path, res, i+1)  //nnn this is index 
            path = path[:len(path)-1]
        }
    }
    return 
}


