// linked list
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	before, after := &ListNode{}, &ListNode{}
	start, mid := before, after
	for head != nil { // 修正认知
		if head.Val < x {
			before.Next = &ListNode{Val: head.Val}
			before = before.Next
		} else {
			after.Next = &ListNode{Val: head.Val}
			after = after.Next
		}
		head = head.Next
	}
	before.Next = mid.Next
	after.Next = nil
	return start.Next
}

