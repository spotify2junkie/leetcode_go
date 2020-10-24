// Traverse Pointer  DFS

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	sum := 0
	helper(root, &sum)
	return root
}

func helper(r *TreeNode, sum *int) {
	if r == nil {
		return
	}
	helper(r.Right, sum)
	*sum = *sum + r.Val
	r.Val = *sum
	helper(r.Left, sum)
	return
}