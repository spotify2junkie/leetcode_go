// backtrack
import "sort"

func combinationSum(c []int, t int) [][]int {
	res := [][]int{}
	if len(c) == 0 {
		return res
	}
	p := []int{}
	sort.Ints(c)
	genPermute(c, t, p, 0, &res)
	return res
}

func genPermute(c []int, t int, p []int, index int, res *[][]int) {
	if t <= 0 {
		if t < 0 {
			return
		}
		temp := make([]int, len(p))
		copy(temp, p)
		*res = append(*res, temp)
		return
	}

	for i := index; i < len(c); i++ {
		if c[i] > t { // 这里可以剪枝优化
			break
		}
		p = append(p, c[i])
		genPermute(c, t-c[i], p, i, res)
		p = p[:len(p)-1]
	}
	return
}

