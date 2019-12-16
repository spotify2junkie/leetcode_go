/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func flatten(root *TreeNode)  {
    if root == nil {
        return 
    }
    flatten(root.Right)
    flatten(root.Left)
    // for r 
    tmp := root.Right // 暂存
    root.Right, root.Left = root.Left, nil 
    if root.Right != nil { 
        root = root.Right
    }
    root.Right = tmp // 连上
}