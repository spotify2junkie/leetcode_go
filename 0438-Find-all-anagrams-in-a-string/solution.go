// sliding window hash table, slice
func findAnagrams(s string, p string) []int {
    var freq[256]int
    result := []int{} 
    if len(s)== 0 || len(s) < len(p) {
		return result
	}
    left,right,count := 0,0,len(p)
    for i:=0;i < count; i++ {
        freq[p[i]-'a']++
    }
    for right < len(s) {
        if freq[s[right]-'a'] >= 1 {
            count--
        }
        freq[s[right]-'a']--
        right++
        if count == 0 {
            result = append(result, left)
        }
        if right-left == len(p) {
            if freq[s[left]-'a'] >= 0 {
                count++
            }
            freq[s[left]-'a']++
            left++
        }
    }
    return result
}

## another solution 
## func findAnagrams(s string, p string) []int {
##     sl, pl := len(s), len(p)
##     ans := []int{}
##     freq1,freq2 := [256]int{}, [256]int{}
##     for i:=0;i < pl; i++ {
##         freq2[p[i]-'a']++
##     }
##     for i:=0;i < sl; i++ {
##         if i >=  pl {
##             freq1[s[i-pl]-'a']--
##         }
##         freq1[s[i]-'a']++
##         if freq1 == freq2 {
##             ans = append(ans,i+1-pl)
##         }
##     }
##     return ans 
## }
