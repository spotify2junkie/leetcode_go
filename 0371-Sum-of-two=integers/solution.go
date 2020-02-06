//nnn bit manipulation 
func getSum(a, b int) int {
    for a != 0 {
        a,b = (a&b)<<1, a^b //nnn 进位与不进位 
    }
    return b 
}
