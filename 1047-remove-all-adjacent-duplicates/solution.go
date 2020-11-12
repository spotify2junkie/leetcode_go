// array , stack 
func removeDuplicates(S string) string {
    bs := []byte(S)
    stack := []byte{}
    for _, b := range bs {
        if len(stack) > 0 && stack[len(stack)-1] == b {
            stack = stack[:len(stack)-1]
            continue
        }
        stack = append(stack, b)
    }
    return string(stack)
    
}
