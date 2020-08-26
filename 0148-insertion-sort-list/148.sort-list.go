/*
 * @lc app=leetcode id=148 lang=golang
 *
 * [148] Sort List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	cur := head
	length := 0
	for cur != nil {
		length++
		cur = cur.Next
	}
	if length <= 1 {
		return head
	}
	middleHead := middleNode(head) // 得到中间点
	cur = middleHead.Next          // 隔断
	middleHead.Next = nil
	middleHead = cur

	left := sortList(head)
	right := sortList(middleHead)
	return mergeTwoSortedList(left, right)
}

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p1, p2 := head, head
	if p2.Next != nil && p2.Next.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	return p1 // contains this node as first half
}

func mergeTwoSortedList(l, r *ListNode) *ListNode {
	if l == nil {
		return r
	}
	if r == nil {
		return l
	}
	if l.Val < r.Val {
		l.Next = mergeTwoSortedList(l.Next, r)
		return l
	}
	r.Next = mergeTwoSortedList(l, r.Next) // 注意
	return r
}

// @lc code=end