// Traverse Pointer  DFS 

ffunc convertBST(root *TreeNode) *TreeNode {
    if root == nil {
        return nil 
    }
    sum := 0 
    getAdd(root,&sum) //取地址
    return root
}

func getAdd(root *TreeNode, sum *int) *TreeNode { // int类型
    if root == nil {
        return nil 
    }
    // right 
    getAdd(root.Right,sum)  // 先后顺序
    *sum += root.Val
    root.Val = *sum
    getAdd(root.Left, sum)
    return root
}