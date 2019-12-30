// dynamic programming 
// palindromic substring 
func countSubstrings(s string) int {
    count := 0 
    for i:=0;i < len(s); i++ {
        count +=  palindrome_counter(s,i,i) // 回文函数
        count += palindrome_counter(s,i,i+1) // 无非就是自己和多一位的其他string
    }
    return count
}


func palindrome_counter(s string,left,right int) int {
    res := 0
    for left >= 0 && right < len(s) && s[left] == s[right] {
        left--
        right++ 
        res++ 
    }
    return res 
}