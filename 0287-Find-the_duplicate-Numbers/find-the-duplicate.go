func findDuplicate(nums []int) int {
    counter := map[int]int{}
    res := 0
    for _,num := range nums {
        counter[num]++
    }
    for key := range counter {
        if counter[key] > 1 {
            res = key// 为啥需要赋值bunengreturn
        }
    }
    return res 
}