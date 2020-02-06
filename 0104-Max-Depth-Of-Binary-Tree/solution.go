//nnn binary tree , recursion 
func maxDepth(root *TreeNode) int {
	if root == nil { // 只用查看root 本身不用看两边
        return 0
    }
    return max(maxDepth(root.Left),maxDepth(root.Right))+1
}

func max(a, b int) int {
    if a > b {
        return a
    } else { return b}
}
