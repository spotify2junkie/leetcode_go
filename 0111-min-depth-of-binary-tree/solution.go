func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	q := []*TreeNode{root}
	lvl := 1

	for len(q) > 0 {
		size := len(q)

		// 向四周散开
		for i := 0; i < size; i++ {
			node := q[i]

			// 只存在左边树或者右边
			if node == nil {
				continue
			}

			if node.Left == nil && node.Right == nil {
				return lvl
			}

			q = append(q, node.Left, node.Right)
		}
		// Resize the slice to avoid adding the same elements.
		q = q[size:]

		lvl++
	}

	return lvl
}