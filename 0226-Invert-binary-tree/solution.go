//binary tree, recursion
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil 
    }
    
    return helper(root)
}


func helper(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    tmp := root.Left
    tmpr := root.Right
    root.Left = tmpr
    root.Right = tmp
    helper(root.Left)
    helper(root.Right)
    return root 
}