func lengthOfLongestSubstring(s string) int {
    location := [256]int{} // 只有256长是因为，假定输入的字符串只有ASCII字符
	for i := range location {
		location[i] = -1 // 先设置所有的字符都没有见过
	}
    
    left, maxlens := 0,0
    for i:=0;i < len(s); i++ {
        if location[s[i]] >= left {
            left = i + 1  // 缩小左边界
        } else if i+1-left > maxlens {  // 查看右边界
            maxlens = i+1-left 
        }
        location[s[i]] = i
    }
    return maxlens
}