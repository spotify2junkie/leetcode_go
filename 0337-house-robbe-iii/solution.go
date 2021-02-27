/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
	memos := make(map[*TreeNode]int)
	if root == nil {
		return 0 //base case
	}
	if _, ok := memos[root]; ok {
		return memos[root]
	}
	yes := root.Val //collect now
	if root.Left != nil {
		yes = yes + rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
		yes = yes + rob(root.Right.Left) + rob(root.Right.Right)
	}

	no := rob(root.Left) + rob(root.Right)
	memos[root] = max(yes, no)
	return memos[root]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}