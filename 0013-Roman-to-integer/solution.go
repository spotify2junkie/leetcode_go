func romanToInt(s string) int {
    res := 0
    m := map[byte]int{
        'I':1,
        'V':5,
        'X':10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }
    /*
    check digit before current value 
    if smaller then minus else plus 
    */
    last := 0
    for i:=len(s)-1;i>=0;i-- {
        cur := m[s[i]]
        sign := 1 
        if cur < last {
            sign = -1
        }
        res += sign*cur
        last = cur //mmm update
    }
    return res 
}