// recursion

var succssor *ListNode

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == 1 {
		return reverseRange(head, n)
	}
	head.Next = reverseBetween(head.Next, m-1, n-1) // 相对位置
	return head
}

func reverseRange(head *ListNode, n int) *ListNode {
	if n == 1 {
		succssor = head.Next
		return head
	}
	last := reverseRange(head.Next, n-1)
	head.Next.Next = head
	head.Next = succssor
	return last
}