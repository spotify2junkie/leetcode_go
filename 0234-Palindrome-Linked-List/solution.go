//nnn linked list 
func isPalindrome(head *ListNode) bool {
    res := []int{}
    for head != nil {
        res = append(res,head.Val)
        head = head.Next
    }
    // 按照规则对比值
	l, r := 0, len(res)-1
	for l < r {
		if res[l] != res[r] {
			return false
		} 
		l++ //这个可以
		r-- // 不用来回
	}
	return true
}
