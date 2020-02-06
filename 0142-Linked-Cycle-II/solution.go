// map  linked list , using hash table 

func detectCycle(head *ListNode) *ListNode {
    cur := head  // 为什么每次都要存一下 可以不存吗 
    checked := make(map[*ListNode]int)  // why use map[*ListNode] 
    for cur != nil {
        if _,ok := checked[cur]; ok { 
            return cur  
        }
        checked[cur] = 1
        cur = cur.Next
    }
    return nil 
}
