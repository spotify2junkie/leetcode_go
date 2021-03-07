/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	l, r := 0, 0
	count := make(map[byte]int)
	for r < len(s) {
		c := r 
		count[c]++
		right++
		if _, ok := count[c]; ok >= 1 {
			d := s[left]
			count[d]--
			left++
		}
		res = max(res, right-left)
	}
	return res
}

func max(a, b int) int {
	if a > b {return a}
	else {return b}
}
// @lc code=end

