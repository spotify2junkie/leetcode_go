// func maxDepth(root *TreeNode) int {
//     if root == nil {
//         return 0
//     }
//     return helper(root.Left,root.Right)
// }

// func helper(l, r *TreeNode) int {
//     depth := 0
//     if l.Left == nil && l.right == nil && r.Left == nil && r.Right nil {
//         return depth 
//     } else { 
//         return helper(l, r)+1
//     }
// }

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