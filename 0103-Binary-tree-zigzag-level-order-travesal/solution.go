lfunc zigzagLevelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    queue := []*TreeNode{} // 没有{}
    queue = append(queue,root)
    val_slice := make([][]int, 0)
    val_slice = append(val_slice,[]int{root.Val}) // append root  int 类型 , 不能(python)index取法
    level := 0
    for loc:=1;len(queue) > 0;loc++ {
        cur_slice := []int{}
        size := len(queue)
        for i:=0;i < size;i++ { //  暂存current size , 分号
            if queue[i].Left != nil {
                l := queue[i].Left
                queue = append(queue,l)
                cur_slice = append(cur_slice,l.Val)
            }
            if queue[i].Right != nil {  // 要判断是否为空
                r := queue[i].Right
                queue = append(queue,r)
                cur_slice = append(cur_slice, r.Val)
            }
        }
        // slice可以和slice append
        queue = queue[size:] //新的栈  
        if level%2 == 0 {
			reverse(cur_slice)  //nnn 加个reverse就好
		}
        if len(cur_slice) != 0 {  // 边界问题 
            val_slice = append(val_slice,cur_slice) 
        }
        level++
    }
    return val_slice
}

func reverse(nums []int) {
	start, end := 0, len(nums)-1
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
