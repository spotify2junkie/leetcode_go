// sliding window
// dont quite get why use count
package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
    var need [128]int
    var have [128]int
    for i:= range s1 {
        need[s1[i]]++
    }
    s1l := len(s1)
    s2l := len(s2)
    for left,right,count := 0,0,0;right<s2l; {
        if have[s2[right]] < need[s2[right]] {
			// 出现了 window 中缺失的字母  针对只出现一次还比较好
			count++
		}
		have[s2[right]]++
		right++
		if count == s1l {
            return true 
        }
        if right >= s1l {
			left = right-s1l
			if have[s2[left]] == need[s2[left]] {
				// 出现了 window 中缺失的字母
				count--
			}
		have[s2[left]]--
		}

        
    }
    return false
}

func main() {
	fmt.Println(checkInclusion("ab","eidbaoo"))
	fmt.Println(checkInclusion("ab","eidboaoo"))
	fmt.Println(checkInclusion("hello","ooolleoooleh"))
}
