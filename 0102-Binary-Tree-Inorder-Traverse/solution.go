/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 怎么控制同一个level append
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	val_slice := [][]int{}
	val_slice = append(val_slice, []int{root.Val})
	for len(queue) > 0 {
		cur_slice := []int{}
		s := len(queue)
		for i := 0; i < s; i++ {
			if queue[i].Left != nil {
				l := queue[i].Left
				queue = append(queue, l)
				cur_slice = append(cur_slice, l.Val)
			}
			if queue[i].Right != nil { // 要判断是否为空
				r := queue[i].Right
				queue = append(queue, r)
				cur_slice = append(cur_slice, r.Val)
			}
		}
		queue = queue[s:] // 更新队列
		if len(cur_slice) != 0 {
			val_slice = append(val_slice, cur_slice)
		}
	}
	return val_slice
}