func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	res := []string{}
	backtrack(n, n, "", &res) // 取值
	return res                // 取局部被改变的res
}

func backtrack(lindex, rindex int, str string, res *[]string) { // *类型搞明白
	if lindex == 0 && rindex == 0 { //边界情况
		*res = append(*res, str)
		return
	}

	if lindex > 0 { // 边界情况
		backtrack(lindex-1, rindex, str+"(", res)
	}

	if rindex > 0 && rindex > lindex { // 边界情况
		backtrack(lindex, rindex-1, str+")", res)
	}
}