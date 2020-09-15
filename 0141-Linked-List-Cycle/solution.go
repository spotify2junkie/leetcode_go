// hash table,linked list , two pointes
func hasCycle(head *ListNode) bool {
	hash := make(map[*ListNode]int)
	for head != nil {
		if _, ok := hash[head]; ok { // 查看是否存在
			return true
		}
		hash[head] = head.val // store value
		head = head.Next
	}
	return false

}

// second solution, use two pointers
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return false
	}
	fast, slow := &ListNode{}, &ListNode{}
	fast = head.Next.Next
	slow = head.Next
	for fast != nil && fast.Next != nil { // 只需关注下面用到的fast.Next 是否为空
		if fast.Val == slow.Val {
			return true
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return false
}


