// Recursion DFS Hash table 

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func buildTree(preorder []int, inorder []int) *TreeNode {
    inPos := make(map[int]int)
    for i:= 0; i < len(inorder);i++ {
        inPos[inorder[i]] = i 
    }

    return build(preorder, 0, len(preorder)-1, 0, inPos) 
}

func build(pre []int, prestart int,preend int,instart int,inPos map[int]int) *TreeNode {
    if prestart > preend {
        return nil 
    }
    root := &TreeNode{Val:pre[prestart]}
    rootidx := inPos[pre[prestart]] // get root_index
    leftlen := rootidx - instart
    root.Left = build(pre,prestart+1,prestart+leftlen,instart,inPos)
    root.Right = build(pre,prestart+leftlen+1,preend,rootidx+1,inPos)
    return root 
}