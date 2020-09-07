// adjust 2 sums 
func twoSum(A []int, target int) []int {
    i, j := 0, len(A)-1
    for i < j {
        if A[i]+A[j] == target {
            return []int{i+1, j+1}
        } else if A[i]+A[j] < target {
            i++
        } else {
            j--
        }
    }
    return []int{-1, -1}
}
