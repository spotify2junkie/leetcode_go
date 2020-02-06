//nnn math 
func mySqrt(x int) int {
    res := x 
    for res*res > x {
        res = (res+x/res) / 2 //nnn newton method 
    }
    return res 
}
