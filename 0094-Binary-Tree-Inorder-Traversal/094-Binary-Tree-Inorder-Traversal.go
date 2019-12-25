/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func inorderTraversal(root *TreeNode) []int {
    stack := []*TreeNode{}
    travel := []int{}
    cur := root
    for cur != nil || len(stack) > 0 {
        for cur != nil {
            stack = append(stack, cur) //入栈
            cur = cur.Left
        }
        cur = stack[len(stack)-1]  //出栈
        stack = stack[:len(stack)-1]
        travel = append(travel,cur.Val)
        cur = cur.Right   //移动节点
    }
    return travel 
}