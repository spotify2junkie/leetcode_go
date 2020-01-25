func lengthOfLongestSubstring(s string) int {
    if len(s) == 0 {
        return 0
    }
    var freq [256]int
    left, right,result := 0,0,0
    for left < len(s) {
        if right < len(s) && freq[s[right]-'a'] == 0 {
            freq[s[right]-'a']++
            right++
        } else {
            freq[s[left]-'a']-- // 经典
            left++
        }
        result = max(result, right-left)
    }
    return result
}

func max(i, j int) int {
    if i > j {
        return i 
    } else {
        return j 
    }
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