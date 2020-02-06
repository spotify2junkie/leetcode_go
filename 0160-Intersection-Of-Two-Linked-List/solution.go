//nnn linked list
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    a, b := headA,headB
    for a != b {
         if a != nil {
             a = a.Next
         } else {
             a = headB
         }

         if b != nil {
             b = b.Next
         } else {
             b = headA
         }
    }
    return b
}
