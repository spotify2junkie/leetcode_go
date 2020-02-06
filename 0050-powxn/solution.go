//nnn math using recursion 
func myPow(x float64, n int) float64 {
    if x < 0 {
        return 1.0 / power(x,-n)
    }
    return pow(x,n)
}

func pow(x float64, n int) float64 {
    if x == 0 {
        return 0 
    }
    if n == 0 {
        return 1 
    }
    res := pow(x,n>>1)
    if n&1 == 0 {  //nnn 可以被整除
        return res*res
    }
    return res*res*x
}
