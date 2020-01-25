func plusOne(digits []int) []int {
    i := len(digits)-1
    for { //nnn new for loop format 
        if i < 0 {
            return append([]int{1},digits...) //nnn three points 
        }
        if digits[i] != 9 {
            digits[i]++
            return digits
        } else {
            digits[i] = 0 
            i--  
        }
    }
    return digits
}