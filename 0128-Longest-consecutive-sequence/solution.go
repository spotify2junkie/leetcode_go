//nnn slice , array 
func longestConsecutive(nums []int) int {
	m := map[int]int{}
	res := 0

	for _, n := range nums {
		_, ok := m[n]
		if !ok {
			left := m[n-1]
			right := m[n+1]
			m[n] = 1 + left + right
			if left > 0 {
				m[n-left] = m[n]
			}
			if right > 0 {
				m[n+right] = m[n]
			}
			if m[n] > res {
				res = m[n]
			}
		}
	}

	return res
}
