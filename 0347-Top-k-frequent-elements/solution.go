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





func canPartition(nums []int) bool {
    sum := 0 
    for i:=0;i < len(nums);i++ {
        sum +=nums[i]
    }
    sum = sum>>1
    dp := make([][]bool,len(nums)+1) // 2-d matrix
    for i := range dp {
        dp[i] = make([]bool,sum+1) // make for row slice
    }
    for i:=0;i < len(nums)+1;i++ { // len(nums)+1
        dp[i][0] = true
    }
    for j:=1; j < sum+1;j++ {  // do not do dp[0][0] 
        dp[0][j] = false
    }
    
    for i:=1; i < len(nums)+1;i++ {
        for j:=1;j<sum+1;j++ {
            if j >= nums[i-1] {
                dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
            }
        }
    }
    return dp[len(nums)][sum]
}
