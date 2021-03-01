func longestPalindrome(s string) string {
	res := ""
	tmp := ""
	for i := 0; i < len(s); i++ {
		if len(pa(s, i, i)) > len(pa(s, i, i+1)) {
			tmp = pa(s, i, i)
		} else {
			tmp = pa(s, i, i+1)
		}
		if len(tmp) > len(res) {
			res = tmp
		}
	}
	return res
}

func pa(s string, l int, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}