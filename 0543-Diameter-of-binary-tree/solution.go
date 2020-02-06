// recursion binary tree
var max_depth int  // 定义一个全局变量比较max_depth 
func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {
        return 0 
    }
    depth(root)
    return max_depth
}


func depth(node *TreeNode) int {
    if node == nil {
        return 0
    }
    max_depth = max(max_depth,depth(node.Right)+depth(node.Left))
    return max(depth(node.Right),depth(node.Left))+1
}

func max(a, b int) int {
    if a >= b {
        return a 
    } else { return b }
}
