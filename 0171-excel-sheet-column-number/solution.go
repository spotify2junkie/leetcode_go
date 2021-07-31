func titleToNumber(columnTitle string) int {
    all := 0 
    for _,key := range columnTitle {
        cur := int(key - 'A') + 1 
        all  = all * 26 + cur 
    }
    return all 
}

