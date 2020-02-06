func evalRPN(tokens []string) int {
	sn := new([]int)
	for _, v := range tokens {
		if helper(sn, v) {
			return 0
		}
	}
	return (*sn)[0]
}

func helper(sn *[]int, s string) bool {
	if s == "+" || s == "-" || s == "*" || s == "/" {
		if len(*sn) < 2 {
			return true
		}
		tmp := 0
		if s == "+" {
			tmp = (*sn)[len(*sn) - 1] + (*sn)[len(*sn) - 2]
		} else if s == "-" {
			tmp = (*sn)[len(*sn) - 2] - (*sn)[len(*sn) - 1]
		} else if s == "*" {
			tmp = (*sn)[len(*sn) - 1] * (*sn)[len(*sn) - 2]
		} else if s == "/" {
			tmp = (*sn)[len(*sn) - 2] / (*sn)[len(*sn) - 1]
		}
		*sn = (*sn)[:len(*sn) - 2]
		*sn = append(*sn, tmp)
	} else {
		*sn = append(*sn, atoi(s))
	}
	return false
}

func atoi(s string) int {
	flag := 1
	if s[0] == '-' {
		flag = - flag
		s = s[1:]
	}
	rst := 0
	for i := 0; i < len(s); i++ {
		rst *= 10
		rst += int(s[i] - '0')
	}
	return rst * flag
}
