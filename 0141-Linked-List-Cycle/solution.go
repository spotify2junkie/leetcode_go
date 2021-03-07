func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return false
	}
	// fast, slow := &ListNode{}, &ListNode{}

	fast, slow := head.Next.Next, head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		if fast.Val == slow.Val {
			return true
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return false
}