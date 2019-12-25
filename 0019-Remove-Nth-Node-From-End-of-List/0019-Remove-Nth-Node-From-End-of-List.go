// func removeNthFromEnd(head *ListNode, n int) *ListNode {
//     if head == nil {
//         return nil
//     }
//     for i:=1;head.Next != nil;i++ {
//         if i+1 == n {
//             curNext := head.Next
//             head.Next = curNext.Next
//         }
//     }
//     return head
// }

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    node := &ListNode{ Next : head } // why ? 因为取地址 ,作为指针
    fast, slow, step := node, node, 0
    for step < n {
        fast = fast.Next
        step++
    }
    for fast.Next != nil {
        fast = fast.Next
        slow = slow.Next
    }
    slow.Next = slow.Next.Next
    return node.Next
}
