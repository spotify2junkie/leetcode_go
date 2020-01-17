// bit manipulation 
func hammingWeight(num uint32) int {
    count := 0
    for num != 0 {
        count++
        num &= num-1 //nnn and 'bit'
        
    }
    return count
}