func guessNumber(n int) int {
    left, right := 1,n
    for left <= right {
        mid := (left+right) >> 1
        res := guess(mid)
        if res == 0 {
            return mid
        } else if res == -1 {
            right = mid -1
        } else {
            left = mid +1 
        }
    }
    return -1 //nnn add default cond  
}
