// linked list 
// 相位相加
// for 里两个if 实现 双循环
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	cur := res
	sum := 0 
	carry := 0 // 初始化carry 
	for l1 != nil || l2 != nil || carry > 0 {
		sum = carry // 传入carry

		if l1 != nil { // if 条件, 我还以为是循环
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		} 

		carry = sum / 10 // 相位相加
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}
	return res.Next
}