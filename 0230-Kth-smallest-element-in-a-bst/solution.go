//nnn recursion 
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func inorderTraversal(root *TreeNode) []int {
    node_val := []int{}
    if root != nil {
        node_val = append(node_val,inorderTraversal(root.Left)...)
        node_val = append(node_val, root.Val)
        node_val = append(node_val,inorderTraversal(root.Right)...)
    }
    return node_val
}

func kthSmallest(root *TreeNode, k int) int {
    res := inorderTraversal(root)
    return res[k-1]
}  
