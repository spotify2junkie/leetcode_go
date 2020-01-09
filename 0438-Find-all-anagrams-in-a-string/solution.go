// hash table, slice
package main

import "fmt"

func findAnagrams(s string, p string) []int {
	var freq [256]int // 至多256
	result := []int{}
	if len(s) == 0 || len(s) < len(p) {
		return result
	}
	for _, i := range p {
		freq[i-'a']++
	}
	left, right, count := 0, 0, len(p)
	for right < len(s) {
		if freq[s[right]-'a'] >= 1 {
			count--
		}
		freq[s[right]-'a']--
		right++ // end of this sweep
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

func main() {
	fmt.Println(findAnagrams("cbaebabacd",
		"abc"))
}
