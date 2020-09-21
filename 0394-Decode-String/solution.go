func decodeString(s string) string {
	var res string
	var loopNum int
	stack := []item{}
	for i := 0; i < len(s); i++ {
		v := s[i]
		if v >= '0' && v <= '9' {
			loopNum = loopNum*10 + int(v-'0')
		} else if v == '[' {
			stack = append(stack, item{preStr: res, loopNum: loopNum})
			res = ""
			loopNum = 0
		} else if v == ']' {
			item := stack[len(stack)-1]
			for j := 0; j < item.loopNum; j++ {
				item.preStr += res
			}
			res = item.preStr
			stack = stack[:len(stack)-1]
		} else {
			res += string(v)
		}
	}
	return res
}

type item struct {
	preStr  string
	loopNum int
}
