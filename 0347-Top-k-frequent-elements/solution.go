// hash table and sort 
import "sort"

func topKFrequent(nums []int, k int) []int {
    // 两次循环可以解决
    res := []int{}
    rec := make(map[int]int,len(nums)) // 为了取indecx 
    for _,n := range nums {
        rec[n]++
    }
    counts := []int{}
    for _,c := range rec {
        counts = append(counts,c)
    }
    sort.Ints(counts)
    // 需要一个boundary
    min := counts[len(counts)-k]
    for n,c := range rec {
        if c >= min {
            res = append(res,n)
        }
    }
    return res 
}