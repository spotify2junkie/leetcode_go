//nnn linked list
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}          // define 一个空的
	tmp := head                  // tmp 作为head 的执政
	for l1 != nil && l2 != nil { // 边界条件
		if l1.Val < l2.Val {
			tmp.Next = l1
			l1 = l1.Next
		} else {
			tmp.Next = l2
			l2 = l2.Next
		}
		tmp = tmp.Next
	}
	if l1 != nil {
		tmp.Next = l1
	} else {
		tmp.Next = l2
	}
	return head.Next
}

 