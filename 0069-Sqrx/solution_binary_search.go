func mySqrt(x int) int {
    if x == 0 || x == 1 {
        return x
    }
    lo, high, res := 1, x, -1
    for lo <= high {
        m := (lo + high) >>1
        if m*m == x {
            return m
        }
        if m * m < x {
            lo = m +1
            res = m
        } else {
            high = m-1
        }
    }
    return res
}
