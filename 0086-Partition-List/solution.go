// linked list
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	beforeHead := &ListNode{Val: 0, Next: nil} // 取地址
	afterHead := &ListNode{Val: 0, Next: nil}  // 取地址
	for head != nil {
		if head.Val < x {
			a.Next = head
			a = a.Next
		} else {
			b.Next = head
			b = b.Next
		}
		head = head.Next
	}
	b.Next = nil // 定义空
	a.Next = afterHead.Next
	return beforeHead.Next
}
