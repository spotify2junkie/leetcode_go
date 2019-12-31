// dynamic programming
// this doesnt work , cuz this return all `s`
// func longestValidParentheses(s string) int {
//     if len(s) < 2 {
//         return 0
//     }
//     dp := make([]int,len(s)) // 加上0 的处理
//     stack := []byte{}
//     stack = append(stack,s[0],s[1])
//     if s[0] == '(' && s[1] == ')' {
//         dp[1] = 1
//     }
//     for i:=2; i < len(s); i++ {
//         stack = append(stack,s[i])
//         // for ,i := range stack {

//         // }
//         if (s[i] == ')' && stack[len(stack)-2] == '(') {
//             stack = stack[:len(stack)-2] // 出栈 , 两个都出
//             dp[i] = dp[i-1] + 1
//         } else {
//             dp[i] = dp[i-1]
//         }
//     }
//     return dp[len(s)-1]*2
// }
func longestValidParentheses(s string) int {
	if len(s) == 0 {
		return 0
	}
	n, left, max := len(s), 0, 0
	d := make([]int, n)

	for i := 0; i < n; i++ {
		if s[i] == '(' {
			left++
		} else if left > 0 {
			d[i] = d[i-1] + 2
			if i-d[i] > 0 {
				d[i] += d[i-d[i]]
			}
			max = Max(max, d[i])
			left--
		}
	}

	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}