/*
本来想着通过map得到
查看后面一个元素，也就不要入栈
检查长度就好
*/

// 一份错误解答
// func isValid(s string) bool {
// 	symbol := map[string]string{ "(" : ")", "[" : "]", "{" : "}"}
//     b := []string{}
//     for index, value := range s {
//         if index == len(s) -1 {
//             break
//         }
//         if string(s[index+1]) == symbol[string(value)] {
//             b = append(append(b,string(value)), symbol[string(value)])
//         }
// 
//     }
// 
//     if len(b) == len(s) {
//         return true 
//     }
// 
//     return false
// }
// 但是没考虑{[]} 这种情况就fail了

func isValid(s string) bool {
	stack := []string{}
	// 后括号映射表
	frontBracket := map[string]string{")": "(", "]": "[", "}": "{"} // 绝了后括号解决了匹配问题

	for _, x := range s {
		if x == '(' || x == '[' || x == '{' { // 遇到前括号，入栈
			stack = append(stack, string(x))
		} else if x == ')' || x == ']' || x == '}' { // 遇到后括号，判断
			if len(stack) != 0 && stack[len(stack)-1] == frontBracket[string(x)] { // 栈非空，和栈顶元素匹配，匹配成功，出栈
				stack = stack[0 : len(stack)-1]
			} else { // 栈空或者匹配失败，返回错误
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}