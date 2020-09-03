// slice 
// map 
// sort 



func rearrangeBarcodes(A []int) []int {
    n := len(A)
    if n <= 2 {
        return A
    }
    count := [10001]int{}  // 指定切片大小
    for _, v := range A {
        count[v]++
    }
    sort.Slice(A, func(i, j int) bool {
        if count[A[i]] == count[A[j]] { // 一定要的条件
            return A[i] > A[j]
        }
        return count[A[i]] > count[A[j]]
    }) // 注意函数这里结尾 
    i := 0 // index 在外面设置好
    res := make([]int, n) // make 指定slice assign 
    for _, v := range A {
        res[i] =v
        i+=2
        if i >= n {
            i =1
        }
    }
    return res 
}
