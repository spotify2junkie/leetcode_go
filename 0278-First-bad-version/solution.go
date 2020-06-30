//nnn binary search 
//nnn from left to right 
//nnn consider the start scenario 

func firstBadVersion(n int) int {
    start, end := 1, n
    for start + 1 < end {
        mid := (start+end)>>1
        if isBadVersion(mid) {
            end = mid
        } else if isBadVersion(mid) == false {
            start = mid
        }
    }
    if isBadVersion(start) {
        return start
    }
    return end
    
}
