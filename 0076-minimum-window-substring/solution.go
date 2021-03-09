// sliding window
// two pointers
func minWindow(s2 string, s1 string) string {
	have := make([]int, 128)
	need := make([]int, 128)
	res := ""
	for _, element := range s1 {
		need[element]++
	}
	left, right, count := 0, 0, 0
	size := len(s2)
	min := 1<<31 - 1
	for right < size {
		rc := s2[right]
		right++
		if need[rc] > 0 {
			have[rc]++
			if need[rc] >= have[rc] {
				count++
			}
		}
		for count == len(s1) { //这个条件
			if right-left < min {
				min = right - left
				res = s2[left:right]
			}
			lc := s2[left]
			left++
			if need[lc] > 0 {
				// have[lc-'a']--
				if need[lc] >= have[lc] {
					count--
				}
				have[lc]--
			}
		}
	}
	return res
}
