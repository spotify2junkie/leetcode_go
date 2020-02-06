/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//nnn linked list 
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    for head != nil {
        temp := head.Next  // 暂存
        head.Next = prev // 指向之前的节点
        prev = head  // 存入之前的节点
        head = temp  // 找到下一个开始的地方
    }
    return prev
}
