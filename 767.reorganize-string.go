/*
 * @lc app=leetcode id=767 lang=golang
 *
 * [767] Reorganize String
 */

// @lc code=start
func reorganizeString(S string) string {
	count := make([][]int, 26)
	maxCount := 0
	for _, b := range S {
		count[b-'a'][0]++
		count[b-'a'][1] = int(b - a)
		maxCount = max(maxCount, count[b-'a'][0])
	}
	if maxCount*2 > n+1 {
		return ""
	}

}

// @lc code=end

