/*
 * @lc app=leetcode id=147 lang=golang
 *
 * [147] Insertion Sort List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := new(ListNode)
	dummy.Next = head
	cur := head.Next // get next to head
	head.Next = nil
	for cur != nil {
		nextHead := cur.Next
		InsertNode(dummy, cur)
		cur = nextHead
	}
	return dummy.Next
}

func InsertNode(head, node *ListNode) {
	for head.Next != nil && head.Next.Val < node.Val {
		head = head.Next
	}
	node.Next = head.Next
	head.Next = node // reset back
}

// @lc code=end
