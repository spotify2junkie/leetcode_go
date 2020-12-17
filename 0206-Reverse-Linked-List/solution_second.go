// recursion reverse linkded list
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := reverseList(he ad.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}