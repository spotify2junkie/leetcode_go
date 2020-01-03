package main

import (
	"fmt"

	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

var maxDepth int

type TreeNode = kit.TreeNode

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// if depth(root.Right) == 0 {
	//     return depth(root.Left)
	// }
	// if depth(root.Left) == 0 {
	//     return depth(root.Right)
	// }
	depth(root)
	return maxDepth
}

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	maxDepth = max(maxDepth, depth(node.Left)+depth(node.Right))
	return max(depth(node.Right), depth(node.Left)) + 1
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

// tcs is testcase slice
var tcs = []struct {
	pre, in []int
	ans     int
}{

	{
		[]int{},
		[]int{},
		0,
	},

	{
		[]int{1, 2, 4, 5, 3},
		[]int{4, 2, 5, 1, 3},
		3,
	},

	// 可以有多个 testcase
}

func main() {
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		root := kit.PreIn2Tree(tc.pre, tc.in)
		if tc.ans == diameterOfBinaryTree(root) {
			fmt.Println("test passed")
		}
	}
}
