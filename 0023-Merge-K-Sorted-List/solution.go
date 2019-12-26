// linked list 
// devide and conquer 
func mergeKLists(lists []*ListNode) *ListNode {
    length := len(lists)
    if length == 0 {
        return nil 
    } 
    if length == 1 {
        return lists[0]
    }
    nums := length /2 
    left := mergeKLists(lists[:nums])
    right := mergeKLists(lists[nums:])
    return merge(left,right)
}

func merge(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2 
    } 

    if l2 == nil {
        return l1
    }

    if l1.Val < l2.Val {
        l1.Next = merge(l1.Next,l2)
        return l1
    } else {
        l2.Next = merge(l2.Next,l1)
        return l2
    }
}
