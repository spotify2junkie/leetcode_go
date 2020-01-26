func longestCommonPrefix(strs []string) (res string) { //nnn return key with bracket 
    shortest := -1
    for _, str := range strs {
        if len(str) < shortest || shortest == -1 {
            shortest = len(str)
        }
    }
    for i:=0; i < shortest; i++ {
        for _,str := range strs {
            if str[i] != strs[0][i] {
                return 
            }
        }
        res += string(strs[0][i])
    }
    return res    
}