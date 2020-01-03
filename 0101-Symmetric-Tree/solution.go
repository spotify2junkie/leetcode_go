/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true 
    }
    return ish(root.Left,root.Right)
}


func ish(l, r *TreeNode) bool {
    if l == nil && r == nil { //mainly for bottom
        return true 
    }

    if l == nil || r == nil { // false- condition
        return false 
    }

    if l.Val != r.Val {   // false condition 
        return false
    }

    return ish(l.Left,r.Right) && ish(l.Right,r.Left)
}