// dfs , tree , local reference 
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func pathSum(root *TreeNode, sum int) int {
    if root == nil {
        return 0 
    }
    count := 0 
    var dfs func(*TreeNode, int)
    dfs = func(node *TreeNode, sum int) {
        if node == nil {
            return 
        }
        sum -= node.Val
        if sum == 0 {
            count++
        }
        dfs(node.Left,sum)
        dfs(node.Right,sum)
    }

    dfs(root,sum)

    return count+pathSum(root.Left,sum)+pathSum(root.Right, sum)
}