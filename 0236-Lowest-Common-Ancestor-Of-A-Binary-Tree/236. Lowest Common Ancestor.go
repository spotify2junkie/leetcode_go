/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// think of what this function do ? 
	// first consider strange situation
	if p == root || q == root || root == nil {
		return root
	}

	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)

	if l == nil {
		return r
	}

	if r == nil {
		return l
	}

   // did i just covered all situations ? 
   return root  // this should be all 

	
}