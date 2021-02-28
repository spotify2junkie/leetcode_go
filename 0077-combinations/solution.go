
func combine(n int, k int) [][]int {
	res := [][]int{}
	choice := []int{}
	for i := 1; i <= n; i++ {
		choice = append(choice, i)
	}
	backtrack([]int{}, &res, 0, choice, k)
	return res
}

func backtrack(track []int, res *[][]int, loc int, choice []int, k int) {
	if len(track) == k {
		tmp := make([]int, k)
		copy(tmp, track)
		*res = append(*res, tmp)
		return
	}

	for i := loc; i < len(choice); i++ {
		track = append(track, choice[i])
		backtrack(track, res, i+1, choice, k)
		track = track[:len(track)-1]
	}
}