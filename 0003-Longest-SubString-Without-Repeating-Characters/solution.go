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

// another solution 
// // another solution
// func lengthOfLongestSubstring(s string) int {
//     if len(s) == 0 {
//         return 0
//     }
//     var freq [256]int
//     left, right,result := 0,-1,0
//     for left < len(s) {
//         if right+1 < len(s) && freq[s[right+1]-'a'] == 0 {
//             freq[s[right+1]-'a']++
//             right++
//         } else {
//             freq[s[left]-'a']-- // 经典
//             left++
//         }
//         result = max(result, right-left+1)
//     }
//     return result
// }

// func max(i, j int) int {
//     if i > j {
//         return i 
//     } else {
//         return j 
//     }
// }