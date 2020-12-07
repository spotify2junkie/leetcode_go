func lengthOfLongestSubstring(s string) int {
	l, r := 0, 0
	count := make(map[byte]int)
	res := 0
	for r < len(s) {
		c := s[r]
		count[c]++
		r++
		for count[c] > 1 { //模板
			d := s[l]
			count[d]--
			l++
		}
		res = max(res, r-l)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}