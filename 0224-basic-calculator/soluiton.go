func calculate(s string) int {
    res := 0 
    stack := []int{}
    sign := 1 
    for i:=0;i < len(s);i++ {
        if s[i] >= '0' && s[i] <= '9' {
            num := 0
            for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
                num = 10* num + int(s[i]-'0')
            }
            res += sign * num
            i--
        } else if s[i] == '+' {
            sign = 1
        } else if s[i] == '-' {
            sign = -1
        } else if s[i] == '(' {
            stack = append(stack, res, sign)
            res = 0
            sign = 1
        } else if s[i] == ')' {
            signtmp := stack[len(stack)-1]
            restemp := stack[len(stack)-2]
            stack = stack[:len(stack)-2] // 更新stack
            res = signtmp*res + restemp
            sign = 1
        }
    }
    return res 
}
